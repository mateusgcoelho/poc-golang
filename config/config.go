package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitDefaultConfig() {
	err := godotenv.Load()
	logger := NewLogger("enviroment-config")
	
	if err != nil {
		logger.Errorf("unable to initialize environment variables: %v", err)
		os.Exit(1);
		return
	}
}
