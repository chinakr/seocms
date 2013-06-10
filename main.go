package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
    // 管理后台：文章、分类管理
    beego.Router("/admin", &controllers.AdminController{})    // 管理后台后页；如果写成`/admin/`无法匹配到
    beego.Router("/:object(article)/:action(list|add)", &controllers.AdminController{})    // 显示文章列表、添加文章
    beego.Router("/:object(article)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑文章、删除文章
    beego.Router("/:object(category)/:action(list|add)", &controllers.AdminController{})    // 显示分类列表、添加分类
    beego.Router("/:object(category)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑分类、删除分类

    // 用户管理
    beego.Router("/user", &controllers.UserController{})    // 用户列表；如果写成`/user/`无法匹配到
    beego.Router("/user/:action(add)", &controllers.UserController{})    // 添加用户
    beego.Router("/user/:action(edit|delete)/:id([0-9]+)", &controllers.UserController{})    // 编辑用户、删除用户

    // 文章列表页和文章内容页
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/t/:tagid([0-9]+)/", &controllers.TagListController{})    // 标签列表页
    beego.Router("/tags", &controllers.TagsController{})    // 标签云；注：`/tags/`就不能匹配到，原因不明
    beego.Router("/:category(.+)/", &controllers.ListController{})    // 分类列表页
    beego.Router("/:category(.+)/:id([0-9]+)", &controllers.ArticleController{})    // 文章内容页

    // 模板函数
    beego.AddFuncMap("breadcrumb", controllers.Breadcrumb)    // 根据管理后台的频道名称、页面名称，返回面包屑
    beego.AddFuncMap("id2category", controllers.Id2category)    // 根据分类ID，获得分类名称
    beego.AddFuncMap("id2categoryEn", controllers.Id2categoryEn)    // 根据分类ID，获得分类英文名称
    beego.AddFuncMap("isSelected", controllers.IsSelected)    // 如果当前分类被选中，返回字符串`selected`
    beego.AddFuncMap("findTags", controllers.FindTags)    // 根据文章ID，获得文章标签列表
    beego.AddFuncMap("getSidebar", controllers.GetSidebar)    // 根据页面类型，获得边栏HTML代码；页面类型包括：首页、分类列表页、标签列表页和文章内容页

    beego.Run()
}

