package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
	"github.com/api/util"
)

type UserUseCase interface {
	GetUserAll(*gorm.DB) ([]*model.User, error)
	GetCurrentUser(DB *gorm.DB, password string, email string) (*model.User, error)
	GetCurrentUserID(DB *gorm.DB, ID int) (*model.User, error)
	AddUser(DB *gorm.DB, name string, age int, icon string, password string, email string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewUserseCase(uu repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: uu,
	}
}

// userデータを全件取得するためのユースケース
func (uu userUseCase) GetUserAll(DB *gorm.DB) ([]*model.User, error) {
	// DBからデータを全て取得
	users, err := uu.userRepository.GetAll(DB)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// ログインユーザー情報を取得
func (uu userUseCase) GetCurrentUser(DB *gorm.DB, password string, email string) (*model.User, error) {
	// DBからデータを取得
	user, err := uu.userRepository.GetCurrentUser(DB, email)

	// DBからの取得でエラーが出たら弾く
	if err != nil {
		return nil, err
	}

	// パスワードがあってるか確認
	if errPass := util.CompareHashAndPassword(user.Password, password); errPass != nil {
		return nil, errPass
	}

	// パスワードを空にする。
	user.Password = ""

	return user, nil
}

// IDでユーザー情報を取得
func (uu userUseCase) GetCurrentUserID(DB *gorm.DB, ID int) (*model.User, error) {
	// DBからデータを取得
	user, err := uu.userRepository.GetCurrentUserID(DB, ID)

	// DBからの取得でエラーが出たら弾く
	if err != nil {
		return nil, err
	}

	return user, nil
}

// userデータを追加するユースケース
func (uu userUseCase) AddUser(DB *gorm.DB, name string, age int, icon string, password string, email string) error {
	// DBにデータを追加
	err := uu.userRepository.AddUser(DB, name, age, icon, password, email)

	return err
}
