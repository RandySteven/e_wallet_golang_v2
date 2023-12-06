package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apps"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func AppPort() string {
	return os.Getenv("APP_PORT")
}

func TimeoutDuration() time.Duration {
	timeOut, _ := strconv.Atoi(os.Getenv("TIME_OUT_DURATION"))
	var timeDuration time.Duration = time.Duration(timeOut)
	return timeDuration
}

func InitHandlers() *apps.Handlers {
	config := InitConfig()

	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = repository.Automigrate()
	if err != nil {
		return nil
	}

	handlers, err := apps.NewHandlers(*repository)
	if err != nil {
		return nil
	}

	return handlers
}
