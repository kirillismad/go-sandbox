package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	opensearch "github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type config struct {
	CSVPath            string
	Index              string
	Addresses          []string
	Username           string
	Password           string
	BatchSize          int
	RequestTimeout     time.Duration
	InsecureSkipVerify bool
	IDPrefix           string
}

type coffeeDocument struct {
	Metadata coffeeMetadata `json:"metadata"`
	Raw      coffeeRaw      `json:"raw"`
}

type coffeeMetadata struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Type          string   `json:"type"`
	Tags          []string `json:"tags"`
	CookingType   string   `json:"cooking_type"`
	Price         int      `json:"price"`
	CommentsCount int      `json:"comments_count"`
	Rating        int      `json:"rating"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type coffeeRaw struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Type          string   `json:"type"`
	Tags          []string `json:"tags"`
	CookingType   string   `json:"cooking_type"`
	Price         int      `json:"price"`
	CommentsCount int      `json:"comments_count"`
	Rating        float64  `json:"rating"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type bulkAction struct {
	Index bulkActionMeta `json:"index"`
}

type bulkActionMeta struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
}

type bulkResponse struct {
	Errors bool                       `json:"errors"`
	Items  []map[string]bulkItemEntry `json:"items"`
}

type bulkItemEntry struct {
	Status int             `json:"status"`
	Error  json.RawMessage `json:"error"`
	ID     string          `json:"_id"`
}

func main() {
	cfg := parseFlags()
	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseFlags() config {
	var addresses string

	cfg := config{}
	flag.StringVar(&cfg.CSVPath, "csv", "opensearch/coffee_items.csv", "path to the CSV file to upload")
	flag.StringVar(&cfg.Index, "index", "coffee-items", "OpenSearch index name")
	flag.StringVar(&addresses, "addr", envOrDefault("OPENSEARCH_URL", "http://localhost:9200"), "comma-separated OpenSearch node addresses")
	flag.StringVar(&cfg.Username, "username", os.Getenv("OPENSEARCH_USERNAME"), "OpenSearch username")
	flag.StringVar(&cfg.Password, "password", os.Getenv("OPENSEARCH_PASSWORD"), "OpenSearch password")
	flag.IntVar(&cfg.BatchSize, "batch-size", 500, "number of documents per bulk request")
	flag.DurationVar(&cfg.RequestTimeout, "timeout", 30*time.Second, "timeout for each bulk request")
	flag.BoolVar(&cfg.InsecureSkipVerify, "insecure", false, "skip TLS certificate verification")
	flag.StringVar(&cfg.IDPrefix, "id-prefix", "coffee-", "document ID prefix")
	flag.Parse()

	cfg.Addresses = splitAndTrim(addresses)
	return cfg
}

func run(cfg config) error {
	if cfg.CSVPath == "" {
		return errors.New("csv path is required")
	}
	if cfg.Index == "" {
		return errors.New("index is required")
	}
	if cfg.BatchSize <= 0 {
		return errors.New("batch-size must be greater than 0")
	}
	if len(cfg.Addresses) == 0 {
		return errors.New("at least one OpenSearch address is required")
	}

	client, err := newClient(cfg)
	if err != nil {
		return fmt.Errorf("create opensearch client: %w", err)
	}

	file, err := os.Open(cfg.CSVPath)
	if err != nil {
		return fmt.Errorf("open csv file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("read csv header: %w", err)
	}

	columns := make(map[string]int, len(header))
	for i, name := range header {
		columns[name] = i
	}

	if err := validateColumns(columns); err != nil {
		return err
	}

	buffer := bytes.NewBuffer(nil)
	batchCount := 0
	totalCount := 0
	lineNumber := 1

	flush := func() error {
		if batchCount == 0 {
			return nil
		}

		if err := sendBulk(context.Background(), client, cfg, buffer); err != nil {
			return err
		}

		buffer.Reset()
		batchCount = 0
		return nil
	}

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("read csv row %d: %w", lineNumber+1, err)
		}

		lineNumber++
		docID := fmt.Sprintf("%s%d", cfg.IDPrefix, lineNumber-1)
		document, err := buildDocument(record, columns)
		if err != nil {
			return fmt.Errorf("parse csv row %d: %w", lineNumber, err)
		}

		if err := appendBulkDocument(buffer, cfg.Index, docID, document); err != nil {
			return fmt.Errorf("encode bulk row %d: %w", lineNumber, err)
		}

		batchCount++
		totalCount++
		if batchCount >= cfg.BatchSize {
			if err := flush(); err != nil {
				return err
			}
		}
	}

	if err := flush(); err != nil {
		return err
	}

	fmt.Printf("uploaded %d documents into index %s\n", totalCount, cfg.Index)
	return nil
}

