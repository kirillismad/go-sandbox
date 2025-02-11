package confluent

import (
	"encoding/json"
	"strconv"

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
		"linger.ms":                             10,
		"batch.num.messages":                    2,
	})
	s.Require().NoError(err)
	defer p.Close()
	defer p.Flush(5000)

	go func() {
		for event := range p.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				s.Require().NoError(ev.TopicPartition.Error)
				s.T().Log(ev.TopicPartition)
			}
		}
	}()

	const cnt = 2
	for range cnt {
		s.produce(p, Message{ID: gofakeit.Int64(), Name: gofakeit.Name()})
	}
}

func (s *KafkaTestSuite) produce(p *kafka.Producer, msg Message) {
	content, err := json.Marshal(msg)
	s.Require().NoError(err)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: lo.ToPtr(topic), Partition: kafka.PartitionAny},
		Value:          content,
		Key:            []byte(strconv.FormatInt(msg.ID, 10)),
		Headers: []kafka.Header{
			{
				Key:   "x-custom-header",
				Value: []byte(gofakeit.UUID()),
			},
		},
	}, nil)
	s.Require().NoError(err)
}
