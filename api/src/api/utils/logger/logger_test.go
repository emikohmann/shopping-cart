package logger

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	Info("Test log")
	Warn("Test log")
	Error("Test log", errors.New("test error"))
	Panic("Test log", errors.New("test error"))
}
