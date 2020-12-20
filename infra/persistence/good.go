package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type GoodPersistence struct{}

func NewGoodPersistence() repository.GoodRepository {
	return &GoodPersistence{}
}

// いいねを追加
func (g GoodPersistence) AddGood(DB *gorm.DB, userId uint, postId int, isGood *bool) error {

	good := &model.Good{
		UserId: userId,
		PostId: postId,
		IsGood: isGood,
	}

	// いいねを追加
	err := DB.Create(&good).Error

	return err
}

// いいねがあるかチェック
func (g GoodPersistence) CheckGood(DB *gorm.DB, userId uint, postId int) bool {
	var good model.Good

	if err := DB.Where("user_id = ? AND post_id >= ?", userId, postId).First(&good).Error; err != nil {
		return false
	}

	return true
}

// いいねを更新
func (g GoodPersistence) UpdateGood(DB *gorm.DB, userId uint, postId int, isGood *bool) error {
	err := DB.Model(&model.Good{}).Where("user_id = ? AND post_id >= ?", userId, postId).Update("is_good", isGood).Error

	return err
}
