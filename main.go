package main

import (
	"fmt"
	"log/slog"
	"os"

	uData "github.com/roihan12/h8-mygram/features/user/data"
	uHandler "github.com/roihan12/h8-mygram/features/user/handler"
	uService "github.com/roihan12/h8-mygram/features/user/service"

	"github.com/roihan12/h8-mygram/app/config"
	"github.com/roihan12/h8-mygram/app/database"
	"github.com/roihan12/h8-mygram/app/router"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBPostgres(*cfg)
	database.InitMigration(db)

	// SETUP DOMAIN
	userData := uData.New(db)
	userService := uService.New(userData)
	userHandler := uHandler.New(userService)

	// Init router
	router, err := router.NewRouter(
		*cfg,
		*userHandler,
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
