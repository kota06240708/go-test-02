package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
)

type PostRepository interface {
	GetPosts(DB *gorm.DB) ([]*model.Post, error)
	AddPost(DB *gorm.DB, post *model.Post) (*model.Post, error)
	GetUserPosts(DB *gorm.DB, userId uint) ([]*model.Post, error)
	GetSelectPost(DB *gorm.DB, commentId uint) (*model.Post, error)
	UpdatePost(DB *gorm.DB, post *model.Post) error
	DeletePost(DB *gorm.DB, id uint) error
}
