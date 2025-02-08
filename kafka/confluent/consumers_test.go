package confluent

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sandbox/utils"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samber/lo"
)

func IsTimeout(err error) bool {
	if e, ok := lo.ErrorsAs[kafka.Error](err); ok && e.IsTimeout() {
		return true
	}
	return false
}

func (s *KafkaTestSuite) TestConsumerAtLeastOnce() {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  "kafka:9092",
		"group.id":           "consumer-group-1",
		"auto.offset.reset":  "earliest", // earliest - читать с начала, если нет коммита
		"enable.auto.commit": false,      // Коммиты делаем вручную
	}

	с, err := kafka.NewConsumer(config)
	s.Require().NoError(err)
	defer с.Close()

	err = с.SubscribeTopics([]string{topic}, nil)
	s.Require().NoError(err)

	// Контекст для graceful shutdown

	for range 2 {
		msg, err := с.ReadMessage(5 * time.Second)
		if err != nil {
			if IsTimeout(err) {
				s.T().Log("skip")
				continue
			}
			s.Require().NoError(err)
		}
		var m Message
		err = json.Unmarshal(msg.Value, &m)
		s.Require().NoError(err)

		key, f := binary.Varint(msg.Key)
		s.Require().False(f < 0)

		s.T().Logf(
			"headers: %v, key: %v, value: %v, partition: %v, offset: %v",
			utils.Map(msg.Headers, func(h kafka.Header) string { return fmt.Sprintf("header: %v, %v", h.Key, string(h.Value)) }),
			key,
			m,
			msg.TopicPartition.Partition,
			msg.TopicPartition.Offset,
		)

		_, err = с.CommitMessage(msg)
		s.Require().NoError(err)
	}
}
