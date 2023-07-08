package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerPort    string `mapstructure:"PORT"`
	ClientOrigin  string `mapstructure:"CLIENT_ORIGIN"`
	DBUrl         string `mapstructure:"DATABASE_URL"`
	DBName        string `mapstructure:"MONGO_INITDB_DATABASE"`
	RedisUri      string `mapstructure:"REDIS_URL"`
	MusixmatchKey string `mapstructure:"MUSIXMATCH_API_KEY"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
