package persistence_test

import (
	"testing"

	"github.com/api/domain/model"

	"gopkg.in/go-playground/assert.v1"
)

func TestGoodAdd(t *testing.T) {
	isGood := true
	var goods []model.Good

	// いいねを追加
	err := goodPersistence.AddGood(db, 1, 2, &isGood)

	db.Find(&goods)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(goods), 1)
}

func TestCheckGood(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// いいねがあるかチェック
		isGood := goodPersistence.CheckGood(db, 1, 2)

		assert.Equal(t, isGood, true)
	})

	t.Run("error postID", func(t *testing.T) {
		// いいねがあるかチェック
		isGood := goodPersistence.CheckGood(db, 1, 20000)

		assert.Equal(t, isGood, false)
	})

	t.Run("error userID", func(t *testing.T) {
		// いいねがあるかチェック
		isGood := goodPersistence.CheckGood(db, 100000000, 2)

		assert.Equal(t, isGood, false)
	})

	t.Run("error userID postID", func(t *testing.T) {
		// いいねがあるかチェック
		isGood := goodPersistence.CheckGood(db, 100000000, 20000000)

		assert.Equal(t, isGood, false)
	})
}

// いいねが更新されているかチェック
func TestUpdateGood(t *testing.T) {
	var good model.Good
	isGood := false

	// goodを更新
	err := goodPersistence.UpdateGood(db, 1, 2, &isGood)

	db.First(&good)

	assert.Equal(t, err, nil)
	assert.Equal(t, good.IsGood, isGood)
}
