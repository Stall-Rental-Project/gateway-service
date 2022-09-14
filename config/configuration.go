package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port                string `mapstructure:"PORT"`
	Host                string `mapstructure:"HOST"`
	EServiceHost        string `mapstructure:"E_SERVICE_HOST"`
	EPaymentServiceHost string `mapstructure:"E_PAYMENT_SERVICE_HOST"`
	UiHost              string `mapstructure:"UI_HOST"`
	AccountClientUrl    string `mapstructure:"ACCOUNT_CLIENT_URL"`
	MarketClientUrl     string `mapstructure:"MARKET_CLIENT_URL"`
	RentalClientUrl     string `mapstructure:"RENTAL_CLIENT_URL"`
	LoggingClientUrl    string `mapstructure:"LOGGING_CLIENT_URL"`
	RateClientUrl       string `mapstructure:"RATE_CLIENT_URL"`
	RedisAddr           string `mapstructure:"REDIS_ADDR"`
	RedisPwd            string `mapstructure:"REDIS_PWD"`
	AwsRegion           string `mapstructure:"AWS_REGION"`
	AwsBucket           string `mapstructure:"AWS_BUCKET"`
	AwsContentCategory  string `mapstructure:"AWS_CONTENT_CATEGORY"`
	UploadMaxSize       int64  `mapstructure:"UPLOAD_MAX_SIZE"`
}

var Config Configuration

func (config *Configuration) LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(config)
	return
}
