package main

import (
	_ "TimeBookingAPI/docs"
	"TimeBookingAPI/internal/bookingModule"
	"TimeBookingAPI/internal/config"
	"TimeBookingAPI/internal/http-server/handlers"
	MWLogger "TimeBookingAPI/internal/http-server/middleware/logger"
	"TimeBookingAPI/internal/storage/PostgreSQL"
	"TimeBookingAPI/internal/userModule"
	"context"
	"github.com/go-ozzo/ozzo-log"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err) // No point in continuing if we can't load the cfg
	}

	logger := setupLogger(cfg.LogsPath)
	logger.Open()
	defer logger.Close()

	logger.Info("Config loaded successfully")
	logger.Info("Starting the application")

	storage, err := PostgreSQL.NewStorage(context.TODO(), cfg.StoragePath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	bookings := bookingModule.NewDB(storage)
	users := userModule.NewDB(storage)
	router := mux.NewRouter()
	router.Use(MWLogger.New(logger))
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	handler := handlers.NewHandler(router, bookings, users, logger)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("failed to start server: %s", err.Error())
		}
	}()

	_ = storage
	logger.Info("Server started successfully")
	<-signalChan

	logger.Info("Shutting down the server")
}

func setupLogger(loggsPath string) *log.Logger {
	logger := log.NewLogger()

	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = loggsPath
	logger.Targets = append(logger.Targets, t1, t2)

	return logger
}
