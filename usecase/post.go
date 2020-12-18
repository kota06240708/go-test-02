package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type PostUseCase interface {
	GetPosts(DB *gorm.DB) ([]*model.Post, error)
	AddPost(DB *gorm.DB, post *model.Post) (*model.Post, error)
	UpdatePost(DB *gorm.DB, post *model.Post) error
	DeletePost(DB *gorm.DB, id uint) error
	GetSelectPost(DB *gorm.DB, postId uint) (*model.Post, error)
	GetUserPosts(DB *gorm.DB, userId uint) ([]*model.Post, error)
}

type postUseCase struct {
	postRepository repository.PostRepository
}

// ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewPostCase(pr repository.PostRepository) PostUseCase {
	return &postUseCase{
		postRepository: pr,
	}
}

// 投稿情報を全て取得
func (pu postUseCase) GetPosts(DB *gorm.DB) ([]*model.Post, error) {

	// DBからデータを取得
	posts, err := pu.postRepository.GetPosts(DB)

	return posts, err
}

// 投稿をアップデート
func (pu postUseCase) UpdatePost(DB *gorm.DB, p *model.Post) error {

	// DBのデータを更新
	err := pu.postRepository.UpdatePost(DB, p)

	return err
}

// 指定した投稿を削除
func (pu postUseCase) DeletePost(DB *gorm.DB, id uint) error {

	// DBのデータを削除
	err := pu.postRepository.DeletePost(DB, id)

	return err
}

// 投稿を追加
func (pu postUseCase) AddPost(DB *gorm.DB, post *model.Post) (*model.Post, error) {

	// DBからデータを取得
	post, err := pu.postRepository.AddPost(DB, post)

	return post, err
}

// 指定したIDの投稿を取得
func (pu postUseCase) GetSelectPost(DB *gorm.DB, id uint) (*model.Post, error) {

	// DBからデータを取得
	post, err := pu.postRepository.GetSelectPost(DB, id)

	return post, err
}

// 指定したユーザーの投稿を全て取得
func (pu postUseCase) GetUserPosts(DB *gorm.DB, id uint) ([]*model.Post, error) {

	// DBからデータを取得
	posts, err := pu.postRepository.GetUserPosts(DB, id)

	return posts, err
}
