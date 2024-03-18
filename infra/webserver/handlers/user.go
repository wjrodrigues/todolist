package handlers

import (
	"encoding/json"
	"net/http"
	"todolist/internal/dto"
	"todolist/internal/entity/user"
	service "todolist/internal/service/user"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dtoUser dto.UserInput

	if err := json.NewDecoder(r.Body).Decode(&dtoUser); err != nil {
		ResponseHeader(w, http.StatusBadRequest)

		return
	}

	user, err := user.NewUser(dtoUser.Name, dtoUser.Email, dtoUser.Password)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, InvalidData)

		return
	}

	err = h.service.Create(user)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, InternalError)

		return
	}

	ResponseHeader(w, http.StatusCreated)
}
