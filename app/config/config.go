package config

import (
	"os"
)

type AppConfig struct {
	Port       string
	MongoDBURI string
	// Add more configuration fields as needed
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Port:       getEnv("PORT"),
		MongoDBURI: getEnv("MONGODB_URI"),
		// Initialize other config fields
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return key
}
