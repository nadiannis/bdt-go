package utils

import (
	"consumeapi/internal/domain/response"
	"log"
	"net/http"
	"strings"
)

func errorResponse(w http.ResponseWriter, status int, message any) {
	res := response.JSONErrorResponse{
		Status:  "error",
		Message: strings.ToLower(http.StatusText(status)),
		Detail:  message,
	}

	err := WriteJSON(w, status, res)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)

	message := "server encountered a problem"
	errorResponse(w, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, err error) {
	errorResponse(w, http.StatusNotFound, err.Error())
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	errorResponse(w, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, errors map[string]string) {
	errorResponse(w, http.StatusUnprocessableEntity, errors)
}
