package route

import (
	"time"

	"github.com/PARMESHWARPANWAR/dev-tinder/bootstrap"
	"github.com/PARMESHWARPANWAR/dev-tinder/api/middleware"
	"github.com/PARMESHWARPANWAR/dev-tinder/mongo"
	"github.com/gin-gonic/gin"
)

func NewHealthCheckRouter(group *gin.RouterGroup) {
    group.GET("/health", func(c *gin.Context) {
        c.JSON(200, map[string]string{
            "status":  "healthy",
            "message": "Service is up and running",
        })
    })
}

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewHealthCheckRouter(publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)

}