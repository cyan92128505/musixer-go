package main

import (
	"fmt"
	"log"
	musixerApp "musixer/api/internal/app"
	"musixer/api/internal/app/initializers"
	"os"
	"os/signal"
	"sync"
	"time"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
	initializers.ConnectRedis(&config)
}

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

	config, _ := initializers.LoadConfig(".")

	if err := app.Listen(":" + config.ServerPort); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()

	fmt.Println("Running cleanup tasks...")
}
