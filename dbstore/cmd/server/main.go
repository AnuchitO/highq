package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anuchito/dbstore/pkg/config"
	"github.com/anuchito/dbstore/pkg/db"
)

func main() {
	config := config.NewConfig()
	fmt.Println("start app...")
	handler := db.NewMainHandler(config.Filename)

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("App started.")

	killSignal := <-signals
	switch killSignal {
	case os.Interrupt:
		fmt.Println("Got SIGINT...")
	case syscall.SIGTERM:
		fmt.Println("got SIGTERM...")
	}
	fmt.Println("App is shutting down...")
	err := srv.Shutdown(context.Background())
	if err != nil {
		fmt.Printf("Error shutting down: %v\n", err)
	}
	fmt.Println("Bye")
}
