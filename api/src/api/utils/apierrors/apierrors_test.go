package apierrors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestApiErrors(t *testing.T) {
	var apiErr APIError

	apiErr = NewBadRequestAPIError("test error")
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "test error", apiErr.Error())

	apiErr = NewUnauthorizedAPIError("test error")
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "test error", apiErr.Error())

	apiErr = NewNotFoundAPIError("test error")
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status())
	assert.EqualValues(t, "test error", apiErr.Error())

	apiErr = NewInternalServerAPIError("test error")
	assert.EqualValues(t, http.StatusInternalServerError, apiErr.Status())
	assert.EqualValues(t, "test error", apiErr.Error())
}
