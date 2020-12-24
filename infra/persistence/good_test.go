package persistence_test

import (
	"os"
	"testing"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"

	"github.com/jinzhu/gorm"

	"github.com/api/test"
	"gopkg.in/go-playground/assert.v1"

	"github.com/api/infra/persistence"
)

var db *gorm.DB
var goodPersistence repository.GoodRepository

func TestMain(m *testing.M) {
	// DB構築
	test.InitDB("../../test/")

	// DBを定義
	db, _ = test.GetDB()

	// 依存関係を注入
	goodPersistence = persistence.NewGoodPersistence()

	run := m.Run()

	// DBのデータを全て削除
	good := model.Good{}
	db.Unscoped().Delete(&good)

	os.Exit(run)
}

func TestGoodAdd(t *testing.T) {
	isGood := true

	err := goodPersistence.AddGood(db, 1, 2, &isGood)

	assert.Equal(t, err, nil)
}
