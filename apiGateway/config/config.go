package config

import(
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
}
