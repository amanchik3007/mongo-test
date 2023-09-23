package config

import (
	"os"
)

type AppConfig struct {
	Port                string
	MongoDBURI          string
	DbName              string
	MongoUserCollection string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Port:                getEnv("PORT"),
		MongoDBURI:          getEnv("MONGODB_URI"),
		DbName:              getEnv("MONGODB_NAME"),
		MongoUserCollection: getEnv("MONGO_USER_COLLECTION"),
	}
}

func getEnv(key string) string {
	var (
		value string
		ok    bool
	)
	if value, ok = os.LookupEnv(key); ok {
		return value
	}
	return os.Getenv(value)
}