func newClient(cfg config) (*opensearch.Client, error) {
	httpTransport := &http.Transport{}
	if cfg.InsecureSkipVerify {
		httpTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	clientConfig := opensearch.Config{
		Addresses:  cfg.Addresses,
		Username:   cfg.Username,
		Password:   cfg.Password,
		Transport:  httpTransport,
		MaxRetries: 3,
	}

	return opensearch.NewClient(clientConfig)
}

func validateColumns(columns map[string]int) error {
	required := []string{
		"metadata_title",
		"metadata_description",
		"metadata_type",
		"metadata_tags",
		"metadata_cooking_type",
		"metadata_price",
		"metadata_comments_count",
		"metadata_rating",
		"metadata_created_at",
		"metadata_updated_at",
		"raw_title",
		"raw_description",
		"raw_type",
		"raw_tags",
		"raw_cooking_type",
		"raw_price",
		"raw_comments_count",
		"raw_rating",
		"raw_created_at",
		"raw_updated_at",
	}

	for _, name := range required {
		if _, ok := columns[name]; !ok {
			return fmt.Errorf("missing required csv column %q", name)
		}
	}

	return nil
}

func buildDocument(record []string, columns map[string]int) (coffeeDocument, error) {
	metadataTags, err := parseTags(valueAt(record, columns, "metadata_tags"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_tags: %w", err)
	}
	rawTags, err := parseTags(valueAt(record, columns, "raw_tags"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_tags: %w", err)
	}

	metadataPrice, err := strconv.Atoi(valueAt(record, columns, "metadata_price"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_price: %w", err)
	}
	metadataCommentsCount, err := strconv.Atoi(valueAt(record, columns, "metadata_comments_count"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_comments_count: %w", err)
	}
	metadataRating, err := strconv.Atoi(valueAt(record, columns, "metadata_rating"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_rating: %w", err)
	}
	if err := validateRFC3339(valueAt(record, columns, "metadata_created_at")); err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_created_at: %w", err)
	}
	if err := validateRFC3339(valueAt(record, columns, "metadata_updated_at")); err != nil {
		return coffeeDocument{}, fmt.Errorf("metadata_updated_at: %w", err)
	}

	rawPrice, err := strconv.Atoi(valueAt(record, columns, "raw_price"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_price: %w", err)
	}
	rawCommentsCount, err := strconv.Atoi(valueAt(record, columns, "raw_comments_count"))
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_comments_count: %w", err)
	}
	rawRating, err := strconv.ParseFloat(valueAt(record, columns, "raw_rating"), 64)
	if err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_rating: %w", err)
	}
	if err := validateRFC3339(valueAt(record, columns, "raw_created_at")); err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_created_at: %w", err)
	}
	if err := validateRFC3339(valueAt(record, columns, "raw_updated_at")); err != nil {
		return coffeeDocument{}, fmt.Errorf("raw_updated_at: %w", err)
	}

	return coffeeDocument{
		Metadata: coffeeMetadata{
			Title:         valueAt(record, columns, "metadata_title"),
			Description:   valueAt(record, columns, "metadata_description"),
			Type:          valueAt(record, columns, "metadata_type"),
			Tags:          metadataTags,
			CookingType:   valueAt(record, columns, "metadata_cooking_type"),
			Price:         metadataPrice,
			CommentsCount: metadataCommentsCount,
			Rating:        metadataRating,
			CreatedAt:     valueAt(record, columns, "metadata_created_at"),
			UpdatedAt:     valueAt(record, columns, "metadata_updated_at"),
		},
		Raw: coffeeRaw{
			Title:         valueAt(record, columns, "raw_title"),
			Description:   valueAt(record, columns, "raw_description"),
			Type:          valueAt(record, columns, "raw_type"),
			Tags:          rawTags,
			CookingType:   valueAt(record, columns, "raw_cooking_type"),
			Price:         rawPrice,
			CommentsCount: rawCommentsCount,
			Rating:        rawRating,
			CreatedAt:     valueAt(record, columns, "raw_created_at"),
			UpdatedAt:     valueAt(record, columns, "raw_updated_at"),
		},
	}, nil
}

func appendBulkDocument(buffer *bytes.Buffer, index string, docID string, document coffeeDocument) error {
	action := bulkAction{
		Index: bulkActionMeta{
			Index: index,
			ID:    docID,
		},
	}

	if err := json.NewEncoder(buffer).Encode(action); err != nil {
		return err
	}
	if err := json.NewEncoder(buffer).Encode(document); err != nil {
		return err
	}

	return nil
}

func sendBulk(parent context.Context, client *opensearch.Client, cfg config, payload *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(parent, cfg.RequestTimeout)
	defer cancel()

	request := opensearchapi.BulkRequest{
		Body:    bytes.NewReader(payload.Bytes()),
		Refresh: "false",
	}

	response, err := request.Do(ctx, client)
	if err != nil {
		return fmt.Errorf("bulk request failed: %w", err)
	}
	defer response.Body.Close()

	if response.IsError() {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("bulk request returned %s: %s", response.Status(), strings.TrimSpace(string(body)))
	}

	var bulkResult bulkResponse
	if err := json.NewDecoder(response.Body).Decode(&bulkResult); err != nil {
		return fmt.Errorf("decode bulk response: %w", err)
	}

	if !bulkResult.Errors {
		return nil
	}

	failures := make([]string, 0, 5)
	for _, item := range bulkResult.Items {
		for op, result := range item {
			if result.Status < 300 {
				continue
			}

			message := fmt.Sprintf("%s %s returned status %d", op, result.ID, result.Status)
			if len(result.Error) > 0 {
				message = fmt.Sprintf("%s: %s", message, strings.TrimSpace(string(result.Error)))
			}
			failures = append(failures, message)
			if len(failures) == cap(failures) {
				return fmt.Errorf("bulk request had item failures: %s", strings.Join(failures, "; "))
			}
		}
	}

	return fmt.Errorf("bulk request had item failures: %s", strings.Join(failures, "; "))
}

func parseTags(value string) ([]string, error) {
	var tags []string
	if err := json.Unmarshal([]byte(value), &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func validateRFC3339(value string) error {
	_, err := time.Parse(time.RFC3339, value)
	return err
}

func valueAt(record []string, columns map[string]int, name string) string {
	idx := columns[name]
	if idx >= len(record) {
		return ""
	}
	return record[idx]
}

func splitAndTrim(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func envOrDefault(name string, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}
