package models

type PageBean[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}
