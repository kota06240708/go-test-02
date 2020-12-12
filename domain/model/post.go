package model

// Todoに関するデータ構造

type Post struct {
	Model

	UserId int     `json:"userId" binding:"-"`
	Text   int     `json:"text" binding:"-"`
	User   User    `json:"user" binding:"-"`
	Goods  []*Good `json:"goods" gorm:"many2many:goods;" binding:"-"`
}

func (p *Post) TableName() string {
	return "posts"
}
