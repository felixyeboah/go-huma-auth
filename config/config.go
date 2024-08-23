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
func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app") // Ensure this matches the config file name without the extension
	viper.SetConfigType("env") // Adjust this if you're using a different format (yaml, json, etc.)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
