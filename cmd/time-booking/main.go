package main

import (
	"TimeBookingAPI/internal/config"
	"TimeBookingAPI/internal/storage/PostgreSQL"
	"context"
	"github.com/go-ozzo/ozzo-log"
	"github.com/gorilla/mux"
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

	//bookings := booking.NewDB(storage)
	//users := user.NewDB(storage)

	router := mux.NewRouter()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
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

	//users.Create(context.TODO(), user.User{
	//	"",
	//	"John",
	//	"Doe",
	//	time.Now(),
	//	time.Now(),
	//})

	//tempUser := user.New(
	//	"ANTONina",
	//	"asdf",
	//)
	//
	//err = users.Create(context.TODO(), tempUser)
	//
	//if err != nil {
	//	logger.Error(err.Error())
	//}

	//err = users.Delete(context.TODO(), bookings, "8f1b198a-11a5-44dd-bc53-b30f6c4f7b3f")
	//err = bookings.Create(context.TODO(), &booking.Booking{
	//	"",
	//	"8f1b198a-11a5-44dd-bc53-b30f6c4f7b3f",
	//	time.Now(),
	//	time.Now().Add(time.Hour),
	//})

	//bookingsArr, err := bookings.FindAll(context.TODO(), "8f1b198a-11a5-44dd-bc53-b30f6c4f7b3f")

	//booking, err := bookings.Find(context.TODO(), "f3a0ddc0-2d0f-42d0-b939-0f46926cc44f")

	//err = bookings.Delete(context.TODO(), "f3a0ddc0-290f-42d0-b939-0f46926cc44f")

}

func setupLogger(loggsPath string) *log.Logger {
	logger := log.NewLogger()

	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = loggsPath
	logger.Targets = append(logger.Targets, t1, t2)

	return logger
}
