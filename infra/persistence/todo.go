package persistence

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/api/domain/model"
	"github.com/api/domain/repository"
)

type TodoPersistence struct{}

func NewTodoPersistence() repository.TodoRepository {
	return &TodoPersistence{}
}

// 値レシーバを使ってtpがGetAllを実装
func (tp TodoPersistence) GetAll(*gin.Context) ([]*model.Todo, error) {
	// 今回はDB接続を使わずに簡単に済ませた
	todo1 := model.Todo{Id: 1, Title: "a", Author: "Bob", CreatedAt: time.Now().Add(-24 * time.Hour)}
	todo2 := model.Todo{Id: 2, Title: "b", Author: "Alisa", CreatedAt: time.Now().Add(-24 * time.Hour)}

	return []*model.Todo{&todo1, &todo2}, nil
}
