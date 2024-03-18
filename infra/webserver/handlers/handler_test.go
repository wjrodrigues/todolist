package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestResponseError(t *testing.T) {
	response := httptest.NewRecorder()

	ResponseError(response, http.StatusBadRequest, "any")

	assert.Equal(t, response.Body.String(), "{\"message\":\"any\"}\n")
	assert.Equal(t, response.Code, 400)
}

func TestResponseHeader(t *testing.T) {
	response := httptest.NewRecorder()

	ResponseHeader(response, http.StatusOK)

	assert.Equal(t, response.Body.String(), "")
	assert.Equal(t, response.Code, 200)
}
