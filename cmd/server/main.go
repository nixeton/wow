package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wow/config"
	"wow/internal/data/repository"
	"wow/internal/service"
	"wow/internal/usecase"
	"wow/pkg/logger"
	"wow/pkg/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	Run(cfg)
}

func Run(config *config.Config) {
	l := logger.New(config.Log.Level)

	repo := repository.NewRepository()

	powUsecase := usecase.NewPowUsecase(config)
	powService := service.NewPowService(powUsecase)

	quoteUsecase := usecase.NewQuoteUsecase(repo)
	quoteService := service.NewQuoteService(quoteUsecase)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case s := <-interrupt:
			l.Info("interrupt signal: " + s.String())
			cancel()
		}
	}()

	srv := server.NewServer(config, powService, quoteService, l)
	err := srv.Start(ctx)
	if err != nil {
		l.Fatal(err)
	}
}
