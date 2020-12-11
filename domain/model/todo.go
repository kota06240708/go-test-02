package model

import (
	"time"
)

// Todoに関するデータ構造

type Todo struct {
	Id        uint
	Title     string
	Author    string
	CreatedAt time.Time
}
