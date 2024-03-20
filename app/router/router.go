package router

import (
	"fmt"
	"log/slog"

	// "strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/app/config"
	"github.com/roihan12/h8-mygram/app/middleware"
	photo "github.com/roihan12/h8-mygram/features/photo/handler"
	user "github.com/roihan12/h8-mygram/features/user/handler"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	app config.AppConfig,
	userHandler user.UserController,
	photoHandler photo.PhotoController,

) (*Router, error) {
	// Disable debug mode in production
	if app.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := app.AllowedOrigins
	fmt.Println(allowedOrigins)
	// originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = []string{"http://127.0.0.1:5173"}

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ROUTE
	v1 := router.Group("/v1")
	user := v1.Group("/users")
	{
		user.POST("/register", userHandler.Register)
		user.POST("/login", userHandler.Login)

		authUser := user.Group("/").Use(middleware.AuthMiddleware)
		{
			authUser.GET("", userHandler.Profile)
			authUser.PUT("", userHandler.Update)
			authUser.DELETE("", userHandler.Delete)

		}
	}

	photo := v1.Group("/photos")
	{
		authphoto := photo.Group("/").Use(middleware.AuthMiddleware)
		{
			authphoto.POST("", photoHandler.Create)
			authphoto.GET("", photoHandler.GetAll)
			authphoto.GET("/:photoId", photoHandler.GetById)
			authphoto.PUT("/:photoId", photoHandler.Update)
			authphoto.DELETE("/:photoId", photoHandler.Delete)

		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
