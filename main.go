package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
    // 管理后台：网站管理
    beego.Router("/site/:object(head)", &controllers.AdminSiteController{})    // 首页head管理
    beego.Router("/site/:object(body)", &controllers.AdminSiteController{})    // 通用body管理

    // 管理后台：文章、分类管理
    beego.Router("/admin", &controllers.AdminController{})    // 管理后台后页；如果写成`/admin/`无法匹配到
    beego.Router("/:object(article)/:action(list|add)", &controllers.AdminController{})    // 显示文章列表、添加文章
    beego.Router("/:object(article)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑文章、删除文章
    beego.Router("/:object(category)/:action(list|add)", &controllers.AdminController{})    // 显示分类列表、添加分类
    beego.Router("/:object(category)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑分类、删除分类

    // 管理后台：标签管理
    beego.Router("/tag", &controllers.AdminTagController{})    // 标签列表
    beego.Router("/tag/:action(edit)/:id([0-9]+)", &controllers.AdminTagController{})    // 编辑标签

    // 管理后台：友情链接管理
    beego.Router("/link", &controllers.AdminLinkController{})    // 友情链接列表
    beego.Router("/link/:action(add)", &controllers.AdminLinkController{})    // 添加友情链接
    beego.Router("/link/:action(edit|delete)/:id([0-9]+)", &controllers.AdminLinkController{})    // 编辑友情链接、删除友情链接

    // 用户管理
    beego.Router("/user", &controllers.UserController{})    // 用户列表；如果写成`/user/`无法匹配到
    beego.Router("/user/:action(add)", &controllers.UserController{})    // 添加用户
    beego.Router("/user/:action(edit|delete)/:id([0-9]+)", &controllers.UserController{})    // 编辑用户、删除用户
    beego.Router("/user/:action(login|logout)", &controllers.UserController{})

    // 文章列表页和文章内容页
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/t/:tagid([0-9]+)/", &controllers.TagListController{})    // 标签列表页
    beego.Router("/tags", &controllers.TagsController{})    // 标签云；注：`/tags/`就不能匹配到，原因不明
    beego.Router("/:category(.+)/", &controllers.ListController{})    // 分类列表页
    beego.Router("/:category(.+)/:id([0-9]+)", &controllers.ArticleController{})    // 文章内容页

    // 文章RSS feed
    beego.Router("/rss", &controllers.RssController{})

    // 模板函数
    beego.AddFuncMap("breadcrumb", controllers.Breadcrumb)    // 根据管理后台的频道名称、页面名称，返回面包屑
    beego.AddFuncMap("alert", controllers.Alert)    // 返回Bootstrap格式的提示信息HTML代码
    beego.AddFuncMap("id2category", controllers.Id2category)    // 根据分类ID，获得分类名称
    beego.AddFuncMap("id2categoryEn", controllers.Id2categoryEn)    // 根据分类ID，获得分类英文名称
    beego.AddFuncMap("isSelected", controllers.IsSelected)    // 如果当前分类被选中，返回字符串`selected`
    beego.AddFuncMap("findTags", controllers.FindTags)    // 根据文章ID，获得文章标签列表
    beego.AddFuncMap("getSidebar", controllers.GetSidebar)    // 根据页面类型，获得边栏HTML代码；页面类型包括：首页、分类列表页、标签列表页和文章内容页

    beego.SessionOn = true    // 启用session

    beego.Run()
}

