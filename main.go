package main

import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/roihan12/h8-mygram/docs"
	sData "github.com/roihan12/h8-mygram/features/socialMedia/data"
	sHandler "github.com/roihan12/h8-mygram/features/socialMedia/handler"
	sService "github.com/roihan12/h8-mygram/features/socialMedia/service"

	cData "github.com/roihan12/h8-mygram/features/comment/data"
	cHandler "github.com/roihan12/h8-mygram/features/comment/handler"
	cService "github.com/roihan12/h8-mygram/features/comment/service"

	pData "github.com/roihan12/h8-mygram/features/photo/data"
	pHandler "github.com/roihan12/h8-mygram/features/photo/handler"
	pService "github.com/roihan12/h8-mygram/features/photo/service"

	uData "github.com/roihan12/h8-mygram/features/user/data"
	uHandler "github.com/roihan12/h8-mygram/features/user/handler"
	uService "github.com/roihan12/h8-mygram/features/user/service"
	"github.com/roihan12/h8-mygram/utils"

	"github.com/roihan12/h8-mygram/app/config"
	"github.com/roihan12/h8-mygram/app/database"
	"github.com/roihan12/h8-mygram/app/router"
)

// @title			MY GRAM API
// @version		1.0
// @description	This is a simple RESTful Social Media Service API written in Go using Gin web framework, PostgreSQL database
//
// @contact.name	Roihan Sori
// @contact.url
// @contact.email	roihansori12@gmail.com
//
// @license.name	MIT
// @license.url
//
// @host
// @BasePath					/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	cfg, _ := config.NewEnv()

	db := database.InitDBPostgres(*cfg)
	database.InitMigration(db)
	cloudinary := utils.NewCloudinary(cfg)

	// SETUP DOMAIN
	userData := uData.New(db)
	userService := uService.New(userData)
	userHandler := uHandler.New(userService)

	photoData := pData.New(db)
	photoService := pService.New(photoData, cloudinary)
	photoHandler := pHandler.New(photoService)

	commentData := cData.New(db)
	commentService := cService.New(commentData)
	commentHandler := cHandler.New(commentService)

	socialMediaData := sData.New(db)
	socialMediaService := sService.New(socialMediaData)
	socialMediaHandler := sHandler.New(socialMediaService)
	// Init router
	router, err := router.NewRouter(
		*cfg,
		*userHandler,
		*photoHandler,
		*commentHandler,
		*socialMediaHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", cfg.URL, cfg.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
