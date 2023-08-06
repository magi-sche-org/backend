package logger

import (
	"testing"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger, err := NewLogger(config.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, logger)
}
