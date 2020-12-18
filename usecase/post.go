package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type PostUseCase interface {
	GetPosts(DB *gorm.DB) ([]*model.Post, error)
	AddPost(DB *gorm.DB, p *model.Post) (*model.Post, error)
	GetSelectPost(DB *gorm.DB, id uint) ([]*model.Post, error)
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

// 投稿を追加
func (pu postUseCase) AddPost(DB *gorm.DB, p *model.Post) (*model.Post, error) {

	// DBにデータを追加
	post, err := pu.postRepository.AddPost(DB, p)

	return post, err
}

// idで投稿を取得
func (pu postUseCase) GetSelectPost(DB *gorm.DB, id uint) ([]*model.Post, error) {

	// DBにデータを追加
	posts, err := pu.postRepository.GetSelectPost(DB, id)

	return posts, err
}
