package mapper

import (
	"big_event/anno"
	"big_event/models"
	"big_event/utils"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type UserMapper struct{}

func (u *UserMapper) Add(username string, password string) error {
	hashedPassword := utils.Md5(password) // 加密密码
	currentTime := time.Now()             // 获取当前时间
	insertSQL := "INSERT INTO user (username, password, create_time, update_time) VALUES (?, ?, ?, ?)"
	_, err := anno.Db.Exec(insertSQL, username, hashedPassword, currentTime, currentTime)
	if err.Error()[6:10] == "1062" {
		return fmt.Errorf("用户名已被占用")
	}
	return err
}

func (u *UserMapper) FindByName(username string) (models.User, error) {
	selectSQL := "SELECT * FROM user WHERE username = ?"
	var user models.User
	err := anno.Db.QueryRow(selectSQL, username).Scan(&user.ID, &user.Username, &user.Password, &user.Nickname, &user.Email, &user.UserPic, &user.CreateTime, &user.UpdateTime)
	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, fmt.Errorf("用户名错误")
	}
	return user, err
}

func (u *UserMapper) Update(user models.User) error {
	updateSQL := "UPDATE user SET nickname = ?,email= ?, update_time = ? WHERE id = ?"
	_, err := anno.Db.Exec(updateSQL, user.Nickname, user.Email, time.Now(), user.ID)
	return err
}

func (u *UserMapper) UpdateAvatar(id int, url string) error {
	updateSQL := "UPDATE user SET user_pic = ? WHERE id = ?"
	_, err := anno.Db.Exec(updateSQL, url, id)
	return err
}

func (u *UserMapper) UpdatePwd(data models.Password, id int) error {
	updateSQL := "UPDATE user SET password = ? WHERE id = ?"
	_, err := anno.Db.Exec(updateSQL, data.NewPwd, id)
	return err
}
