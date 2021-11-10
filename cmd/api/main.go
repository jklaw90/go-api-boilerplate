package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewDevelopment()
	defer l.Sync()
	logger := l.Sugar()

	logger.Infow("initalizing boilerplate-api")

	r := mux.NewRouter()
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("request received")
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":9000",
	}

	defer logger.Infow("shutting down boilerplate-api")
	log.Fatal(srv.ListenAndServe())
}
