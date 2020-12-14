package persistence

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type RefreshTokenPersistence struct{}

func NewRefreshTokenPersistence() repository.RefreshTokenRepository {
	return &RefreshTokenPersistence{}
}

// リフレッシュトークンを追加
func (refreshToken RefreshTokenPersistence) AddRefreshToken(DB *gorm.DB, token string, expire *time.Time) error {

	setRefreshToken := model.RefreshToken{
		Token:  token,
		Expire: expire,
	}

	// リフレッシュトークンを作成
	err := DB.Create(&setRefreshToken).Error

	return err
}

// リフレッシュトークンをチェック
func (refreshToken RefreshTokenPersistence) CheckRefreshToken(DB *gorm.DB, token string) error {

	var setRefreshToken model.RefreshToken

	// リフレッシュトークンを作成
	err := DB.Where("token = ? and expire > ?", token, time.Now()).First(&setRefreshToken).Error

	return err
}
