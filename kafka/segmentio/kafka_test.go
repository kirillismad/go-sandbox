package segmentio

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sandbox/utils"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/suite"
)

type KafkaTestSuite struct {
	suite.Suite
}

func (s *KafkaTestSuite) SetupSuite() {
	// make sure the topic exists
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, 0)
	if err != nil {
		s.Require().NoError(err)
	}
	conn.Close()
}

type Message struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *KafkaTestSuite) TestProducerSync() {
	ctx := context.Background()

	writer := &kafka.Writer{
		Addr:                   kafka.TCP("kafka:9092"),
		Topic:                  topic,
		Balancer:               &kafka.CRC32Balancer{},
		RequiredAcks:           kafka.RequireAll,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	m := Message{
		ID:   gofakeit.Int64(),
		Name: gofakeit.Name(),
	}
	content, err := json.Marshal(m)
	s.Require().NoError(err)

	msg := kafka.Message{
		Key:   []byte(strconv.FormatInt(m.ID, 10)),
		Value: content,
		Headers: []kafka.Header{
			{
				Key:   "x-custom-header",
				Value: []byte(gofakeit.UUID()),
			},
		},
		Time: time.Now(),
	}

	err = writer.WriteMessages(ctx, msg)
	// retry if the topic does not exist
	if errors.Is(err, kafka.UnknownTopicOrPartition) {
		time.Sleep(1 * time.Second)
		err = writer.WriteMessages(ctx, msg)
	}
	s.Require().NoError(err)
}

func (s *KafkaTestSuite) TestProducerAsync() {
	ctx := context.Background()

	done := make(chan struct{})

	writer := &kafka.Writer{
		Addr:         kafka.TCP("kafka:9092"),
		Topic:        topic,
		Balancer:     &kafka.CRC32Balancer{},
		RequiredAcks: kafka.RequireOne,
		Async:        true,
		Completion: func(messages []kafka.Message, err error) {
			s.Require().NoError(err)
			close(done)
		},
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	m := Message{
		ID:   gofakeit.Int64(),
		Name: gofakeit.Name(),
	}
	content, err := json.Marshal(m)
	s.Require().NoError(err)

	msg := kafka.Message{
		Key:   []byte(strconv.FormatInt(m.ID, 10)),
		Value: content,
		Headers: []kafka.Header{
			{
				Key:   "x-custom-header",
				Value: []byte(gofakeit.UUID()),
			},
		},
		Time: time.Now(),
	}

	err = writer.WriteMessages(ctx, msg)
	s.Require().NoError(err)
	<-done
}

func (s *KafkaTestSuite) TestConsumer() {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		GroupID: "segmentio-group",
		Topic:   topic,
	})
	defer reader.Close()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	m, err := reader.FetchMessage(ctx)
	if errors.Is(err, context.DeadlineExceeded) {
		s.T().Log("no message received")
		return
	}
	s.Require().NoError(err)

	headers := utils.Map(m.Headers, func(h kafka.Header) string {
		return string(h.Key) + ":" + string(h.Value)
	})

	key, err := strconv.Atoi(string(m.Key))
	s.Require().NoError(err)

	var msg Message
	err = json.Unmarshal(m.Value, &msg)
	s.Require().NoError(err)

	s.T().Logf("partition: %d, offset: %d headers: %v, key: %d, message: %v", m.Partition, m.Offset, headers, key, msg)

	err = reader.CommitMessages(ctx, m)
	s.Require().NoError(err)
}

func (s *KafkaTestSuite) TestPipeline() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	logs := make(chan string)
	defer close(logs)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case log, ok := <-logs:
				if !ok {
					return
				}
				s.T().Log(log)
			}
		}
	}()

	writer := &kafka.Writer{
		Addr:         kafka.TCP("kafka:9092"),
		Topic:        topic,
		Balancer:     &kafka.CRC32Balancer{},
		RequiredAcks: kafka.RequireAll,
		Compression:  kafka.Snappy,
		Completion: func(messages []kafka.Message, err error) {
			if len(messages) == 1 {
				partition := messages[0].Partition
				offset := messages[0].Offset
				key := string(messages[0].Key)
				logs <- fmt.Sprintf("PRODUCER // partition: %d, offset: %d, key: %s", partition, offset, key)
				return
			}
		},
	}
	defer writer.Close()

	producer := func(ctx context.Context, id int) chan error {
		out := make(chan error)
		go func() {
			for {
				m := Message{
					ID:   gofakeit.Int64(),
					Name: gofakeit.Name(),
				}
				content, err := json.Marshal(m)
				if err != nil {
					out <- fmt.Errorf("failed to marshal message: %w", err)
					return
				}

				msg := kafka.Message{
					Headers: []kafka.Header{{Key: "x-id", Value: []byte(strconv.Itoa(id))}},
					Key:     []byte(strconv.FormatInt(m.ID, 10)),
					Value:   content,
					Time:    time.Now(),
				}

				err = writer.WriteMessages(ctx, msg)
				if err != nil {
					out <- fmt.Errorf("failed to write message: %w", err)
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()
		return out
	}
	consumer := func(ctx context.Context) chan error {
		out := make(chan error)
		go func() {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers: []string{"kafka:9092"},
				GroupID: "segmentio-group",
				Topic:   topic,
			})
			defer reader.Close()

			for {
				select {
				case <-ctx.Done():
					out <- ctx.Err()
					return
				default:
					m, err := reader.FetchMessage(ctx)
					if err != nil {
						out <- fmt.Errorf("failed to fetch message: %w", err)
						return
					}

					key, err := strconv.Atoi(string(m.Key))
					if err != nil {
						out <- fmt.Errorf("failed to convert key: %w", err)
						return
					}

					var msg Message
					err = json.Unmarshal(m.Value, &msg)
					if err != nil {
						out <- fmt.Errorf("failed to unmarshal message: %w", err)
						return
					}

					logs <- fmt.Sprintf("CONSUMER // partition: %d, offset: %d, key: %d", m.Partition, m.Offset, key)
					err = reader.CommitMessages(ctx, m)
					if err != nil {
						out <- fmt.Errorf("failed to commit message: %w", err)
						return
					}
				}
			}
		}()
		return out
	}

	select {
	case <-ctx.Done():
		return
	// case err := <-producer(ctx, 1):
	// 	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
	// 		return
	// 	}
	// 	s.Require().NoError(err)
	case err := <-producer(ctx, 2):
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return
		}
		s.Require().NoError(err)
	// case err := <-consumer(ctx):
	// 	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
	// 		return
	// 	}
	// 	s.Require().NoError(err)
	case err := <-consumer(ctx):
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return
		}
		s.Require().NoError(err)
	}
}

func TestKafka(t *testing.T) {
	suite.Run(t, new(KafkaTestSuite))
}
