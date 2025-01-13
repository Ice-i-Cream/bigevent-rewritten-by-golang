package routers

import (
	"big_event/anno"
	"big_event/controllers"
	"big_event/interceptors"
	"big_event/mapper"
	"context"
	"database/sql"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, interceptors.LoginInterceptor)
	beego.SetStaticPath("/image", "file/")
	beego.Router("/", &controllers.MainController{})
	beego.Router("user/register", &controllers.UserController{}, "post:Register")
	beego.Router("user/login", &controllers.UserController{}, "post:Login")
	beego.Router("user/userInfo", &controllers.UserController{}, "get:UserInfo")
	beego.Router("user/update", &controllers.UserController{}, "put:Update")
	beego.Router("user/updateAvatar", &controllers.UserController{}, "patch:UpdateAvatar")
	beego.Router("user/updatePwd", &controllers.UserController{}, "patch:UpdatePwd")
	beego.Router("upload", &controllers.FileUploadController{}, "post:Upload")
	beego.Router("category", &controllers.CategoryController{}, "get:List")
	beego.Router("category", &controllers.CategoryController{}, "post:Add")
	beego.Router("category", &controllers.CategoryController{}, "put:Update")
	beego.Router("category", &controllers.CategoryController{}, "delete:Delete")
	beego.Router("category/detail", &controllers.CategoryController{}, "get:Detail")
	beego.Router("article", &controllers.ArticleController{}, "get:List")
	beego.Router("article", &controllers.ArticleController{}, "post:Add")
	beego.Router("article", &controllers.ArticleController{}, "put:Update")
	beego.Router("article", &controllers.ArticleController{}, "delete:Delete")

	anno.Db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/big_event?parseTime=True&loc=Local")
	anno.RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	anno.Ctx = context.Background()
	anno.UserService = &mapper.UserMapper{}
	anno.CategoryService = &mapper.CategoryMapper{}
	anno.ArticleService = &mapper.ArticleMapper{}

}
