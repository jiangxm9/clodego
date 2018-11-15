package main

import (
	"github.com/astaxie/beego"
	"github.com/github-user/cloudgo/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run(":9527")
}
