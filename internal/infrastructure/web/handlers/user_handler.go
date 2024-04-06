package handlers

import (
	"encoding/json"
	"net/http"

	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
)

type UserHandler struct {
	UseCase *user.CreateUser
}

func NewUserHandler(usecase *user.CreateUser) *UserHandler {
	return &UserHandler{
		usecase,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user user.Input

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := h.UseCase.Execute(user)
	SetResponse(w, err, output)
}
