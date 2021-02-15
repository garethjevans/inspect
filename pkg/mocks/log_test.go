package mocks_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockLogger(t *testing.T) {
	logger := mocks.LoggerMock{}

	logger.Println("message1")
	logger.Println("message2")
	logger.Println("message3")

	assert.Equal(t, 3, len(logger.Messages))
}
