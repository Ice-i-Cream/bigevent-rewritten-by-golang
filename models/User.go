package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id" validate:"required"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Nickname   string    `json:"nickname" validate:"required,min=1,max=10"`
	Email      string    `json:"email" validate:"required,email"`
	UserPic    string    `json:"userPic"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
