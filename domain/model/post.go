package model

// Todoに関するデータ構造

type Post struct {
	Model

	UserId    uint    `json:"userId" binding:"-"`
	Text      string  `json:"text" binding:"-"`
	User      *User   `json:"user" binding:"-"`
	GoodCount *int    `json:"goodCount" binding:"-"`
	Goods     []*Good `json:"goods" binding:"-"`
}

func (p *Post) TableName() string {
	return "posts"
}
