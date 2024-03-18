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

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Error{Message: message})
}

func ResponseHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
