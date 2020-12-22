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
func (pp PostPersistence) GetPosts(DB *gorm.DB) ([]*model.Post, error) {

	var posts []*model.Post
	user := &model.User{}

	// ユーザー全て取得
	err := DB.Preload("Goods").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(user.GetResParam)
	}).Find(&posts).Error

	return posts, err
}

// 投稿を追加
func (pp PostPersistence) AddPost(DB *gorm.DB, p *model.Post) (*model.Post, error) {

	// 投稿を追加
	err := DB.Create(&p).Error

	return p, err
}

// 指定した投稿IDを取得
func (pp PostPersistence) GetSelectPost(DB *gorm.DB, postId uint) (*model.Post, error) {
	p := &model.Post{}

	// idで絞り込む
	err := DB.Where("id = ?", postId).Preload("Goods").First(&p).Error

	return p, err
}

func (pp PostPersistence) GetUserPosts(DB *gorm.DB, userId uint) ([]*model.Post, error) {
	var posts []*model.Post
	post := &model.Post{}

	err := DB.Table(post.TableName()).
		Where("posts.user_id = ?", userId).
		Joins("left join goods on posts.id = goods.post_id").
		Group("posts.id").
		Select("count(goods.id) as good_count, posts.id, posts.created_at, posts.updated_at, posts.user_id, posts.text").
		Preload("Goods").
		Find(&posts).
		Error

	// for _, post := range posts {
	// 	count := len(post.Goods)
	// 	post.GoodCount = &count
	// }

	return posts, err
}

// 投稿をアップデート
func (pp PostPersistence) UpdatePost(DB *gorm.DB, p *model.Post) error {
	// 投稿を更新
	err := DB.Save(&p).Error

	return err
}

// 指定した投稿を削除
func (pp PostPersistence) DeletePost(DB *gorm.DB, id uint) error {
	// 指定したコメントを削除
	err := DB.Where("id = ?", id).Delete(&model.Post{}).Error

	return err
}
