package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestResponse(t *testing.T) {
	response := httptest.NewRecorder()

	body := struct {
		Name string `json:"name"`
	}{Name: "any"}

	Response(response, http.StatusOK, body)

	assert.Equal(t, response.Body.String(), "{\"name\":\"any\"}\n")
	assert.Equal(t, response.Code, 200)
}

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

func TestHeader(t *testing.T) {
	response := httptest.NewRecorder()

	Header(response, "Content-Type", "application/json")

	header := response.Header()

	assert.Equal(t, header.Get("Content-Type"), "application/json")
}
