package controllers

import (
	"big_event/anno"
	"big_event/models"
	"big_event/utils"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"time"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Register() {
	exec := func(c *UserController) error {
		param, err := c.Input()
		if err != nil {
			return err
		}
		username := param.Get("username")
		password := param.Get("password")
		if anno.UsernamePattern.MatchString(username) && anno.PasswordPattern.MatchString(password) {
			return anno.UserService.Add(username, password)
		} else {
			return fmt.Errorf("请求参数错误")
		}
	}
	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}

func (c *UserController) Login() {
	exec := func(c *UserController) (string, error) {
		param, err := c.Input()
		if err != nil {
			return "", err
		}
		username := param.Get("username")
		password := param.Get("password")
		var user models.User
		user, err = anno.UserService.FindByName(username)
		if err != nil {
			return "", err
		}
		if user.Password != utils.Md5(password) {
			return "", fmt.Errorf("密码错误")
		}
		claims := jwt.MapClaims{
			"claims": jwt.MapClaims{
				"id":       user.ID,
				"username": user.Username,
			},
		}
		token, err := utils.GenToken(claims)
		if err != nil {
			return "", err
		}
		anno.RedisDb.Set(anno.Ctx, token, token, time.Hour)
		return token, nil
	}
	token, err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, token)
}

func (c *UserController) UserInfo() {
	exec := func(c *UserController) (models.User, error) {
		var username = anno.Thread.Get()["username"].(string)
		return anno.UserService.FindByName(username)
	}
	user, err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, user)
}

func (c *UserController) Update() {
	exec := func(c *UserController) error {
		var user = models.User{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		if err != nil {
			return err
		}
		validate := validator.New()
		err = validate.Struct(user)
		if err != nil {
			return err
		}
		return anno.UserService.Update(user)
	}
	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}

func (c *UserController) UpdateAvatar() {
	exec := func(c *UserController) error {
		param, err := c.Input()
		if err != nil {
			return err
		}
		url := param.Get("avatarUrl")
		id := anno.Thread.Get()["id"].(float64)
		return anno.UserService.UpdateAvatar(int(id), url)
	}

	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}

func (c *UserController) UpdatePwd() {
	exec := func(c *UserController) error {
		id := anno.Thread.Get()["id"].(float64)
		username := anno.Thread.Get()["username"].(string)
		user, err := anno.UserService.FindByName(username)
		if err != nil {
			return err
		}
		data := c.Ctx.Input.RequestBody
		var password models.Password
		err = json.Unmarshal(data, &password)
		if err != nil {
			return err
		}
		if !anno.PasswordPattern.MatchString(password.NewPwd) {
			return fmt.Errorf("新密码格式错误")
		} else if !(utils.Md5(password.OldPwd) == user.Password) {
			return fmt.Errorf("原密码错误")
		} else {
			password.NewPwd = utils.Md5(password.NewPwd)
			return anno.UserService.UpdatePwd(password, int(id))
		}
	}
	err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, nil)
}
