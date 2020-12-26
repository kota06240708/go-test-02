package persistence_test

import (
	"testing"

	"github.com/api/domain/model"

	"gopkg.in/go-playground/assert.v1"
)

func TestAddPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		post := model.Post{
			UserId: 100,
			Text:   "test",
		}

		_, err := postPersistence.AddPost(db, &post)

		getPosts := []*model.Post{}

		db.Where("user_id = ?", 100).Find(&getPosts)

		assert.Equal(t, err, nil)
		assert.Equal(t, len(getPosts), 1)
	})
}

func TestGetSelectPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		post, err := postPersistence.GetSelectPost(db, 1)

		assert.Equal(t, err, nil)
		assert.Equal(t, post.Text, "test01")
	})

	t.Run("error", func(t *testing.T) {
		_, err := postPersistence.GetSelectPost(db, 10000000)

		assert.Equal(t, err != nil, true)
	})
}

func TestGetUserPosts(t *testing.T) {
	posts, err := postPersistence.GetUserPosts(db, 1)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(posts), 3)
}

func TestUpdatePost(t *testing.T) {
	p := model.Post{
		Text: "ばか",
	}

	p.ID = 1

	var post model.Post

	err := postPersistence.UpdatePost(db, &p, 1)

	db.Where("id = ?", 1).First(&post)

	assert.Equal(t, err, nil)
	assert.Equal(t, post.Text, "ばか")
}

func TestDeletePost(t *testing.T) {
	var post model.Post

	err := postPersistence.DeletePost(db, 1)

	errDB := db.Where("id = ?", 1).First(&post).Error

	assert.Equal(t, err, nil)
	assert.Equal(t, true, errDB != nil)
}
