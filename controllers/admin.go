package controllers

import (
    "github.com/astaxie/beego"
)

type AdminController struct {
    beego.Controller
}

func (this *AdminController) Get() {
    this.Layout = "layout_admin.tpl"
    object := this.Ctx.Params[":object"]
    action := this.Ctx.Params[":action"]
    if object == "article" {
        switch action {
        case "list":
            this.Data["PageTitle"] = "文章列表_文章管理_SEOCMS"
            this.Data["Articles"] = []string{
                "这是第1篇博客",
                "这是第2篇博客",
                "这是第3篇博客",
            }
            this.TplNames = "admin/article_list.tpl"
        case "add":
            this.Data["PageTitle"] = "添加文章_文章管理_SEOCMS"
            this.Data["Categories"] = []string{"博客", "笔记"}
            this.TplNames = "admin/add_article.tpl"
        case "edit":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "编辑文章_文章管理_SEOCMS"
            this.TplNames = "admin/edit_article.tpl"
        case "delete":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "删除文章_文章管理_SEOCMS"
            this.TplNames = "admin/del_article.tpl"
        }
    } else if object == "category" {
        switch action {
        case "list":
            this.Data["PageTitle"] = "分类列表_文章管理_SEOCMS"
            this.Data["Categories"] = []string{
                "博客",
                "笔记",
            }
            this.TplNames = "admin/category_list.tpl"
        case "add":
            this.Data["PageTitle"] = "添加分类_文章管理_SEOCMS"
            this.TplNames = "admin/add_category.tpl"
        case "edit":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "修改分类_文章管理_SEOCMS"
            this.TplNames = "admin/edit_category.tpl"
        case "delete":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "删除分类_文章管理_SEOCMS"
            this.TplNames = "admin/delete_category.tpl"
        }
    }
}

func (this *AdminController) Post() {
    this.Ctx.Redirect(302, "/article/list")
}
