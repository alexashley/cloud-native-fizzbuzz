package server

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var server http.Server
var mux *http.ServeMux
var port string

func Init(app string) {
	initLogger(app)
	go initSigHandler()
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	mux = http.NewServeMux()
	server = http.Server{Addr: port, Handler: mux}
}

func Route(pattern string, handler http.HandlerFunc) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), "correlationId", id.String()))
		r = r.WithContext(context.WithValue(r.Context(), "path", r.URL.Path))

		Info(r.Context(), "start")
		start := time.Now()
		handler(w, r)
		end := time.Now()
		diff := end.Sub(start)
		Infof(r.Context(), "end (elapsed %s)", diff.String())
	})
}

func Start() {
	Infof(context.Background(), "Starting server on %s", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Fatalf(context.Background(), "Shutdown error %v", err)
	}
}

func initSigHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := <-c

	Infof(context.Background(), "Got signal %s, stopping server", s.String())

	if err := server.Shutdown(context.Background()); err != nil {
		Fatalf(context.Background(), "Error occurred while stopping server %v", err)
	}

	Info(context.Background(), "Finished server shutdown")
	os.Exit(0)
}
