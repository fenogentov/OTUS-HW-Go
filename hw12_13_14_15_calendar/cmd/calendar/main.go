package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/config"
	"github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/fenogentov/OTUS-HW-Go/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./internal/config/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("can't get config: %v", err)
	}

	logg := logger.New(config.Logger.File, config.Logger.Level)

	//	memStorageEn := !config.DB.Enable
	//	if config.DB.Enable {
	//	storage, err := sqlstorage.New(config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.NameDB)
	//	if err != nil {
	//		logg.Error(err.Error())
	//		memStorageEn = true
	//	}
	//	calendar := app.New(logg, storage)
	//	}
	//	if memStorageEn {
	//		storage := memorystorage.New()
	//		calendar := app.New(logg, storage)
	//	}

	server := internalhttp.NewServer(*logg, config.HTTPServer.Host, config.HTTPServer.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		defer os.Exit(1)
	}
}
