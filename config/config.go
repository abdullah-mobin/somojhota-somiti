package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

var AppLocation, _ = time.LoadLocation("Asia/Dhaka")

func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️ No .env file found (using environment variables)")
	} else {
		log.Println("✅ Loaded .env file")
	}
	return nil
}
