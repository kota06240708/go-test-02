package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
)

type PostRepository interface {
	GetPosts(DB *gorm.DB) ([]*model.PostRes, error)
	AddPost(DB *gorm.DB, post *model.Post) (*model.Post, error)
	GetUserPosts(DB *gorm.DB, userId uint) ([]*model.PostRes, error)
	GetSelectPost(DB *gorm.DB, commentId uint) (*model.PostRes, error)
	UpdatePost(DB *gorm.DB, post *model.Post, id uint) error
	DeletePost(DB *gorm.DB, id uint) error
}
