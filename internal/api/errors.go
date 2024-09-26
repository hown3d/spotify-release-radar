package api

import (
	"net/http"
)

type apiError struct {
	statusCode int
	message    string
}

func (e apiError) handle(w http.ResponseWriter) {
	w.WriteHeader(e.statusCode)
	w.Write([]byte(e.message))
}
