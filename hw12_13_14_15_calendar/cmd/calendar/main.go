package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hw12_13_14_15_calendar/internal/config"
	"hw12_13_14_15_calendar/internal/logger"
	httpserver "hw12_13_14_15_calendar/internal/server/http"
	storagememory "hw12_13_14_15_calendar/internal/storage/memory"
	storagesql "hw12_13_14_15_calendar/internal/storage/sql"
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

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storagememoryEnable := config.DB.Enable
	if storagememoryEnable {
		storage, err := storagesql.New(logg, config.DB.Host, config.DB.Port, config.DB.NameDB, config.DB.User, config.DB.Password)
		if err != nil {
			logg.Error(err.Error())
			storagememoryEnable = false
		}
		storage.GetEvents(time.Now(), time.Now().Add(time.Second*35))
	}
	if !storagememoryEnable {
		storage := storagememory.New()
		storage.GetEvents(time.Now(), time.Now().Add(time.Second*35))
	}

	httpServer := httpserver.NewServer(*logg, config.HTTPServer.Host, config.HTTPServer.Port)

	go func() {
		<-ctx.Done()

		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctxTimeout); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := httpServer.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		defer os.Exit(1)
	}
}

// go func() {
// 	<-ctx.Done()
// 	httpServer.Shutdown()
// 	grpcServer.Stop()
// }()

// go httpServer.Start(ctx)

// go grpcServer.Start(ctx)
// ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
// defer cancel()

// if err := server.Stop(ctxTimeout); err != nil {
// 	logg.Error("failed to stop http server: " + err.Error())
// }
//}()

// lis, err := net.Listen("tcp", "0.0.0.0:50051")
// if err == nil {
// 	log.Fatalf("failed to listen: %v", err)
// }
// grpcServer := gprc.NewServer()
// pb.RegisterCalendarServer(grpcServer, pb.CalendarServer{})

//<-ctx.Done()
