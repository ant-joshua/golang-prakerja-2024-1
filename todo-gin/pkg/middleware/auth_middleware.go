package middleware

import (
	"log"
	"strings"
	"todo-gin/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Println("AuthMiddleware")
		// Do something here
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "Unauthenticated",
				"data":       nil,
			})
			c.Abort()

			return
		}

		split := strings.Split(authHeader, "Bearer ")

		if len(split) != 2 {
			c.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "Unauthenticated",
				"data":       nil,
			})

			c.Abort()
			return
		}

		token := split[1]

		verifyToken, err := helpers.VerifyJwtToken(token)

		if err != nil {
			c.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "Unauthenticated",
				"data":       nil,
			})
			c.Abort()
			return
		}

		c.Set("user", verifyToken)

		c.Next()
	}
}
