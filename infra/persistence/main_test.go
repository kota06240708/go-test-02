package persistence_test

import (
	"os"
	"testing"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
	"github.com/api/infra/persistence"
	"github.com/api/test"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var postPersistence repository.PostRepository
var goodPersistence repository.GoodRepository

func TestMain(m *testing.M) {
	// DBを定義
	db, _ = test.GetDB()

	// 依存関係を注入
	goodPersistence = persistence.NewGoodPersistence()
	postPersistence = persistence.NewPostPersistence()

	run := m.Run()

	// DBのデータを全て削除
	post := model.Post{}
	db.Unscoped().Delete(&post)

	good := model.Good{}
	db.Unscoped().Delete(&good)

	os.Exit(run)
}
