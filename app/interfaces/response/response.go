package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Success Status Code:200
func Success(writer http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(writer, "marshal error")
		return
	}
	_, _ = writer.Write(data)
}

// BadRequest Status Code:400
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

// Unauthorized Status Code:401
func Unauthorized(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusUnauthorized, message)
}

// StatusNotFound Status Code:404
func StatusNotFound(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusNotFound, message)
}

// InternalServerError Status Code:500
func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

// httpError Other Error Response
func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		_, _ = writer.Write(data)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
