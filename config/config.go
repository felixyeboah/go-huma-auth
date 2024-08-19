package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application
// It should be populated by the respective environment variables
type Config struct {
	JwtSecret        string `mapstructure:"JWT_SECRET"`
	RedisURL         string `mapstructure:"REDIS_URL"`
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
	Port             string `mapstructure:"PORT"`
	SymmetricKey     string `mapstructure:"SYMMETRIC_KEY"`
	VerificationLink string `mapstructure:"VERIFICATION_LINK"`
	ResendApiKey     string `mapstructure:"RESEND_API_KEY"`
}

// LoadConfig loads the application configuration from the environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
