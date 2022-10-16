package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port                   string `mapstructure:"PORT"`
	Host                   string `mapstructure:"HOST"`
	AccountClientUrl       string `mapstructure:"ACCOUNT_CLIENT_URL"`
	MarketClientUrl        string `mapstructure:"MARKET_CLIENT_URL"`
	RentalClientUrl        string `mapstructure:"RENTAL_CLIENT_URL"`
	CloudinaryCloudName    string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey       string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret    string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder string `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
	UploadMaxSize          int64  `mapstructure:"UPLOAD_MAX_SIZE"`
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
