package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	DevelopmentEnvironment string `env:"DEVELOPMENT_ENVIRONMENT" env-default:"local"`
	StorageConfig
	HTTPServer
}

type StorageConfig struct {
	Postgresql PostgresqlConfig
}

type PostgresqlConfig struct {
	Database string `env:"POSTGRES_DB" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	User     string `env:"POSTGRES_USER" env-required:"true"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
}

type HTTPServer struct {
	Port string `env:"HTTPSERVER_PORT" env-default:"8080"`
	// TODO мб можно удалить
	Address     string        `env:"HTTPSERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"HTTPSERVER_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `env:"HTTPSERVER_IDLE_TIMEOUT" env-default:"60s"`
}

func MustLoad() *Config {
	var cfg Config

	loadEnv()

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("UNABLE TO READ ENVIRONMENT VARIABLES. SET ENVIRONMENT VARIABLES: %s", err.Error())
	}

	return &cfg
}

func loadEnv() {
	err := godotenv.Load("./service/docker/.env")
	if err != nil {
		log.Fatal("ERROR LOADING .ENV FILE")
	}
}
