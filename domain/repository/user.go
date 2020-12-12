package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
)

type UserRepository interface {
	GetAll(DB *gorm.DB) ([]*model.User, error)
	// GetCurrentUser(DB *gorm.DB, c *gin.Context) (model.User, error)
	AddUser(DB *gorm.DB, name string, age int, icon string) error
}