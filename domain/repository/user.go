package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
)

type UserRepository interface {
	GetAll(DB *gorm.DB) ([]*model.User, error)
	GetCurrentUser(DB *gorm.DB, email string) (*model.User, error)
	AddUser(DB *gorm.DB, name string, age int, icon string, password string, email string) error
}
