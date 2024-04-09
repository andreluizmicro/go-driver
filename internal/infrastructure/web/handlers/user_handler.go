package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	request "github.com/andreluizmicro/go-driver-api/internal/infrastructure/web/requests"
	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/update"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	createUser *user.CreateUser
	findUser   *find.FindUser
	updateUser *update.UpdateUser
}

func NewUserHandler(createUser *user.CreateUser, findUser *find.FindUser, updateUser *update.UpdateUser) *UserHandler {
	return &UserHandler{
		createUser: createUser,
		findUser:   findUser,
		updateUser: updateUser,
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

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	RequestId := r.PathValue("id")
	_, err := strconv.Atoi(RequestId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var input update.Input

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(RequestId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	input.ID = int64(id)

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			SetHeader(w)
			setError(w, request.ValidationErrors(e.ActualTag(), e.Field()), http.StatusUnprocessableEntity)
		}
		return
	}

	output, err := h.updateUser.Execute(input)
	SetResponse(w, err, output)
}
