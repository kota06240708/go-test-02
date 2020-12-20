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
func (rp RefreshTokenPersistence) AddRefreshToken(DB *gorm.DB, token string, expire *time.Time) error {

	refreshToken := &model.RefreshToken{
		Token:  token,
		Expire: expire,
	}

	// リフレッシュトークンを作成
	err := DB.Create(&refreshToken).Error

	return err
}

// リフレッシュトークンをチェック
func (rp RefreshTokenPersistence) CheckRefreshToken(DB *gorm.DB, token string) error {

	refreshToken := &model.RefreshToken{}

	// リフレッシュトークンをチェック
	err := DB.Where("token = ? and expire > ?", token, time.Now()).First(&refreshToken).Error

	return err
}
