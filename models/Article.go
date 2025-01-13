package models

import "time"

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title" validate:"required,min=1,max=10"`
	Content    string    `json:"content" validate:"required"`
	CoverImg   string    `json:"coverImg" validate:"required,url"`
	State      string    `json:"state" validate:"required"`
	CategoryID int       `json:"categoryId" validate:"required"`
	CreateUser int       `json:"createUser,omitempty"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
	CTime      string    `json:"createTime,omitempty"`
	UTime      string    `json:"updateTime,omitempty"`
}
