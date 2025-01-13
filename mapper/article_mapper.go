package mapper

import (
	"big_event/anno"
	"big_event/models"
	"database/sql"
	"strconv"
	"time"
)

type ArticleMapper struct{}

func (a *ArticleMapper) ListArticle(pagenum string, pagesize string, categoryid string, state string) (models.PageBean[models.Article], error) {
	pageNum, err := strconv.Atoi(pagenum)
	if err != nil {
		return models.PageBean[models.Article]{}, err
	}
	pageSize, err := strconv.Atoi(pagesize)
	if err != nil {
		return models.PageBean[models.Article]{}, err
	}
	categoryId, err := strconv.Atoi(categoryid)
	if err != nil {
		if categoryid == "" {
			categoryId = -1
		} else {
			return models.PageBean[models.Article]{}, err
		}
	}

	selectSQL := "SELECT * FROM article"
	var args []interface{}
	whereClause := ""

	if categoryId != -1 {
		whereClause += " category_id=? AND"
		args = append(args, categoryId)
	}
	if state != "" {
		whereClause += " state=? AND"
		args = append(args, state)
	}

	if whereClause != "" {
		whereClause = whereClause[:len(whereClause)-4]
		selectSQL += " WHERE " + whereClause
	}

	var rows *sql.Rows
	rows, err = anno.Db.Query(selectSQL, args...)

	if err != nil {
		return models.PageBean[models.Article]{}, err
	}

	var PageBean models.PageBean[models.Article]
	var order int
	order = 1
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.CoverImg, &article.State, &article.CategoryID, &article.CreateUser, &article.CreateTime, &article.UpdateTime)
		if err != nil {
			return models.PageBean[models.Article]{}, err
		} else if order > (pageNum-1)*pageSize && order <= (pageNum-1)*pageSize+pageSize {
			article.CTime = article.CreateTime.Format("2006-01-02 15:04:05")
			article.UTime = article.UpdateTime.Format("2006-01-02 15:04:05")
			PageBean.Items = append(PageBean.Items, article)
		}
		order++
	}
	PageBean.Total = int64(order - 1)
	return PageBean, rows.Err()

}

func (a *ArticleMapper) Add(article models.Article, id int) error {
	insertSQL := "insert into article (title, content, cover_img, state, category_id, create_user, create_time, update_time) VALUES (?,?,?,?,?,?,?,?)"
	_, err := anno.Db.Exec(insertSQL, article.Title, article.Content, article.CoverImg, article.State, article.CategoryID, id, time.Now(), time.Now())
	return err
}

func (a *ArticleMapper) Update(article models.Article) error {
	updateSQL := "update article set title=?,content=?,category_id=?,cover_img=?,state=?,update_time=? where id = ?"
	_, err := anno.Db.Exec(updateSQL, article.Title, article.Content, article.CategoryID, article.CoverImg, article.State, time.Now(), article.ID)
	return err
}

func (a *ArticleMapper) Delete(id int) error {
	deleteSQL := "delete from article where id=?"
	_, err := anno.Db.Exec(deleteSQL, id)
	return err
}
