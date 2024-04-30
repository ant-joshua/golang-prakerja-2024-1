package main

import (
	"todo-gin/pkg/controllers"
	"todo-gin/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitDB()

	defer func() {
		dbInstance, _ := db.DB()

		err := dbInstance.Close()
		if err != nil {
			panic("failed to close database")
		}
	}()

	r := gin.Default()

	// Recovery middleware will recover any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")

	userController := controllers.NewUserController(db)
	userController.Routes(api)

	r.Run(":5000")
}
