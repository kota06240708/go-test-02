package model

import "time"

type RefreshToken struct {
	Model

	Token  string     `json:"token" binding:"-"`
	Expire *time.Time `json:"user" binding:"-"`
}

func (r *RefreshToken) TableName() string {
	return "refresh_tokens"
}
