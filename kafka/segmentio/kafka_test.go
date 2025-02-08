package segmentio

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/suite"
)

func int64ToBytes(val int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	return buf[:binary.PutVarint(buf, val)]
}

type KafkaTestSuite struct {
	suite.Suite
}

func (s *KafkaTestSuite) SetupSuite() {

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
		RequiredAcks:           kafka.RequireOne,
		Async:                  false,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
	}

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
	if errors.Is(err, kafka.UnknownTopicOrPartition) {
		time.Sleep(1 * time.Second)
		err = writer.WriteMessages(ctx, msg)
	}
	s.Require().NoError(err)
}

func TestKafka(t *testing.T) {
	suite.Run(t, new(KafkaTestSuite))
}
