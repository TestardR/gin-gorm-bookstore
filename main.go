package main

import (
	"github.com/gin-gonic/gin"

	"github.com/TestardR/gin-gorm-bookstore/controllers"
	"github.com/TestardR/gin-gorm-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
