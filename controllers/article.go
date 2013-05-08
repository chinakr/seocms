package controllers

import (
    "github.com/astaxie/beego"
)

type ArticleController struct {
    beego.Controller
}

func (this *ArticleController) Get() {
    this.Layout = "layout.tpl"
    this.TplNames = "article.tpl"
}
