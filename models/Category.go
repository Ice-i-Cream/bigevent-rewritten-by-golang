package models

import "time"

// Category represents a category in the system.
type Category struct {
	ID            int       `json:"id"`
	CategoryName  string    `json:"categoryName" validate:"required"`
	CategoryAlias string    `json:"categoryAlias" validate:"required"`
	CreateUser    int       `json:"createUser,omitempty"`
	CreateTime    time.Time `json:"-"`
	UpdateTime    time.Time `json:"-"`
	CTime         string    `json:"createTime,omitempty" validate:"omitempty"`
	UTime         string    `json:"updateTime,omitempty" validate:"omitempty"`
}
