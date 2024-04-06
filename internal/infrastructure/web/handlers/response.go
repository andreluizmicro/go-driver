package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
)

var (
	ErrInternalServer    = "Erro interno do servidor"
	ErrUserAlreadyExists = repository.ErrUserAlreadyExists
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SetHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func SetResponse(w http.ResponseWriter, err error, bodyResponse any) {
	SetHeader(w)
	if err != nil {
		checkErrorType(w, err)
		return
	}

	setBodyResponse(w, bodyResponse)
}

func setBodyResponse(w http.ResponseWriter, body any) {
	json.NewEncoder(w).Encode(body)
}

func checkErrorType(w http.ResponseWriter, err error) {
	switch err {
	case ErrUserAlreadyExists:
		setConflictError(w, err)
	}
}

func setConflictError(w http.ResponseWriter, err error) {
	errResponse := ErrorResponse{
		Message: err.Error(),
	}
	errJSON, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintf(w, "%s", string(errJSON))
}
