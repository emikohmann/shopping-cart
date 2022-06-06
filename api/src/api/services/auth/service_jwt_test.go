package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultClaims(t *testing.T) {
	assert.EqualValues(t, "user_name", claimUserName)
	assert.EqualValues(t, "expires_at", claimExpiresAt)
}
