package middleware

import (
	"log"
	"strings"
	"todo-gin/pkg/domains/models"
	"todo-gin/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Authorization(db *gorm.DB, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("user").(*helpers.JwtCustomClaims)

		var user models.User

		err := db.Where("id = ?", userData.ID).First(&user).Error

		if err != nil {
			c.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "Unauthenticated",
				"data":       nil,
			})
			c.Abort()
			return
		}

		for _, role := range roles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		// Do something here
		c.JSON(403, gin.H{
			"statusCode": 403,
			"message":    "Forbidden",
			"data":       nil,
		})
		c.Abort()

		return

	}

}
