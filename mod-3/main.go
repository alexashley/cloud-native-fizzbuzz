package main

import (
	"encoding/json"
	"github.com/alexashley/cloud-native-fizzbuzz/domain"
	"github.com/alexashley/cloud-native-fizzbuzz/server"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		server.BadRequest(w)
		server.Warnf(r.Context(), "Failed to read body: %v", err)
		return
	}

	var query domain.Mod3Query

	server.Debugf(r.Context(), "body: %s", body)

	if err = json.Unmarshal(body, &query); err != nil {
		server.Warnf(r.Context(), "Failed to parse body: %v", err)
		server.BadRequest(w)
		return
	}

	mod3 := domain.Mod3Result{
		IsDivisibleByThree: query.Value%3 == 0,
	}

	res, err := json.Marshal(mod3)

	if err != nil {
		server.Internal(w)
		return
	}

	if written, err := w.Write(res); err != nil || written != len(res) {
		server.Internal(w)
		return
	}
}

func main() {
	server.Init("mod-3")
	server.Route("/api/v1/math/mod/3", handler)
	server.Start()
}
