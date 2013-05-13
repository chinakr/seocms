package controllers

import (
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
)

type ListController struct {
    beego.Controller
}

const (
    SiteName = "SEOCMS"    // 网站名称
)

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName

    // 获取分类列表，用于导航栏
    orm = InitDb()
    categories := []Category{}
    err = orm.OrderBy("name").FindAll(&categories)
    Check(err)
    this.Data["Categories"] = categories

    this.TplNames = "list.tpl"
}
