package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/infra/database/pg/user"
	service "todolist/internal/service/user"
	"todolist/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestCreatesWithSucess(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))
	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	body.Write([]byte("{\"name\": \"pedro\", \"email\": \"test_create_http@email.com\", \"password\": \"123\"}"))

	request, _ := http.NewRequest(http.MethodPost, "/api/user", &body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	got := response.Body.String()
	assert.Empty(t, got)
	assert.Equal(t, response.Code, http.StatusCreated)

	t.Cleanup(func() {
		userDb.Delete("test_create_http@email.com")
	})
}

func TestNotCreateWithEmptyPayload(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))
	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	body.Write([]byte(""))

	request, _ := http.NewRequest(http.MethodPost, "/api/user", &body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	got := response.Body.String()
	assert.Empty(t, got)
	assert.Equal(t, response.Code, http.StatusBadRequest)
}

func TestNotCreateWithInvalidPassword(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))
	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	pass := "abcdefghijklmnopqrstuvwxy1zABCD1EFGHIJKLMNOPQRSTU3VWXYZ1!@#a3s1d!@#459871"
	payload := fmt.Sprintf("{\"name\": \"pedro\", \"email\": \"test_create_http@email.com\", \"password\": \"%s\"}", pass)
	body.Write([]byte(payload))

	request, _ := http.NewRequest(http.MethodPost, "/api/user", &body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	got := response.Body.String()
	assert.Equal(t, got, "{\"message\":\"Invalid data\"}\n")
	assert.Equal(t, response.Code, http.StatusBadRequest)
}

func TestNotCreateWithErrorOnService(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))
	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)
	userDb.DB.Close()

	body := bytes.Buffer{}
	body.Write([]byte("{\"name\": \"pedro\", \"email\": \"test_create_http@email.com\", \"password\": \"123\"}"))

	request, _ := http.NewRequest(http.MethodPost, "/api/user", &body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	got := response.Body.String()
	assert.Equal(t, got, "{\"message\":\"Server error, please try again\"}\n")
	assert.Equal(t, response.Code, http.StatusInternalServerError)
}
