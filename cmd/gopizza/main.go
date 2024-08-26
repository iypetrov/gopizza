package main

import (
	"context"
	"fmt"
	"github.com/iypetrov/gopizza/pkg/config"
	"github.com/iypetrov/gopizza/pkg/config/logger"
	"github.com/iypetrov/gopizza/pkg/config/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	ctx    context.Context
	cfg    *config.Config
	logger *logger.Logger
	server *http.Server
}

var (
	app *App
)

func init() {
	app = &App{}
	app.ctx = context.Background()
	app.cfg = config.New()
	app.logger = logger.New()
	app.server = server.New(app.cfg)
}

func main() {
	if err := Run(app.ctx); err != nil {
		return
	}
}

func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	fmt.Printf("server started on %s\n", app.cfg.App.Port)
	if err := app.server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	<-setupGracefulShutdown(cancel)
	return nil
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
