package controllers

import (
	"big_event/anno"
	"big_event/models"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ArticleController struct {
	web.Controller
}

func (a *ArticleController) List() {
	exec := func(a *ArticleController) (models.PageBean[models.Article], error) {
		param, err := a.Input()
		if err != nil {
			return models.PageBean[models.Article]{}, err
		}
		pageNum := param.Get("pageNum")
		pageSize := param.Get("pageSize")
		categoryId := param.Get("categoryId")
		state := param.Get("state")
		return anno.ArticleService.ListArticle(pageNum, pageSize, categoryId, state)
	}
	listArticle, err := exec(a)
	anno.PostProcess((*struct{ web.Controller })(a), err, listArticle)
}

func (a *ArticleController) Add() {
	exec := func(a *ArticleController) error {
		data := a.Ctx.Input.RequestBody
		var article models.Article
		err := json.Unmarshal(data, &article)
		if err != nil {
			return err
		}
		validate := validator.New()
		err = validate.Struct(article)

		if err != nil {
			return err
		}
		id := anno.Thread.Get()["id"].(float64)
		return anno.ArticleService.Add(article, int(id))
	}
	err := exec(a)
	anno.PostProcess((*struct{ web.Controller })(a), err, nil)
}

func (a *ArticleController) Update() {
	exec := func(a *ArticleController) error {
		data := a.Ctx.Input.RequestBody
		var article models.Article
		err := json.Unmarshal(data, &article)
		if err != nil {
			return err
		}
		validate := validator.New()
		err = validate.Struct(article)
		if err != nil {
			return err
		}
		return anno.ArticleService.Update(article)
	}

	err := exec(a)
	anno.PostProcess((*struct{ web.Controller })(a), err, nil)
}

func (a *ArticleController) Delete() {
	exec := func(a *ArticleController) error {
		param, _ := a.Input()
		id, err := strconv.Atoi(param.Get("id"))
		if err != nil {
			return err
		}
		return anno.ArticleService.Delete(id)
	}
	err := exec(a)
	anno.PostProcess((*struct{ web.Controller })(a), err, nil)
}
