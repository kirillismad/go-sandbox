package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var (
	coffeeTypes  = []string{"single_origin", "mixed", "blend", "flavored", "decaf"}
	cookingTypes = []string{"espresso", "pour_over", "french_press", "aeropress", "cold_brew", "turkish"}
	tagPool      = []string{"chocolate", "citrus", "nutty", "berry", "caramel", "floral", "bold", "sweet", "spicy", "smooth", "balanced", "fruity"}
)

type coffeeMetadata struct {
	ID            string
	Title         string
	Description   string
	Type          string
	Tags          []string
	CookingType   string
	Price         int
	CommentsCount int
	Rating        int
	CreatedAt     string
	UpdatedAt     string
}

type coffeeRaw struct {
	ID            string
	Title         string
	Description   string
	Type          string
	Tags          []string
	CookingType   string
	Price         int
	CommentsCount int
	Rating        float64
	CreatedAt     string
	UpdatedAt     string
}

type coffeeItem struct {
	Metadata coffeeMetadata
	Raw      coffeeRaw
}

func main() {
	count := flag.Int("n", 100, "number of coffee items to generate")
	outPath := flag.String("out", "coffee_items.csv", "path to the output CSV file")
	seed := flag.Int64("seed", time.Now().UnixNano(), "random seed")
	flag.Parse()

	if *count <= 0 {
		fmt.Fprintln(os.Stderr, "n must be greater than 0")
		os.Exit(1)
	}

	gofakeit.Seed(*seed)

	if err := ensureOutputDir(*outPath); err != nil {
		fmt.Fprintf(os.Stderr, "prepare output path: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(*outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(csvHeader()); err != nil {
		fmt.Fprintf(os.Stderr, "write header: %v\n", err)
		os.Exit(1)
	}

	for range *count {
		item := generateCoffeeItem()
		if err := writer.Write(item.csvRecord()); err != nil {
			fmt.Fprintf(os.Stderr, "write record: %v\n", err)
			os.Exit(1)
		}
	}

	if err := writer.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "flush csv: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generated %d coffee items in %s\n", *count, *outPath)
}

func ensureOutputDir(outPath string) error {
	dir := filepath.Dir(outPath)
	if dir == "." {
		return nil
	}

	return os.MkdirAll(dir, 0o755)
}

func csvHeader() []string {
	return []string{
		"metadata_id",
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
		"raw_id",
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
}

func generateCoffeeItem() coffeeItem {
	id, err := uuid.NewV7()
	if err != nil {
		panic(fmt.Sprintf("generate uuidv7: %v", err))
	}
	itemID := id.String()

	createdAt := gofakeit.DateRange(time.Now().AddDate(-2, 0, 0), time.Now().AddDate(0, 0, -1)).UTC().Truncate(time.Second)
	updatedAt := gofakeit.DateRange(createdAt, time.Now()).UTC().Truncate(time.Second)
	rawTags := randomTags()
	rawType := gofakeit.RandomString(coffeeTypes)
	rawCookingType := gofakeit.RandomString(cookingTypes)
	rawRating := math.Round(gofakeit.Float64Range(2.5, 5.0)*100) / 100
	rawTitle := fmt.Sprintf("%s %s", titleWord(), titleWord())
	rawDescription := strings.TrimSpace(fmt.Sprintf("%s %s", gofakeit.Sentence(8), gofakeit.Sentence(10)))
	rawPrice := gofakeit.Number(700, 4500) * 10
	rawCommentsCount := gofakeit.Number(0, 500)

	raw := coffeeRaw{
		ID:            itemID,
		Title:         rawTitle,
		Description:   rawDescription,
		Type:          rawType,
		Tags:          rawTags,
		CookingType:   rawCookingType,
		Price:         rawPrice,
		CommentsCount: rawCommentsCount,
		Rating:        rawRating,
		CreatedAt:     createdAt.Format(time.RFC3339),
		UpdatedAt:     updatedAt.Format(time.RFC3339),
	}

	metadata := coffeeMetadata{
		ID:            itemID,
		Title:         normalizeForMetadata(raw.Title),
		Description:   normalizeForMetadata(raw.Description),
		Type:          normalizeForMetadata(raw.Type),
		Tags:          normalizeTags(raw.Tags),
		CookingType:   normalizeForMetadata(raw.CookingType),
		Price:         raw.Price,
		CommentsCount: raw.CommentsCount,
		Rating:        int(math.Round(raw.Rating * 100)),
		CreatedAt:     raw.CreatedAt,
		UpdatedAt:     raw.UpdatedAt,
	}

	return coffeeItem{
		Metadata: metadata,
		Raw:      raw,
	}
}

func titleWord() string {
	return strings.Title(strings.ToLower(gofakeit.RandomString([]string{
		"velvet", "mountain", "sunrise", "ember", "forest", "gold", "harbor", "midnight", "amber", "dune", "cedar", "bloom",
	})))
}

func randomTags() []string {
	count := gofakeit.Number(2, 5)
	seen := make(map[string]struct{}, count)
	tags := make([]string, 0, count)
	for len(tags) < count {
		tag := gofakeit.RandomString(tagPool)
		if _, ok := seen[tag]; ok {
			continue
		}
		seen[tag] = struct{}{}
		tags = append(tags, tag)
	}
	return tags
}

func normalizeForMetadata(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeTags(tags []string) []string {
	normalized := make([]string, len(tags))
	for i, tag := range tags {
		normalized[i] = normalizeForMetadata(tag)
	}
	slices.Sort(normalized)
	return normalized
}

func tagsAsJSON(tags []string) string {
	encoded, err := json.Marshal(tags)
	if err != nil {
		return "[]"
	}
	return string(encoded)
}

func (item coffeeItem) csvRecord() []string {
	return []string{
		item.Metadata.ID,
		item.Metadata.Title,
		item.Metadata.Description,
		item.Metadata.Type,
		tagsAsJSON(item.Metadata.Tags),
		item.Metadata.CookingType,
		fmt.Sprintf("%d", item.Metadata.Price),
		fmt.Sprintf("%d", item.Metadata.CommentsCount),
		fmt.Sprintf("%d", item.Metadata.Rating),
		item.Metadata.CreatedAt,
		item.Metadata.UpdatedAt,
		item.Raw.ID,
		item.Raw.Title,
		item.Raw.Description,
		item.Raw.Type,
		tagsAsJSON(item.Raw.Tags),
		item.Raw.CookingType,
		fmt.Sprintf("%d", item.Raw.Price),
		fmt.Sprintf("%d", item.Raw.CommentsCount),
		fmt.Sprintf("%.2f", item.Raw.Rating),
		item.Raw.CreatedAt,
		item.Raw.UpdatedAt,
	}
}
