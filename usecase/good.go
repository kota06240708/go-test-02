package usecase

import (
	"github.com/jinzhu/gorm"

	"github.com/api/domain/repository"
)

type GoodUseCase interface {
	UpdateGood(DB *gorm.DB, userId uint, postId int, isGood bool) error
}

type goodUseCase struct {
	goodRepository repository.GoodRepository
}

// ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewGoodCase(gr repository.GoodRepository) GoodUseCase {
	return &goodUseCase{
		goodRepository: gr,
	}
}

// 投稿情報を全て取得
func (gu goodUseCase) UpdateGood(DB *gorm.DB, userId uint, postId int, isGood bool) error {

	// 既にDBに存在するかチェック
	getIsGood := gu.goodRepository.CheckGood(DB, userId, postId)

	var err error

	if getIsGood {
		err = gu.goodRepository.UpdateGood(DB, userId, postId, isGood)
	} else {
		err = gu.goodRepository.AddGood(DB, userId, postId, isGood)
	}

	return err
}
