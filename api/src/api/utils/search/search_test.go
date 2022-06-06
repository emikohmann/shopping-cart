package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeywords(t *testing.T) {
	assert.EqualValues(t, "limit", keywordLimit)
	assert.EqualValues(t, "offset", keywordOffset)
	assert.EqualValues(t, ",", valuesSeparator)
}
