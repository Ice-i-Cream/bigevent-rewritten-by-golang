package main

import (
	_ "big_event/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
