package middleware

import (
	"net/http"

	"github.com/PARMESHWARPANWAR/dev-tinder/domain"
	"github.com/PARMESHWARPANWAR/dev-tinder/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("access_token")
		if err != nil {
            c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "No token found"})
            c.Abort()
            return
        }
		authorized, err := tokenutil.IsAuthorized(authToken, secret)
		if authorized {
				userId, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id",userId)
				c.Next()
				return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		c.Abort()

		
	}
}