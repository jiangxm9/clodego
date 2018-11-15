package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller //继承beego中的Controller类
}

//重写Get方法，这里参照beego中的todo实例，设置浏览器访问的是之前其他课实现的静态页面
func (this *MainController) Get() { 
	this.TplName = "tmnt.html" //this.Tpl就是需要渲染的模板，这里会默认读取views文件夹中的文件
	this.Render() //调用beego已实现的Render函数渲染
}
