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
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrRuturnDefault("ENVIRONMENT","develop"))
	c.PostgresHost = cast.ToString(getOrRuturnDefault("POSTGRES_HOST","localhost"))
	c.PostgresPort = cast.ToInt(getOrRuturnDefault("POSTGRES_PORT",5432))
	c.PostgresDatabase = cast.ToString(getOrRuturnDefault("POSTGRES_DATABASE","tasks"))
	c.PostgresUser = cast.ToString(getOrRuturnDefault("POSTGRES_USER","khurshid"))
	c.PostgresPassword = cast.ToString(getOrRuturnDefault("POSTGRES_PASSWORD","X"))

	c.LogLevel =  cast.ToString(getOrRuturnDefault("LOG_LEVEL","debug"))
	c.RPCPort = cast.ToString(getOrRuturnDefault("RPC_PORT",":9000"))

	return c
}


func getOrRuturnDefault(key string, defaultValue interface{}) interface{}{
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
