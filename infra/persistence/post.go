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

// 指定した投稿IDを取得
func (post PostPersistence) GetSelectPost(DB *gorm.DB, postId uint) (*model.Post, error) {
	p := &model.Post{}

	// idで絞り込む
	err := DB.Where("id = ?", postId).Preload("Goods").First(&p).Error

	return p, err
}

func (post PostPersistence) GetUserPosts(DB *gorm.DB, userId uint) ([]*model.Post, error) {
	var p []*model.Post

	// idで絞り込む
	err := DB.Where("user_id = ?", userId).Preload("Goods").Find(&p).Error

	return p, err
}

// 投稿をアップデート
func (post PostPersistence) UpdatePost(DB *gorm.DB, p *model.Post) error {
	// 投稿を更新
	err := DB.Save(&p).Error

	return err
}

// 指定した投稿を削除
func (post PostPersistence) DeletePost(DB *gorm.DB, id uint) error {
	// 指定したコメントを削除
	err := DB.Where("id = ?", id).Delete(&model.Post{}).Error

	return err
}
