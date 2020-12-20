package repository

import (
	"github.com/jinzhu/gorm"
)

type GoodRepository interface {
	AddGood(DB *gorm.DB, userId uint, postId int, isGood *bool) error
	CheckGood(DB *gorm.DB, userId uint, postId int) bool
	UpdateGood(DB *gorm.DB, userId uint, postId int, isGood *bool) error
}
