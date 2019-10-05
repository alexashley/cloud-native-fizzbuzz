package main

import (
	"github.com/alexashley/cloud-native-fizzbuzz/domain"
	"github.com/alexashley/cloud-native-fizzbuzz/server"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var query domain.Mod3Query

	if !server.ParseBody(w, r, &query) {
		return
	}

	server.WriteBody(w, r, domain.Mod3Result{
		IsDivisibleByThree: query.Value%3 == 0,
	})
}

func main() {
	server.Init("mod-3")
	server.Route("/api/v1/math/mod/3", handler)
	server.Start()
}
