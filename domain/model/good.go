package model

// Todoに関するデータ構造

type Good struct {
	Model

	UserId int `json:"userId" binding:"-"`
	PostId int `json:"postId" binding:"-"`
}

func (Post) TableName() string {
	return "goods"
}
