package router

import (
	// "fmt"
	"log/slog"

	// "strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/app/config"
	"github.com/roihan12/h8-mygram/app/middleware"
	comment "github.com/roihan12/h8-mygram/features/comment/handler"
	photo "github.com/roihan12/h8-mygram/features/photo/handler"
	socialMedia "github.com/roihan12/h8-mygram/features/socialMedia/handler"
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
	commentHandler comment.CommentController,
	socialMediaHandler socialMedia.SocialMediaController,

) (*Router, error) {
	// Disable debug mode in production
	if app.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	// allowedOrigins := app.AllowedOrigins
	// originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowAllOrigins = true

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
		authPhoto := photo.Group("/").Use(middleware.AuthMiddleware)
		{
			authPhoto.POST("", photoHandler.Create)
			authPhoto.GET("", photoHandler.GetAll)
			authPhoto.GET("/:photoId", photoHandler.GetById)
			authPhoto.PUT("/:photoId", photoHandler.Update)
			authPhoto.DELETE("/:photoId", photoHandler.Delete)

		}
	}
	comment := v1.Group("/comments")
	{
		authComment := comment.Group("/").Use(middleware.AuthMiddleware)
		{
			authComment.POST("", commentHandler.Create)
			authComment.GET("", commentHandler.MyAllComment)
			authComment.GET("/:commentId", commentHandler.GetById)
			authComment.GET("/photos/:photoId", commentHandler.GetCommentByPhotoID)
			authComment.PUT("/:commentId", commentHandler.Update)
			authComment.DELETE("/:commentId", commentHandler.Delete)

		}
	}

	socialMedia := v1.Group("/socialmedias")
	{
		authSocialMedia := socialMedia.Group("/").Use(middleware.AuthMiddleware)
		{
			authSocialMedia.POST("", socialMediaHandler.Create)
			authSocialMedia.GET("", socialMediaHandler.GetAll)
			authSocialMedia.GET("/:socialMediaId", socialMediaHandler.GetById)
			authSocialMedia.PUT("/:socialMediaId", socialMediaHandler.Update)
			authSocialMedia.DELETE("/:socialMediaId", socialMediaHandler.Delete)

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
