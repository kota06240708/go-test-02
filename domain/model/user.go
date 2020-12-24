package model

import "github.com/jinzhu/gorm"

// Todoに関するデータ構造

type User struct {
	Model

	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Icon     string     `json:"icon"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Goods    []*Good    `json:"goods" binding:"-"`
	Posts    []*PostRes `json:"posts" binding:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetResParam(DB *gorm.DB) *gorm.DB {
	query := `id, name, age, icon, email, created_at, updated_at`

	return DB.Select(query)
}
