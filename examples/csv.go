package examples

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sandbox/utils"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

const filePath = "./src/csvExample.csv"
const count = 1000

func DemoCsv() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := writeCSV(ctx, filePath, []string{"id", "name", "age"}, genRecords(ctx, count))
	if err != nil {
		log.Fatalln(err)
	}

	records, err := readCSV(ctx, filePath)
	if err != nil {
		log.Fatalln(err)
	}
	i := 1
	for record := range records {
		if record.err != nil {
			log.Fatalln(record.err)
		}
		fmt.Fprintf(os.Stdout, "%v) %v\n", i, record.result)
		i++
	}
}

type CSVRecord interface {
	ToRecord() []string
}
type csvRecord struct {
	ID   uuid.UUID
	Name string
	Age  uint
}

func (record csvRecord) ToRecord() []string {
	return []string{record.ID.String(), record.Name, strconv.Itoa(int(record.Age))}
}

func randomRecord() csvRecord {
	r := csvRecord{
		ID:   uuid.New(),
		Name: rstring(),
		Age:  uint(rand.Intn(100-20) + 20),
	}
	return r
}

func genRecords(ctx context.Context, count int) <-chan CSVRecord {
	const workers = 10
	quotioent, module := count/workers, count%workers
	chunks := make([]int, workers)
	for i := range chunks {
		chunks[i] = quotioent
		if module-i > 0 {
			chunks[i]++
		}
	}

	out := make(chan CSVRecord, workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	go func() {
		defer close(out)
		wg.Wait()
	}()

	for i := 0; i < workers; i++ {
		i := i
		go func() {
			defer wg.Done()
			for j := 0; j < chunks[i]; j++ {
				select {
				case out <- randomRecord():
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	return out
}

func writeCSV(ctx context.Context, filePath string, header []string, input <-chan CSVRecord) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(bufio.NewWriter(file))

	if err = writer.Write(header); err != nil {
		return err
	}

loop:
	for record := range input {
		select {
		case <-ctx.Done():
			break loop
		default:
			if err = writer.Write(record.ToRecord()); err != nil {
				return err
			}
		}
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		return err
	}
	return nil
}

type csvRow[T any] struct {
	err    error
	result T
}

func readCSV(ctx context.Context, filePath string) (chan csvRow[csvRecord], error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	out := make(chan csvRow[csvRecord], 1)
	reader := csv.NewReader(bufio.NewReader(file))
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	go func() {
		defer close(out)
		defer file.Close()
		for {
			r, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				select {
				case <-ctx.Done():
				case out <- csvRow[csvRecord]{err: err}:
					return
				}
			}
			record := csvRecord{
				ID:   uuid.MustParse(r[0]),
				Name: r[1],
				Age:  uint(utils.Must(strconv.Atoi(r[2]))),
			}
			select {
			case <-ctx.Done():
				return
			case out <- csvRow[csvRecord]{result: record}:
			}

		}
	}()

	return out, nil
}
