package mapper

import (
	"big_event/anno"
	"big_event/models"
	"fmt"
	"time"
)

type CategoryMapper struct{}

func (c *CategoryMapper) ListCategory(id int) ([]models.Category, error) {
	selectSQL := "SELECT id, category_name, category_alias, create_user, create_time, update_time FROM category WHERE create_user = ?"
	rows, err := anno.Db.Query(selectSQL, id)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.CategoryName, &category.CategoryAlias, &category.CreateUser, &category.CreateTime, &category.UpdateTime)
		if err != nil {
			return nil, err
		}
		category.CTime = category.CreateTime.Format("2006-01-02 15:04:05")
		category.UTime = category.UpdateTime.Format("2006-01-02 15:04:05")
		categories = append(categories, category)
	}
	return categories, rows.Err()
}

func (c *CategoryMapper) Add(category models.Category, id int) error {
	insertSQL := "insert into category (category_name, category_alias, create_user, create_time, update_time) VALUES (?,?,?,?,?)"
	_, err := anno.Db.Exec(insertSQL, category.CategoryName, category.CategoryAlias, id, time.Now(), time.Now())
	return err
}

func (c *CategoryMapper) Update(category models.Category) error {
	updateSQL := "update category set category_name = ? , category_alias=?,update_time=? where id = ?"
	_, err := anno.Db.Exec(updateSQL, category.CategoryName, category.CategoryAlias, time.Now(), category.ID)
	return err
}

func (c *CategoryMapper) Delete(id int) error {
	deleteSQL := "delete from category where id = ?"
	_, err := anno.Db.Exec(deleteSQL, id)
	return err
}

func (c *CategoryMapper) Detail(id int) (models.Category, error) {
	selectSQL := "select  * from category where id = ?"
	var category models.Category
	rows, err := anno.Db.Query(selectSQL, id)
	if err != nil {
		return category, err
	} else {
		for rows.Next() {
			err := rows.Scan(&category.ID, &category.CategoryName, &category.CategoryAlias, &category.CreateUser, &category.CreateTime, &category.UpdateTime)
			if err != nil {
				return category, err
			} else {
				category.CTime = category.CreateTime.Format("2006-01-02 15:04:05")
				category.UTime = category.UpdateTime.Format("2006-01-02 15:04:05")
				return category, nil
			}
		}
	}
	return category, fmt.Errorf("category not found")
}
