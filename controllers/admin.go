package controllers

import (
    "github.com/astaxie/beego"
)

type AdminController struct {
    beego.Controller
}

func (this *AdminController) Get() {
    this.Layout = "layout.tpl"
    action := this.Ctx.Params[":action"]
    switch action {
    case "add":
        this.TplNames = "admin/add_article.tpl"
    case "edit":
        this.Data["Id"] = this.Ctx.Params[":id"]
        this.TplNames = "admin/edit_article.tpl"
    case "delete":
        this.Data["Id"] = this.Ctx.Params[":id"]
        this.TplNames = "admin/del_article.tpl"
    }
}
