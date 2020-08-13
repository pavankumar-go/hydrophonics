package config

import (
	"github.com/spf13/viper"
)

var config Config

// LoadConfig reads configuration details from a config.toml file
func LoadConfig() error {
	viper.SetConfigName("config")       // name of config file (without extension)
	viper.SetConfigType("toml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/app/.config") // look for config in the working directory
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		return err
	}

	err = viper.Unmarshal(&config)
	return err
}

// Config contains all configuration
type Config struct {
	Database Database `toml:"database"`
}

// Database contains db configuration info
type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

// GetConfig eliminates process of reading toml each time on a call
// instead it reads toml only when config instance is nil
func GetConfig() *Config {
	if &config == nil {
		LoadConfig()
	}

	return &config
}
