package main

import (
	"assignment_4/apps"
	"assignment_4/configs"
	"assignment_4/entities/models"
	"log"
	"os"
	"strconv"
	"time"
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
