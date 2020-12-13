package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type UserPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}

// ユーザー情報全て取得
func (user UserPersistence) GetAll(DB *gorm.DB) ([]*model.User, error) {

	var users []*model.User
	// good := model.Good{}
	// post := model.Post{}

	// ユーザー全て取得
	err := DB.Select("name, age, icon, email").Preload("Posts").Find(&users).Error

	return users, err
}

// ユーザー情報登録
func (user UserPersistence) AddUser(DB *gorm.DB, name string, age int, icon string, password string, email string) error {

	setUser := model.User{
		Name:     name,
		Age:      age,
		Icon:     icon,
		Password: password,
		Email:    email,
	}

	// ユーザー情報を登録
	err := DB.Create(&setUser).Error

	return err
}
