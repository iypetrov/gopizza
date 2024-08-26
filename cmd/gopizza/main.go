package main

import (
	"context"
	"fmt"
	"github.com/iypetrov/gopizza/pkg/config/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	app.Init()
}

func main() {
	_, cancel := context.WithCancel(context.Background())

	app.Log.Info("server started on %s\n", app.Cfg.App.Port)
	if err := app.Server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	<-setupGracefulShutdown(cancel)
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Println("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Println("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Println("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}
