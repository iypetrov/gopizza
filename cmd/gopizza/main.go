package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/router"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	configs.Init()
	db, err := configs.CreateDatabaseConnection()
	if err != nil {
		fmt.Printf("cannot connect to database %s", err.Error())
	}
	queries := configs.NewDatabase(db)
	if err := configs.RunSchemaMigration(db); err != nil {
		fmt.Printf("cannot run schema migration %s", err.Error())
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Printf("cannot load aws config %s", err.Error())
	}

	s3Client := s3.NewFromConfig(awsCfg)
	cognitoClient := cip.NewFromConfig(awsCfg)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", configs.Get().App.Port),
		Handler:      router.NewRouter(ctx, db, queries, s3Client, cognitoClient),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("server started on %s\n", configs.Get().App.Port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("cannot start server: %s", err.Error())
	}

	<-setupGracefulShutdown(cancel)
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			fmt.Printf("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		fmt.Printf("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			fmt.Print("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}
