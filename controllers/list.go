package controllers

import (
    "github.com/astaxie/beego"
)

type ListController struct {
    beego.Controller
}

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.TplNames = "list.tpl"
}
