package main

import (
	"context"
	"fmt"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/common"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	configs.Init()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", configs.Get().App.Port),
		Handler:      common.NewRouter(ctx),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("server started on %s\n", configs.Get().App.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}

	<-setupGracefulShutdown(cancel)
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Printf("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Printf("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Print("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}
