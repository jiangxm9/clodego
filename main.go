package main

import (
	"github.com/astaxie/beego"
	"github.com/github-user/cloudgo/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})// 设置路由，传入controller处理
	beego.Run(":9527") //开放9527端口运行
}
