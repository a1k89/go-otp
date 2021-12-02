package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sms/models"
)

func GetConfiguration() models.Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Can't load env")
	}

	conf := models.Configuration{
		RedisHost: os.Getenv("REDIS_HOST"),
	}

	return conf

}
