package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.New(os.Stdout, "[mod-3] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile)
var server http.Server

type mod3Query struct {
	Value int `json:"value"`
}

type mod3Result struct {
	IsDivisibleByThree bool `json:"divisibleByThree"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	var query mod3Query

	if err = json.Unmarshal(body, &query); err != nil {
		w.WriteHeader(400)
		return
	}

	mod3 := mod3Result{
		IsDivisibleByThree: query.Value%3 == 0,
	}

	res, err := json.Marshal(mod3)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if written, err := w.Write(res); err != nil || written != len(res) {
		w.WriteHeader(500)
		return
	}
}

func initSigHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := <-c

	logger.Printf("Got signal %s, stopping server", s.String())

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("Error occurred while stopping server %v", err)
	}

	logger.Printf("Finished server shutdown")
	os.Exit(0)
}

func main() {
	go initSigHandler()
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	m := http.NewServeMux()
	server = http.Server{Addr: port, Handler: m}

	m.HandleFunc("/api/v1/math/mod/3", handler)

	logger.Printf("Starting server on %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}
