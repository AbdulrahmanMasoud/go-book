package main

import (
	"github.com/AbdulrahmanMasoud/go-book/controllers"
	"github.com/AbdulrahmanMasoud/go-book/database"
	"github.com/AbdulrahmanMasoud/go-book/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	conn := database.Connection(db)

	defer conn.Close()
	//db.Migrator().DropTable(&models.Book{})
	db.AutoMigrate(&models.Book{})

	route := gin.Default()
	//Book Resource
	books := route.Group("/blogs")
	{
		books.GET("/", controllers.Index)
		books.GET("/:id", controllers.Show)
		books.POST("/store", controllers.Store)
		books.PUT("/:id/update", controllers.Update)
		books.DELETE("/:id/delete", controllers.Delete)
	}

	route.Run()
}
