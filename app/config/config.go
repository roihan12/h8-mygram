package config

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
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


// New creates a new container instance
func NewEnv() (*AppConfig, error) {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := AppConfig{}

	app.JWT_KEY = os.Getenv("JWT_KEY")
	app.DBUser = os.Getenv("DBUSER")
	app.DBPass = os.Getenv("DBPASS")
	app.DBHost = os.Getenv("DBHOST")
	app.DBName = os.Getenv("DBNAME")
	app.Port = os.Getenv("HTTP_PORT")
	app.URL = os.Getenv("HTTP_URL")
	app.Env = os.Getenv("APP_ENV")
	app.AllowedOrigins = os.Getenv("HTTP_ALLOWED_ORIGINS")
	app.CLOUDINARY_CLOUD_NAME = os.Getenv("CLOUDINARY_CLOUD_NAME")
	app.CLOUDINARY_API_KEY = os.Getenv("CLOUDINARY_API_KEY")
	app.CLOUDINARY_API_SECRET = os.Getenv("CLOUDINARY_API_SECRET")
	app.CLOUDINARY_UPLOAD_FOLDER = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	if portStr := os.Getenv("DBPORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			app.DBPort = port
		} else {
			log.Println("error converting DBPORT to int:", err)
		}
	}

	return &app, nil
	
}