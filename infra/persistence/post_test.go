package persistence_test

import (
	"testing"

	"github.com/api/domain/model"

	"gopkg.in/go-playground/assert.v1"
)

func TestPostAdd(t *testing.T) {
	post := model.Post{
		UserId: 100,
		Text:   "test",
	}

	_, err := postPersistence.AddPost(db, &post)

	getPosts := []*model.Post{}

	db.Where("user_id = ?", 100).Find(&getPosts)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(getPosts), 1)
}
