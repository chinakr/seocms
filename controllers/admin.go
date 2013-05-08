package controllers

import (
    "github.com/astaxie/beego"
)

type AdminController struct {
    beego.Controller
}

func (this *AdminController) Get() {
    this.Layout = "layout_admin.tpl"
    action := this.Ctx.Params[":action"]
    switch action {
    case "add":
        this.Data["PageTitle"] = "添加文章_文章管理_SEOCMS"
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
}
