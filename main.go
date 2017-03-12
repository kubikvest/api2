package main

import (
	"context"
	"fmt"
	"github.com/kubikvest/api2/app"
	"github.com/kubikvest/api2/games"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Start")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, _ := app.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/billing?charset=utf8&parseTime=True&loc=Local", user, pass, host, port))
	appctx := &app.Context{DB: db}

	mux := http.NewServeMux()

	games.SetHandlers(mux, appctx)

	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(1 * time.Second),
		WriteTimeout: time.Duration(1 * time.Second),
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)

	go s.ListenAndServe()

	<-sigchan

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	s.Shutdown(ctx)
	fmt.Println("Stop")
}
