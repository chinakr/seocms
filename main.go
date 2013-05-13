package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
    // 管理后台：文章、分类管理
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/:object(article)/:action(list|add)", &controllers.AdminController{})    // 显示文章列表、添加文章
    beego.Router("/:object(article)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑文章、修改文章
    beego.Router("/:object(category)/:action(list|add)", &controllers.AdminController{})    // 显示分类列表、添加分类
    beego.Router("/:object(category)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑分类、修改分类

    // 文章列表页和文章内容页
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/:category(.+)/", &controllers.ListController{})    // 分类列表页(以及SEO过的标签列表页)
    beego.Router("/:category(.+)/:id([0-9]+)", &controllers.ArticleController{})    // 文章内容页

    beego.AddFuncMap("id2category", controllers.Id2category)
    beego.AddFuncMap("isSelected", controllers.IsSelected)

    beego.Run()
}

