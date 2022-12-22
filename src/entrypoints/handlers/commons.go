package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorRS struct {
	Error string `json:"error"`
}

var validate = validator.New()

func jsonError(writer http.ResponseWriter, error string, code int) {
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(&ErrorRS{
		Error: error,
	})
}
