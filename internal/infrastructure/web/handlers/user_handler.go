package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
)

type UserHandler struct {
	createUser *user.CreateUser
	findUser   *find.FindUser
}

func NewUserHandler(createUser *user.CreateUser, findUser *find.FindUser) *UserHandler {
	return &UserHandler{
		createUser: createUser,
		findUser:   findUser,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user user.Input

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := h.createUser.Execute(user)
	SetResponse(w, err, output)
}

func (h *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	RequestId := r.PathValue("id")
	id, err := strconv.Atoi(RequestId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.findUser.Execute(find.Input{ID: int64(id)})
	SetResponse(w, err, output)
}
