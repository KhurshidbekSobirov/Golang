package config

import(
	"os"
	"github.com/spf13/cast"
)

type Config struct {
	Enivorentment   string
	TaskServiceHost string
	TaskServicePort int

	UserServiceHost string
	UserServicePort int
	CtxTimeout int

	EmailServiceHost string
	EmailServicePort int
	LogLevel string
	HTTPPort string

	SigninKey string
	RedisHost string
	RedisPort int
	
}

func Load() Config {
	c := Config{}

	c.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.TaskServiceHost = cast.ToString(getOrReturnDefault("TASK_SERVICE_HOST", "task_service"))
	c.TaskServicePort = cast.ToInt(getOrReturnDefault("TASK_SERVICE_PORT", 9000))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST","user_service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9001))
	c.EmailServiceHost = cast.ToString(getOrReturnDefault("EMAIL_SERVICE_HOST","email_service"))
	c.EmailServicePort = cast.ToInt(getOrReturnDefault("EMAIL_SERVICE_PORT", 9002))
	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "najottalimsecretkey"))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "redisDB"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	
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

