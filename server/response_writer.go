package server

import "net/http"

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func Internal(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
