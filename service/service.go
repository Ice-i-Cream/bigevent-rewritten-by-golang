package service

import (
	"big_event/models"
)

type UserService interface {
	Add(username string, password string) error
	FindByName(username string) (models.User, error)
	Update(user models.User) error
	UpdateAvatar(id int, url string) error
	UpdatePwd(data models.Password, id int) error
}

type CategoryService interface {
	ListCategory(id int) ([]models.Category, error)
	Add(category models.Category, id int) error
	Update(category models.Category) error
	Delete(id int) error
	Detail(id int) (category models.Category, err error)
}

type ArticleService interface {
	ListArticle(pageNum string, pageSize string, categoryId string, state string) (models.PageBean[models.Article], error)
	Add(article models.Article, id int) error
	Update(article models.Article) error
	Delete(id int) error
}
