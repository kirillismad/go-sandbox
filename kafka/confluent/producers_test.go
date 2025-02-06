package confluent

import (
	"encoding/binary"
	"encoding/json"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samber/lo"
)

type Message struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *KafkaTestSuite) TestProducerAtLeastOnce() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     "kafka:9092",
		"acks":                                  "all",
		"retries":                               3,
		"retry.backoff.ms":                      200,
		"max.in.flight.requests.per.connection": 5,
		"socket.keepalive.enable":               true,
		"compression.type":                      "snappy",
		"linger.ms":                             10,
		"batch.num.messages":                    2,
		"queue.buffering.max.ms":                0,
	})
	s.Require().NoError(err)
	defer p.Close()
	defer p.Flush(1000)

	go func() {
		for event := range p.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				s.Require().NoError(ev.TopicPartition.Error)
				s.T().Log(ev.TopicPartition)
			}
		}
	}()

	const cnt = 4
	for range cnt {
		s.produce(p, Message{ID: gofakeit.Int64(), Name: gofakeit.Name()})
	}
}

func (s *KafkaTestSuite) TestProducerAtMostOnce() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     "kafka:9092",
		"acks":                                  "1", // acks = 1
		"retries":                               1,
		"retry.backoff.ms":                      200,
		"max.in.flight.requests.per.connection": 5,
		"socket.keepalive.enable":               true,
		"compression.type":                      "snappy",
		"linger.ms":                             10,
		"batch.num.messages":                    2,
		"queue.buffering.max.ms":                0,
	})
	s.Require().NoError(err)
	defer p.Close()
	defer p.Flush(1000)

	go func() {
		for event := range p.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				s.Require().NoError(ev.TopicPartition.Error)
				s.T().Log(ev.TopicPartition)
			}
		}
	}()

	const cnt = 4
	for range cnt {
		s.produce(p, Message{ID: gofakeit.Int64(), Name: gofakeit.Name()})
	}
}

func (s *KafkaTestSuite) produce(p *kafka.Producer, msg Message) {
	content, err := json.Marshal(msg)
	s.Require().NoError(err)

	buf := make([]byte, binary.MaxVarintLen64)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: lo.ToPtr(topic), Partition: kafka.PartitionAny},
		Value:          content,
		Key:            buf[:binary.PutVarint(buf, msg.ID)],
		Headers: []kafka.Header{
			{
				Key:   "x-custom-header",
				Value: []byte(gofakeit.UUID()),
			},
		},
	}, nil)
	s.Require().NoError(err)
}
