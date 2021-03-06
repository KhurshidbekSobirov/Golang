package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	SMTPHost         string
	SMTPPort         int
	Smshost string
	SMTPUser         string
	SMTPUserPass     string
	EmailFromHeader  string
}

// Load ...
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DB", "postgres"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "khurshid"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "X"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9002"))

	c.SMTPHost = cast.ToString(getOrReturnDefault("SMTP_HOST", "smtp.gmail.com"))
	c.SMTPPort = cast.ToInt(getOrReturnDefault("SMTP_PORT", 587))
	c.SMTPUser = cast.ToString(getOrReturnDefault("SMTP_USER", "goguruh01@gmail.com"))
	c.Smshost = cast.ToString(getOrReturnDefault("SMS_HOST", "https://rest.nexmo.com/sms/json"))
	c.SMTPUserPass = cast.ToString(getOrReturnDefault("SMTP_USER_PASSWORD", "Qwertyu!op"))
	c.EmailFromHeader = cast.ToString(getOrReturnDefault("EMAIL_FROM_HEADER", "goguruh01@gmail.com"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
