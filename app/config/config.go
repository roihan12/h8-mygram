package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	JWT_KEY                  string
	CLOUDINARY_CLOUD_NAME    string
	CLOUDINARY_API_KEY       string
	CLOUDINARY_API_SECRET    string
	CLOUDINARY_UPLOAD_FOLDER string
)

type AppConfig struct {
	DBUser                   string
	DBPass                   string
	DBHost                   string
	DBPort                   int
	DBName                   string
	JWT_KEY                  string
	CLOUDINARY_CLOUD_NAME    string
	CLOUDINARY_API_KEY       string
	CLOUDINARY_API_SECRET    string
	CLOUDINARY_UPLOAD_FOLDER string
	Env                      string
	URL                      string
	Port                     string
	AllowedOrigins           string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("error reading config:", err)
	}

	app.JWT_KEY = viper.GetString("JWT_KEY")
	app.DBUser = viper.GetString("DBUSER")
	app.DBPass = viper.GetString("DBPASS")
	app.DBHost = viper.GetString("DBHOST")
	app.DBName = viper.GetString("DBNAME")
	app.Port = viper.GetString("HTTP_PORT")
	app.URL = viper.GetString("HTTP_URL")
	app.Env = viper.GetString("APP_ENV")
	app.AllowedOrigins = viper.GetString("HTTP_ALLOWED_ORIGINS")
	app.CLOUDINARY_CLOUD_NAME = viper.GetString("CLOUDINARY_CLOUD_NAME")
	app.CLOUDINARY_API_KEY = viper.GetString("CLOUDINARY_API_KEY")
	app.CLOUDINARY_API_SECRET = viper.GetString("CLOUDINARY_API_SECRET")
	app.CLOUDINARY_UPLOAD_FOLDER = viper.GetString("CLOUDINARY_UPLOAD_FOLDER")

	if viper.IsSet("DBPORT") {
		app.DBPort = viper.GetInt("DBPORT")
	}

	return &app
}
