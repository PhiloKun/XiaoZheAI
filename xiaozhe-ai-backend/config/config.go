package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")

	// Enable environment variable override
	viper.AutomaticEnv()
	viper.SetEnvPrefix("XIAOZHE")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read config file: %s", err.Error()))
	}

	log.Printf("Using config file: %s", viper.ConfigFileUsed())

	// Unmarshal to struct
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal config: %s", err.Error()))
	}

	// Validate required fields
	if err := validateConfig(); err != nil {
		panic(fmt.Sprintf("Config validation failed: %s", err.Error()))
	}

	log.Println("Configuration loaded successfully")
}

func validateConfig() error {
	if AppConfig.Server.Port == "" {
		return fmt.Errorf("server.port is required")
	}
	if AppConfig.Database.Host == "" {
		return fmt.Errorf("database.host is required")
	}
	if AppConfig.Database.Database == "" {
		return fmt.Errorf("database.database is required")
	}
	return nil
}

// Helper functions for easy access
func GetServerAddress() string {
	return fmt.Sprintf("%s:%s", AppConfig.Server.Host, AppConfig.Server.Port)
}

func GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Database.Username,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Database,
	)
}
