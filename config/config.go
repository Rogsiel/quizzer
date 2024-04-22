package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
    OriginHost		    string		`mapstructure:"ORIGIN_HOST"`
    DBDriver		    string		`mapstructure:"DB_DRIVER"`
    DBSource		    string		`mapstructure:"DB_SOURCE"`
    ServerAddress	    string		`mapstructure:"SERVER_ADDRESS"`
    TokenSymmetricKey	    string		`mapstructure:"TOKEN_SYMMETRIC_KEY"`
    AccessTokenDuration	    time.Duration	`mapstructure:"ACCESS_TOKEN_DURATION"`
    RefreshTokenDuration    time.Duration	`mapstructure:"REFRESH_TOKEN_DURATION"`
    SMTPAuthAddress	    string		`mapstructure:"SMTP_AUTH_ADDRESS"`
    SMTPServerAddress	    string		`mapstructure:"SMTP_SERVER_ADDRESS"`
    EmailSenderName	    string		`mapstructure:"EMAIL_SENDER_NAME"`
    EmailSenderAddress	    string		`mapstructure:"EMAIL_SENDER_ADDRESS"`
    EmailSenderPassword	    string		`mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("api")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
	return
    }

    err = viper.Unmarshal(&config)
    return
}
