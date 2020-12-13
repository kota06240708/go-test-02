package usecase

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/api/domain/repository"
)

type RefreshTokenUseCase interface {
	AddRefreshToken(DB *gorm.DB, token string, expire *time.Time) error
}

type refreshTokenUseCase struct {
	refreshTokenRepository repository.RefreshTokenRepository
}

// ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewRefreshTokenCase(rt repository.RefreshTokenRepository) RefreshTokenUseCase {
	return &refreshTokenUseCase{
		refreshTokenRepository: rt,
	}
}

// refreshTokenデータを追加するユースケース
func (rt refreshTokenUseCase) AddRefreshToken(DB *gorm.DB, token string, expire *time.Time) error {
	// DBにデータを追加
	err := rt.refreshTokenRepository.AddRefreshToken(DB, token, expire)

	return err
}
