package main

import (
	"context"
	"fmt"
	"github.com/iypetrov/gopizza/internal/config"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/log"
	"github.com/iypetrov/gopizza/internal/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	db     *database.Queries
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())
	config.New()
}

func main() {
	conn, err := config.CreateDatabaseConnection(config.Get())
	if err != nil {
		log.Error("cannot connect to database %s", err.Error())
	}
	db = config.NewDatabase(conn)
	if err := config.RunSchemaMigration(conn); err != nil {
		log.Error("cannot run schema migration %s", err.Error())
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Get().App.Port),
		Handler:      router.New(ctx, db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Info("server started on %s\n", config.Get().App.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Error("cannot start server: %s", err.Error())
	}

	<-setupGracefulShutdown(cancel)
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Info("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Info("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Info("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}
