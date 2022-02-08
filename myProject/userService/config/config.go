package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	TaskServicePort  int
	TaskServiceHost   string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "users"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "khurshid"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "X"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9001"))
	c.TaskServicePort = cast.ToInt(getOrReturnDefault("TASK_SERVICE_PORT", 9000))
	c.TaskServiceHost = cast.ToString(getOrReturnDefault("TASK_SERVICE_HOST", "localhost"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
