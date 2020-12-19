package model

// Todoに関するデータ構造

type Good struct {
	Model

	UserId uint `json:"userId" binding:"-"`
	PostId int  `json:"postId" binding:"-"`
	IsGood bool `json:"isGood" binding:"-"`
}

func (g *Good) TableName() string {
	return "goods"
}
