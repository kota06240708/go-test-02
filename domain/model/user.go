package model

// Todoに関するデータ構造

type User struct {
	Model

	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Icon  string  `json:"icon"`
	Goods []*Good `json:"goods" gorm:"many2many:goods;" binding:"-"`
	Posts []*Post `json:"posts" binding:"-"`
}

func (User) TableName() string {
	return "users"
}
