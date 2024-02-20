package utils

import "github.com/spf13/viper"

type Config struct {
	DatabaseDestination string `mapstructure:"DATABASE_DESTINATION"`
	MailSenderHost      string `mapstructure:"MAIL_HOST"`
	MailSenderPort      int    `mapstructure:"MAIL_PORT"`
	MailSenderUsername  string `mapstructure:"MAIL_USERNAME"`
	MailSenderPassword  string `mapstructure:"MAIL_PASSWORD"`
	JwtSecretKey        string `mapstructure:"JWT_SECRET_KEY"`
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
