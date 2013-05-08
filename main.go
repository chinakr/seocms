package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/c/1", &controllers.ArticleController{})    // 文章内容页
    beego.Router("/article/:action(add)", &controllers.AdminController{})    // 添加文章
    beego.Router("/article/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑或修改文章
	beego.Run()
}

