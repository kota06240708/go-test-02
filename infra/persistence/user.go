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
	good := model.Good{}
	post := model.Post{}

	// ユーザー全て取得
	err := DB.Find(&users).Preload(good.TableName()).Preload(post.TableName()).Error

	return users, err
}

// ユーザー情報登録
func (user UserPersistence) AddUser(DB *gorm.DB, name string, age int, icon string) error {

	setUser := model.User{
		Name: name,
		Age:  age,
		Icon: icon,
	}

	// ユーザー情報を登録
	err := DB.Create(&setUser).Error

	return err
}
