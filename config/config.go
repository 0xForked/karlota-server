package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct{}

func init() {
	loadEnv()
}

func loadEnv() Config {
	log.Println("Load configuration file . . . .")
	// find environment file
	viper.SetConfigFile(`.env`)
	// error handling for specific case
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(".env file not found!, please copy .env.example and paste as .env")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
	log.Println("configuration file: ready")

	return Config{}
}

func (cfg Config) GetAppName() string {
	return viper.GetString("APP_NAME")
}

func (cfg Config) GetAppDesc() string {
	return viper.GetString("APP_DESC")
}

func (cfg Config) GetAppDebug() bool {
	return viper.GetBool("APP_DEBUG")
}

func (cfg Config) GetAppVersion() string {
	return viper.GetString("APP_VERSION")
}

func (cfg Config) GetAppUrl() string {
	return viper.GetString("APP_URL")
}

func (cfg Config) GetDbDriver() string {
	return viper.GetString("DB_DRIVER")
}

func (cfg Config) GetDbDsnUrl() string {
	return viper.GetString("DB_DSN_URL")
}

func (cfg Config) GetJWTSecretKey() string {
	return viper.GetString("JWT_SECRET_KEY")
}

func (cfg Config) GetJWTLifespan() int64 {
	return viper.GetInt64("JWT_LIFETIME")
}
