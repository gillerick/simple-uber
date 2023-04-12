package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-uber/internal/auth"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	errors := make(map[string]string)
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			errors["unauthorized"] = "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errors,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
