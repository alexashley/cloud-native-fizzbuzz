package server

import "net/http"

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
}

func Internal(w http.ResponseWriter) {
	w.WriteHeader(500)
}
