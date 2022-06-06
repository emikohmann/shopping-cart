package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionDSN(t *testing.T) {
	assert.EqualValues(t, "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", connectionDSN)
}
