package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/c/1", &controllers.ArticleController{})    // 文章内容页
    beego.Router("/:object(article)/:action(list|add)", &controllers.AdminController{})    // 显示文章列表、添加文章
    beego.Router("/:object(article)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑文章、修改文章
    beego.Router("/:object(category)/:action(list|add)", &controllers.AdminController{})    // 显示分类列表、添加分类
    beego.Router("/:object(category)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑分类、修改分类
	beego.Run()
}

