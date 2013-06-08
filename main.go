package main

import (
	"seocms/controllers"
	"github.com/astaxie/beego"
)

func main() {
    // 管理后台：文章、分类管理
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/", &controllers.AdminController{})
    beego.Router("/:object(article)/:action(list|add)", &controllers.AdminController{})    // 显示文章列表、添加文章
    beego.Router("/:object(article)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑文章、修改文章
    beego.Router("/:object(category)/:action(list|add)", &controllers.AdminController{})    // 显示分类列表、添加分类
    beego.Router("/:object(category)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminController{})    // 编辑分类、修改分类

    // 文章列表页和文章内容页
    beego.Router("/", &controllers.ListController{})    // 首页(文章列表)
    beego.Router("/t/:tagid([0-9]+)/", &controllers.TagListController{})    // 标签列表页
    beego.Router("/tags", &controllers.TagsController{})    // 标签云；注：`/tags/`就不能匹配到，原因不明
    beego.Router("/:category(.+)/", &controllers.ListController{})    // 分类列表页
    beego.Router("/:category(.+)/:id([0-9]+)", &controllers.ArticleController{})    // 文章内容页

    beego.AddFuncMap("id2category", controllers.Id2category)    // 根据分类ID，获得分类名称
    beego.AddFuncMap("id2categoryEn", controllers.Id2categoryEn)    // 根据分类ID，获得分类英文名称
    beego.AddFuncMap("isSelected", controllers.IsSelected)    // 如果当前分类被选中，返回字符串`selected`
    beego.AddFuncMap("findTags", controllers.FindTags)    // 根据文章ID，获得文章标签列表
    beego.AddFuncMap("getSidebar", controllers.GetSidebar)    // 根据页面类型，获得边栏HTML代码；页面类型包括：首页、分类列表页、标签列表页和文章内容页

    beego.Run()
}

