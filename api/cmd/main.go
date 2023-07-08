package main

import (
	"fmt"
	"log"
	musixerApp "musixer/api/internal/app"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	app := musixerApp.NewAPP()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	if err := app.Listen(":7000"); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()

	fmt.Println("Running cleanup tasks...")
}
