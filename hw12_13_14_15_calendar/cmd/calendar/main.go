package main

import (
	"context"
	"flag"
	memorystorage "hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "hw12_13_14_15_calendar/internal/storage/sql"
	"log"
	"os/signal"
	"syscall"

	"hw12_13_14_15_calendar/internal/logger"
	"hw12_13_14_15_calendar/internal/storage"

	grpcserv "hw12_13_14_15_calendar/internal/server/grpc"
	httpserv "hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
	config, err := NewConfig(configFile)
	if err != nil {
		log.Fatalf("can't get config > %v", err)
	}
	logg := logger.New(config.Logger.File, config.Logger.Level)

	logg.Info("calendar is running ...")

	memoryStorageEn := !config.DB.Enable
	var storage storage.Storage
	if config.DB.Enable {
		storage, err = sqlstorage.New(logg, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.NameDB)
		if err != nil {
			logg.Error(err.Error())
			memoryStorageEn = true
		}
	}
	if memoryStorageEn {
		storage = memorystorage.New()
	}

	logg.Info("http server > " + config.HTTPServer.Host + ":" + config.HTTPServer.Port)
	httpServer := httpserv.NewServer(logg, config.HTTPServer.Host, config.HTTPServer.Port, storage)

	grpcServer := grpcserv.NewServer(logg, config.GRPCServer.Host, config.GRPCServer.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()
		httpServer.Shutdown()
		grpcServer.Stop()
	}()

	go httpServer.Start(ctx)

	go grpcServer.Start(ctx)

	// lis, err := net.Listen("tcp", "0.0.0.0:50051")
	// if err == nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// grpcServer := gprc.NewServer()
	// pb.RegisterCalendarServer(grpcServer, pb.CalendarServer{})

	<-ctx.Done()

}
