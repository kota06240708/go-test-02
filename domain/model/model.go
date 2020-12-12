package model

import "time"

// Model is basic struct
type Model struct {
	ID        uint       `json:"id" binding:"-"`
	CreatedAt time.Time  `json:"createdAt" binding:"-"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"-"`
	DeletedAt *time.Time `json:"-" binding:"-"`
}
