package confluent

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type KafkaTestSuite struct {
	suite.Suite
}

func TestKafkaTestSuite(t *testing.T) {
	suite.Run(t, new(KafkaTestSuite))
}
