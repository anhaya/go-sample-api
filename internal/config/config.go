package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	DB db `mapstructure:"db"`
}

type db struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

func LoadConfig() (config config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Set default values
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "goapi")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.database", "goapi")

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return
}
