package model

import "github.com/jinzhu/gorm"

// Todoに関するデータ構造

type Post struct {
	Model

	UserId uint    `json:"userId" binding:"-"`
	Text   string  `json:"text" binding:"-"`
	User   *User   `json:"user" binding:"-"`
	Goods  []*Good `json:"goods" binding:"-"`
}

type PostRes struct {
	Post

	GoodCount *int `json:"goodCount" binding:"-"`
}

func (Post) TableName() string {
	return "posts"
}

func (Post) GetGoodCount(DB *gorm.DB) *gorm.DB {
	return DB.Joins("left join goods on goods.post_id = posts.id").
		Group("posts.id").
		Select("count(goods.id) as good_count, posts.id, posts.created_at, posts.updated_at, posts.user_id, posts.text")
}
