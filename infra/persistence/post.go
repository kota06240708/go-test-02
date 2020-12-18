package persistence

import (
	"github.com/api/domain/model"
	"github.com/api/domain/repository"
	"github.com/jinzhu/gorm"
)

type PostPersistence struct{}

func NewPostPersistence() repository.PostRepository {
	return &PostPersistence{}
}

// 投稿情報を全て取得
func (post PostPersistence) GetPosts(DB *gorm.DB) ([]*model.Post, error) {

	var posts []*model.Post

	// ユーザー全て取得
	err := DB.Preload("Goods").Preload("User").Find(&posts).Error

	return posts, err
}

// 投稿を追加
func (post PostPersistence) AddPost(DB *gorm.DB, p *model.Post) (*model.Post, error) {

	// 投稿を追加
	err := DB.Create(&p).Error

	return p, err
}

// 投稿を絞り込む
func (post PostPersistence) GetSelectPost(DB *gorm.DB, id uint) ([]*model.Post, error) {
	var p []*model.Post

	// idで絞り込む
	err := DB.Where("user_id = ?", id).Preload("Goods").Preload("User").Find(&p).Error

	return p, err
}
