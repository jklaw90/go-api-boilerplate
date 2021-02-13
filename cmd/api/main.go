package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jklaw90/go-api-boilerplate/pkg/http/handlers"
	"github.com/jklaw90/go-api-boilerplate/pkg/http/middelwares"
	"go.uber.org/zap"
)

func main() {

	l, _ := zap.NewDevelopment()
	defer l.Sync()
	logger := l.Sugar()

	logger.Infow("initalizing boilerplate-api")

	r := mux.NewRouter()
	r.Use(middelwares.LoggingMiddleware(logger))
	r.HandleFunc("/health", handlers.HealthCheck)

	srv := &http.Server{
		Handler: r,
		Addr:    ":80",
	}

	log.Fatal(srv.ListenAndServe())
}
