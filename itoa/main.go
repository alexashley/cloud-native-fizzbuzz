package main

import (
	"github.com/alexashley/cloud-native-fizzbuzz/domain"
	"github.com/alexashley/cloud-native-fizzbuzz/server"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var req domain.ItoaRequest

	if !server.ParseBody(w, r, &req) {
		return
	}

	server.WriteBody(w, r, domain.ItoaResponse{
		String: strconv.Itoa(req.Integer),
	})
}

func main() {
	server.Init("itoa")
	server.Route("/api/v1/str/itoa", handler)
	server.Start()
}
