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

// ユーザーを取得する時のクエリー
func getUserQuery(DB *gorm.DB) *gorm.DB {
	query := `id, name, age, icon, email, created_at, updated_at`

	return DB.Select(query)
}

// ユーザー情報全て取得
func (user UserPersistence) GetAll(DB *gorm.DB) ([]*model.User, error) {

	var users []*model.User

	// ユーザー全て取得
	err := DB.Scopes(getUserQuery).Preload("Posts").Preload("Goods").Find(&users).Error

	return users, err
}

// ログインユーザーを取得
func (user UserPersistence) GetCurrentUser(DB *gorm.DB, email string) (*model.User, error) {

	currentUser := &model.User{}

	// メールでユーザーを絞り込む
	err := DB.Preload("Posts").Preload("Goods").Where("email = ?", email).First(&currentUser).Error

	return currentUser, err
}

// IDでユーザー情報を取得
func (user UserPersistence) GetCurrentUserID(DB *gorm.DB, ID float64) (*model.User, error) {

	currentUser := &model.User{}

	// メールでユーザーを絞り込む
	err := DB.Scopes(getUserQuery).Preload("Posts").Preload("Goods").Where("id = ?", ID).First(&currentUser).Error

	return currentUser, err
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
