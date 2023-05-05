package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/romandnk/crud/internal/infastructure"
	"github.com/romandnk/crud/internal/infastructure/api/http"
	"github.com/romandnk/crud/internal/infastructure/repository"
	"github.com/romandnk/crud/internal/infastructure/repository/postgres"
	"github.com/romandnk/crud/internal/infastructure/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func configInit() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configInit(); err != nil {
		logrus.Fatalf("error loading config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := postgres.NewPostgresDB(ctx, connString)
	if err != nil {
		logrus.Fatalf("error connecting to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := http.NewHandler(services)

	srv := infastructure.Server{}
	go func() {
		if err := srv.Start(viper.GetString("port"), handlers.NewRouter()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("server has started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done
	logrus.Print("server is shutting down")

	if err := srv.Stop(ctx); err != nil {
		logrus.Errorf("error occured while server shutting down: %s", err.Error())
	}
	logrus.Print("server has stopped")
}
