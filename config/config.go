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

type JWTConfig struct {
	JwtSecret            string `mapstructure:"JWT_SECRET"`
	AccessTokenLifetime  int    `mapstructure:"ACCESS_TOKEN_LIFETIME"`
	RefreshTokenLifetime int    `mapstructure:"REFRESH_TOKEN_LIFETIME"`
}

func LoadJWTConfig(path string) (config JWTConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("jwt")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
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
