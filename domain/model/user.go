package model

// Todoに関するデータ構造

type User struct {
	Model

	Name  string `json:"name" binding:"-"`
	Age   string `json:"age" binding:"-"`
	Icon  string `json:"icon" binding:"-"`
	Goods []Good `json:"icon" gorm:"many2many:goods;" binding:"-"`
	Posts []Post `json:"posts" binding:"-"`
}

func (User) TableName() string {
	return "users"
}
