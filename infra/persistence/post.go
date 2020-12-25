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
func (pp PostPersistence) GetPosts(DB *gorm.DB) ([]*model.PostRes, error) {

	var posts []*model.PostRes
	post := &model.Post{}
	user := &model.User{}

	// ユーザー全て取得
	err := DB.Preload("Goods").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(user.GetResParam)
		}).
		Scopes(post.GetGoodCount).
		Find(&posts).Error

	return posts, err
}

// 投稿を追加
func (pp PostPersistence) AddPost(DB *gorm.DB, p *model.Post) (*model.Post, error) {

	// 投稿を追加
	err := DB.Create(&p).Error

	return p, err
}

// 指定した投稿IDを取得
func (pp PostPersistence) GetSelectPost(DB *gorm.DB, postId uint) (*model.PostRes, error) {
	p := &model.PostRes{}
	post := &model.Post{}

	// idで絞り込む
	err := DB.
		Where("posts.id = ?", postId).
		Scopes(post.GetGoodCount).
		Preload("Goods").
		First(&p).Error

	return p, err
}

func (pp PostPersistence) GetUserPosts(DB *gorm.DB, userId uint) ([]*model.PostRes, error) {
	var posts []*model.PostRes
	post := &model.Post{}

	err := DB.Table(post.TableName()).
		Where("posts.user_id = ?", userId).
		Scopes(post.GetGoodCount).
		Preload("Goods").
		Find(&posts).
		Error

	return posts, err
}

// 投稿をアップデート
func (pp PostPersistence) UpdatePost(DB *gorm.DB, p *model.Post, id uint) error {
	// 投稿を更新
	err := DB.Model(&model.Post{}).Where("id = ?", id).Update("text", p.Text).Error

	return err
}

// 指定した投稿を削除
func (pp PostPersistence) DeletePost(DB *gorm.DB, id uint) error {
	// 指定したコメントを削除
	err := DB.Where("id = ?", id).Delete(&model.Post{}).Error

	return err
}
