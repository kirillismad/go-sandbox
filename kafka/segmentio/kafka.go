package segmentio

import (
	"context"
	"errors"
	"sync"

	"github.com/segmentio/kafka-go"
)

const topic = "segmentio-topic"

type ConsumerGroup struct {
	count int
}

func HandleMessage(m kafka.Message) error {
	return nil
}

func (g *ConsumerGroup) Start(ctx context.Context) func() {
	var wg sync.WaitGroup
	for i := 0; i < g.count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer := kafka.NewReader(kafka.ReaderConfig{
				Brokers: []string{"kafka:9092"},
				GroupID: "segmentio-group",
				Topic:   topic,
			})
			defer consumer.Close()
			for {
				m, err := consumer.FetchMessage(ctx)
				if err != nil {
					switch {
					case errors.Is(err, context.Canceled):
						return
					case errors.Is(err, context.DeadlineExceeded):
						continue

					}
				}
				err = HandleMessage(m)
				if err != nil {
					return
				}

				err = consumer.CommitMessages(ctx, m)
				if err != nil {
					return
				}
			}
		}()
	}
	return func() {
		wg.Wait()
	}
}

func (g *ConsumerGroup) Stop() {

}
