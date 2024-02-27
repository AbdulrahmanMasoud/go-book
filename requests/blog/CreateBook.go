package requests

import (
	"gorm.io/gorm"
	"strings"
)

type CreateBookRequest struct {
	gorm.Model
	Title   string `gorm:"type:varchar(500)" json:"title" binding:"required"`
	Content string `gorm:"type:text" json:"content" binding:"required"`
	Slug    string `gorm:"type:varchar(500);unique" json:"slug"`
}

func (CreateBookRequest) TableName() string {
	return "books"
}

func (book *CreateBookRequest) BeforeCreate(tx *gorm.DB) (err error) {
	slug := strings.ToLower(strings.Replace(book.Title, " ", "-", -1))
	book.Slug = slug
	return
}
