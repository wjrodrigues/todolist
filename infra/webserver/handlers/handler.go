package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	InternalError = "Server error, please try again"
	InvalidData   = "Invalid data"
)

type Error struct {
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, status int, value interface{}) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(value)
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Error{Message: message})
}

func ResponseHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func Header(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}
