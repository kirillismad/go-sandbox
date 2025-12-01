//go:build nats

package nats

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/suite"
)

type message struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type NatsTestSuite struct {
	suite.Suite
}

func TestNatsTestSuite(t *testing.T) {
	suite.Run(t, new(NatsTestSuite))
}

func (s *NatsTestSuite) marshalMsg(m message) []byte {
	data, err := json.Marshal(m)
	s.Require().NoError(err)
	return data
}

func (s *NatsTestSuite) unmarshalMsg(data []byte) message {
	var m message
	err := json.Unmarshal(data, &m)
	s.Require().NoError(err)
	return m
}

func (s *NatsTestSuite) TestPubSub() {
	const subjectName = "subject.name"
	conn, err := nats.Connect("nats://nats:4222")
	s.Require().NoError(err)
	defer conn.Drain()

	for i := range 2 {
		_, err := conn.Subscribe(subjectName, func(msg *nats.Msg) {
			m := s.unmarshalMsg(msg.Data)
			s.T().Logf("Received message: sub: %d %+v", i, m)
		})
		s.Require().NoError(err)
	}

	for range 3 {
		data := s.marshalMsg(message{
			ID:   gofakeit.Int64(),
			Name: gofakeit.Name(),
		})

		err = conn.Publish(subjectName, data)
		s.Require().NoError(err)
	}
	time.Sleep(1 * time.Second)
}

func (s *NatsTestSuite) TestPubSub1() {
	const subjectName = "subject.name"
	conn, err := nats.Connect(
		"nats://nats:4222",
		nats.Name("client_name"),
	)
	s.Require().NoError(err)
	defer conn.Drain()

	sub, err := conn.SubscribeSync(subjectName)
	s.Require().NoError(err)

	go func() {
		for range 3 {
			conn.Publish(subjectName, s.marshalMsg(message{
				ID:   gofakeit.Int64(),
				Name: gofakeit.Name(),
			}))
		}
	}()

	next := func() (*nats.Msg, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		return sub.NextMsgWithContext(ctx)
	}

	for range 3 {
		msg, err := next()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				continue
			}
			s.Require().NoError(err)
		}

		m := s.unmarshalMsg(msg.Data)
		s.T().Logf("Received message: %+v", m)
	}
}

func (s *NatsTestSuite) TestRequestReply() {
	const subjectName = "subject.name"
	conn, err := nats.Connect("nats://nats:4222")
	s.Require().NoError(err)
	defer conn.Drain()

	_, err = conn.Subscribe(subjectName, func(msg *nats.Msg) {
		m := s.unmarshalMsg(msg.Data)
		s.Require().NoError(err)
		s.T().Logf("Received message: %+v", m)

		msg.Respond([]byte("OK"))
	})
	s.Require().NoError(err)

	for range 3 {
		m := message{
			ID:   gofakeit.Int64(),
			Name: gofakeit.Name(),
		}
		data, err := json.Marshal(m)
		s.Require().NoError(err)

		msg, err := conn.Request(subjectName, data, 1*time.Second)
		s.Require().NoError(err)

		s.T().Logf("Received response: %s", string(msg.Data))
	}
}

func (s *NatsTestSuite) TestQueue() {
	const subjectName = "subject.name"
	conn, err := nats.Connect("nats://nats:4222")
	s.Require().NoError(err)
	defer conn.Drain()

	for i := range 2 {
		_, err := conn.QueueSubscribe(subjectName, "queue", func(msg *nats.Msg) {
			m := s.unmarshalMsg(msg.Data)
			s.Require().NoError(err)
			s.T().Logf("Received message: sub: %d %+v", i, m)
		})
		s.Require().NoError(err)
	}

	for range 4 {
		data := s.marshalMsg(message{
			ID:   gofakeit.Int64(),
			Name: gofakeit.Name(),
		})

		err = conn.Publish(subjectName, data)
		s.Require().NoError(err)
	}
	time.Sleep(1 * time.Second)
}

func (s *NatsTestSuite) TestQueue1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const subjectName = "subject.name"

	conn, err := nats.Connect("nats://nats:4222")
	s.Require().NoError(err)

	for i := range 2 {
		sub, err := conn.QueueSubscribeSync(subjectName, "queue")
		s.Require().NoError(err)
		defer sub.Unsubscribe()

		next := func() (*nats.Msg, error) {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()
			return sub.NextMsgWithContext(ctx)
		}
		go func() {
			for {
				msg, err := next()
				if err != nil {
					switch {
					case errors.Is(err, context.DeadlineExceeded):
						continue
					default:
						return
					}
				}
				m := s.unmarshalMsg(msg.Data)
				msg.Respond([]byte("OK"))
				s.T().Logf("Received message: sub: %d %+v", i, m)
			}
		}()
	}
	for range 4 {
		data := s.marshalMsg(message{
			ID:   gofakeit.Int64(),
			Name: gofakeit.Name(),
		})

		msg, err := conn.Request(subjectName, data, 1*time.Second)
		s.Require().NoError(err)
		s.T().Logf("Received response: %s", string(msg.Data))
	}
	time.Sleep(1 * time.Second)
}
