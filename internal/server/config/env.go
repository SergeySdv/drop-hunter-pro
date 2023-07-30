package config

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	Database struct {
		Conn string
	}

	ENV string

	Config struct {
		Env           ENV
		GRPCAddr      string
		GWAddr        string
		StaticPort    string
		ProxyPort     string
		Database      Database
		App           App
		Auth          Auth
		InstanceId    string
		MigrationsDir string
		TelegramToken string
		SnapshotHost  string
		Lazanya       string

		PrometheusPort     string
		JaegerUrl          string
		JaegerServiceName  string
		PayServiceGRPCAddr string

		WebKey string
		WebIV  string
	}

	App struct {
		Domain string
		Schema string
		Port   string
	}

	Auth struct {
		GoogleClientId     string
		GoogleClientSecret string
		CookieName         string

		RedirectOnSuccess string
		RedirectOnFailure string
	}
)

const (
	Dev   ENV = "dev"
	Local ENV = "local"
	Prod  ENV = "prod"
)

func envFromString(s string) ENV {
	switch s {
	case "dev":
		return Dev
	case "local":
		return Local
	case "prod":
		return Prod
	}
	panic("invalid env value: " + s)
}

var CFG *Config

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	c := &Config{
		Env:        envFromString(mustenv("ENV")),
		GRPCAddr:   mustenv("GRPC_ADDR"),
		GWAddr:     mustenv("GW_ADDR"),
		StaticPort: mustenv("STATIC_PORT"),
		ProxyPort:  mustenv("PROXY_PORT"),
		Database: Database{
			Conn: mustenv("POSTGRES_CONN"),
		},
		App: App{
			Domain: mustenv("DOMAIN"),
			Schema: mustenv("SCHEMA"),
			Port:   mayend("PORT"),
		},
		Auth: Auth{
			GoogleClientId:     mustenv("GOOGLE_CLIENT_ID"),
			GoogleClientSecret: mustenv("GOOGLE_CLIENT_SECRET"),

			CookieName:        mustenv("COOKIE_NAME"),
			RedirectOnSuccess: mustenv("REDIRECT_ON_SUCCESS"),
			RedirectOnFailure: mustenv("REDIRECT_ON_FAILURE"),
		},
		InstanceId:         uuid.New().String(),
		MigrationsDir:      mustenv("MIGRATIONS_DIR"),
		TelegramToken:      mustenv("TELEGRAM_TOKEN"),
		SnapshotHost:       mustenv("SNAPSHOT_ORG_HOST"),
		Lazanya:            mustenv("LAZANYA"),
		PrometheusPort:     mustenv("PROMETHEUS_PORT"),
		JaegerUrl:          mustenv("JAEGER_URL"),
		JaegerServiceName:  mustenv("JAEGER_SERVICE_NAME"),
		PayServiceGRPCAddr: mustenv("PAY_SERVICE_GRPC_ADDR"),
		WebKey:             mustenv("WEB_KEY"),
		WebIV:              mustenv("WEB_IV"),
	}

	CFG = c

	return c, nil
}
func mayend(key string) string {
	return os.Getenv(key)
}
func mustenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is not set")
	}
	return v
}
