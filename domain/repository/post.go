package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
)

type PostRepository interface {
	GetPosts(DB *gorm.DB) ([]*model.Post, error)
	AddPost(DB *gorm.DB, post *model.Post) (*model.Post, error)
	GetSelectPost(DB *gorm.DB, id uint) ([]*model.Post, error)
}
