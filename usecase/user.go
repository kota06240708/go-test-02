package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type UserUseCase interface {
	GetUserAll(*gorm.DB) ([]*model.User, error)
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

// userデータを追加するユースケース
func (uu userUseCase) AddUser(DB *gorm.DB, name string, age int, icon string, password string, email string) (err error) {
	// DBにデータを追加
	err = uu.userRepository.AddUser(DB, name, age, icon, password, email)

	return err
}
