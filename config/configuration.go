package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	Port                   string `mapstructure:"PORT"`
	Host                   string `mapstructure:"HOST"`
	EServiceHost           string `mapstructure:"E_SERVICE_HOST"`
	EPaymentServiceHost    string `mapstructure:"E_PAYMENT_SERVICE_HOST"`
	UiHost                 string `mapstructure:"UI_HOST"`
	AccountClientUrl       string `mapstructure:"ACCOUNT_CLIENT_URL"`
	MarketClientUrl        string `mapstructure:"MARKET_CLIENT_URL"`
	RentalClientUrl        string `mapstructure:"RENTAL_CLIENT_URL"`
	LoggingClientUrl       string `mapstructure:"LOGGING_CLIENT_URL"`
	RateClientUrl          string `mapstructure:"RATE_CLIENT_URL"`
	RedisAddr              string `mapstructure:"REDIS_ADDR"`
	RedisPwd               string `mapstructure:"REDIS_PWD"`
	AwsRegion              string `mapstructure:"AWS_REGION"`
	AwsBucket              string `mapstructure:"AWS_BUCKET"`
	AwsContentCategory     string `mapstructure:"AWS_CONTENT_CATEGORY"`
	UploadMaxSize          int64  `mapstructure:"UPLOAD_MAX_SIZE"`
	CloudinaryCloudName    string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey       string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret    string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder string `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
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

func EnvCloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config.CloudinaryCloudName
}

func EnvCloudAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config.CloudinaryApiKey
}

func EnvCloudAPISecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config.CloudinaryApiSecret
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
