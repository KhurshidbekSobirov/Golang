package config

import(
	"os"
	"github.com/spf13/cast"
)

type Config struct {
	Enivorentment   string
	TaskServiceHost string
	TaskServicePort int

	CtxTimeout int

	LogLevel string
	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.TaskServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.TaskServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 50321))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

