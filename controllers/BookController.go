package controllers

import (
	"github.com/AbdulrahmanMasoud/go-book/database"
	"github.com/AbdulrahmanMasoud/go-book/models"
	requests "github.com/AbdulrahmanMasoud/go-book/requests/book"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	db := database.Connect()
	var books []models.Book
	db.Where("visible =?", true).Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Show book by id or slug
func Show(c *gin.Context) {
	db := database.Connect()
	var book models.Book
	db.Where("id =?", c.Param("id")).Or("slug =?", c.Param("id")).First(&book)
	if book.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found this book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// ShowByUser Show blogs by user
func ShowByUser(c *gin.Context) {
	db := database.Connect()
	var book []models.Book
	db.Debug().Where("user_id =?", c.Param("user_id")).Find(&book)
	//if book.ID == 0 {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found this book"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func Store(c *gin.Context) {
	db := database.Connect()
	var book requests.CreateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := db.Create(&book)
	if created.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": created.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &book})
}

func Update(c *gin.Context) {
	db := database.Connect()
	var book requests.UpdateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := db.Where("id =?", c.Param("id")).Updates(&book)

	if updated.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": updated.Error})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "updated done"})
}

func Delete(c *gin.Context) {
	db := database.Connect()
	var book models.Book
	deleted := db.Where("id =?", c.Param("id")).Delete(&book)
	if deleted.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleted.Error})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
