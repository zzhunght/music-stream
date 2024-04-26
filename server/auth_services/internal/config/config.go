package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	JwtSecretKey         string        `mapstructure:"JWT_SECRET_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	ServerPort           string        `mapstructure:"SERVER_ADDRESS"`
}

func LoadEnv(path string) (c *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
