package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todolist/infra/database/pg/user"
	entity "todolist/internal/entity/user"

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

func TestAuthWithSuccess(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))
	user, _ := entity.NewUser("Maria", "auth@email.com", "123")
	userDb.Create(user)

	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	body.Write([]byte("{\"email\": \"auth@email.com\", \"password\": \"123\"}"))

	request, _ := http.NewRequest(http.MethodPost, "/api/user/auth", &body)
	response := httptest.NewRecorder()

	handler.Auth(response, request, "123")

	got := response.Body.String()

	assert.True(t, strings.Contains(got, "access_token"))
	assert.Equal(t, response.Code, http.StatusOK)

	t.Cleanup(func() {
		userDb.Delete("auth@email.com")
	})
}

func TestAuthWithInvalidUser(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))

	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	body.Write([]byte("{\"email\": \"auth@email.com\", \"password\": \"123\"}"))

	request, _ := http.NewRequest(http.MethodPost, "/api/user/auth", &body)
	response := httptest.NewRecorder()

	handler.Auth(response, request, "123")

	got := response.Body.String()

	assert.Equal(t, got, "{\"message\":\"email or password are invalid\"}\n")
	assert.Equal(t, response.Code, http.StatusUnauthorized)
}

func TestAuthWithInvalidBody(t *testing.T) {
	userDb := user.NewUserDB(test.Conn(t))

	service := service.NewUserService(userDb)
	handler := NewUserHandler(service)

	body := bytes.Buffer{}
	body.Write([]byte("{"))

	request, _ := http.NewRequest(http.MethodPost, "/api/user/auth", &body)
	response := httptest.NewRecorder()

	handler.Auth(response, request, "/")

	got := response.Body.String()

	assert.Equal(t, got, "")
	assert.Equal(t, response.Code, http.StatusBadRequest)
}
