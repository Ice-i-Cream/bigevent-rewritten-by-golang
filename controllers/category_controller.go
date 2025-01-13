package controllers

import (
	"big_event/anno"
	"big_event/models"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type CategoryController struct {
	web.Controller
}

func (c *CategoryController) List() {
	exec := func(c *CategoryController) ([]models.Category, error) {
		id := anno.Thread.Get()["id"].(float64)
		return anno.CategoryService.ListCategory(int(id))
	}
	categoryList, err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, categoryList)
}

func (c *CategoryController) Add() {
	exec := func(c *CategoryController) error {
		var category models.Category
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &category)
		if err != nil {
			return err
		}
		validate := validator.New()
		err = validate.Struct(category)
		if err != nil {
			return err
		}
		id := anno.Thread.Get()["id"].(float64)
		return anno.CategoryService.Add(category, int(id))
	}

	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}

func (c *CategoryController) Update() {
	exec := func(c *CategoryController) error {
		var category models.Category
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &category)
		if err != nil {
			return err
		}
		validate := validator.New()
		err = validate.Struct(category)
		if err != nil {
			return err
		}
		return anno.CategoryService.Update(category)
	}

	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}

func (c *CategoryController) Delete() {
	exec := func(c *CategoryController) error {
		param, _ := c.Input()
		id, err := strconv.Atoi(param.Get("id"))
		if err != nil {
			return err
		}
		return anno.CategoryService.Delete(id)
	}

	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)

}

func (c *CategoryController) Detail() {
	exec := func(c *CategoryController) (models.Category, error) {
		param, _ := c.Input()
		id, err := strconv.Atoi(param.Get("id"))
		if err != nil {
			return models.Category{}, err
		}
		return anno.CategoryService.Detail(id)
	}
	category, err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, category)
}
