package model

// Todoに関するデータ構造

type User struct {
	Model

	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Icon     string  `json:"icon"`
	Email    string  `json:"email"`
	Password string  `json:"password" gorm:"-" sql:"-"`
	Goods    []*Good `json:"goods" binding:"-"`
	Posts    []*Post `json:"posts" binding:"-"`
}

func (u *User) TableName() string {
	return "users"
}
