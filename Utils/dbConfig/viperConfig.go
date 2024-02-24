package dbConfig

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ViperConfigDev() *DevConfig {
	// read configs from file
	devConfig := DevConfig{}
	dbCreds := DbConfig{}
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	devConfig.Env = viper.GetString("DEV.ENV")
	devConfig.Port = viper.GetString("DEV.PORT")
	devConfig.UserContextKey = viper.GetString("DEV.USER_CONTEXT_KEY")
	devConfig.JwtSecretKey = viper.GetString("DEV.JWT_SECRET_KEY")
	dbCreds.DbName = viper.GetString("DEV.DATABASE_NAME")
	dbCreds.DbHost = viper.GetString("DEV.DATABASE_SERVER_HOST")
	dbCreds.DbPassword = viper.GetString("DEV.DATABASE_PASSWORD")
	dbCreds.DbPort = viper.GetInt("DEV.DATABASE_PORT")
	dbCreds.DbUserName = viper.GetString("DEV.DATABASE_USERNAME")
	devConfig.DbCreds = &dbCreds
	return &devConfig
}

func ViperConfigStage() *DevConfig {
	// read configs from file
	devConfig := DevConfig{}
	dbCreds := DbConfig{}
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	devConfig.Env = viper.GetString("STAGE.ENV")
	devConfig.Port = viper.GetString("STAGE.PORT")
	devConfig.UserContextKey = viper.GetString("STAGE.USER_CONTEXT_KEY")
	devConfig.JwtSecretKey = viper.GetString("STAGE.JWT_SECRET_KEY")
	dbCreds.DbName = viper.GetString("STAGE.DATABASE_NAME")
	dbCreds.DbHost = viper.GetString("STAGE.DATABASE_SERVER_HOST")
	dbCreds.DbPassword = viper.GetString("STAGE.DATABASE_PASSWORD")
	dbCreds.DbPort = viper.GetInt("STAGE.DATABASE_PORT")
	dbCreds.DbUserName = viper.GetString("STAGE.DATABASE_USERNAME")
	devConfig.DbCreds = &dbCreds
	return &devConfig
}

func ViperConfigLive() *DevConfig {
	// read configs from file
	devConfig := DevConfig{}
	dbCreds := DbConfig{}
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	devConfig.Env = viper.GetString("LIVE.ENV")
	devConfig.Port = viper.GetString("LIVE.PORT")
	devConfig.UserContextKey = viper.GetString("LIVE.USER_CONTEXT_KEY")
	devConfig.JwtSecretKey = viper.GetString("LIVE.JWT_SECRET_KEY")
	dbCreds.DbName = viper.GetString("LIVE.DATABASE_NAME")
	dbCreds.DbHost = viper.GetString("LIVE.DATABASE_SERVER_HOST")
	dbCreds.DbPassword = viper.GetString("LIVE.DATABASE_PASSWORD")
	dbCreds.DbPort = viper.GetInt("LIVE.DATABASE_PORT")
	dbCreds.DbUserName = viper.GetString("LIVE.DATABASE_USERNAME")
	devConfig.DbCreds = &dbCreds
	return &devConfig
}
