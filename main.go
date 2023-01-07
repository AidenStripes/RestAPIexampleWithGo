package main

import (
	"go/mod/controllers"
	"go/mod/models"
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors.Default())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run()
	if err != nil {
		return
	}
}
