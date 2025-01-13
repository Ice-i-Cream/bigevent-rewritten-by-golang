package anno

import (
	"big_event/models"
	"big_event/service"
	"context"
	"database/sql"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis/v8"
	"github.com/timandy/routine"
	"regexp"
)

var UsernamePattern = regexp.MustCompile(`^\S{5,16}$`)
var PasswordPattern = regexp.MustCompile(`^\S{5,16}$`)
var Thread = routine.NewThreadLocal[map[string]interface{}]()

var UserService service.UserService
var CategoryService service.CategoryService
var ArticleService service.ArticleService

var Ctx context.Context
var RedisDb *redis.Client
var Db *sql.DB

var PostProcess = func(a *struct{ web.Controller }, err error, data interface{}) {
	if err != nil {
		a.Data["json"] = models.Error(err.Error())
	} else {
		if data != nil {
			a.Data["json"] = models.Success(data)
		} else {
			a.Data["json"] = models.SuccessNoData()
		}
	}
	err = a.ServeJSON()
	if err != nil {
		a.Ctx.Output.SetStatus(500)
	}
}
