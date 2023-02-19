package config

import (
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBPort     uint16 `mapstructure:"POSTGRES_PORT"`
	DBName     string `mapstructure:"POSTGRES_DB"`
}

func LoadDatabaseConfig(path string) (config DatabaseConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
