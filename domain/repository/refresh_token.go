package repository

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RefreshTokenRepository interface {
	AddRefreshToken(DB *gorm.DB, token string, expire *time.Time) error
}
