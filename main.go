package main

import (
	"net/http"
	"time"
	"github.com/kubikvest/api2/app"
	"github.com/kubikvest/api2/games"
)

func main() {
	context := &app.Context{DB:nil}

	mux := http.NewServeMux()

	games.SetHandlers(mux, context)

	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(1 * time.Second),
		WriteTimeout: time.Duration(1 * time.Second),
	}

	go s.ListenAndServe()
}
