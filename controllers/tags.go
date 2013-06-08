/* 标签云 */

package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
)

type TagsController struct {
    beego.Controller
}

func (this *TagsController) Get() {
    this.Layout = "layout.tpl"    // 模板布局文件
    this.Data["SiteName"] = SiteName    // 网站名称
    this.Data["Categories"] = GetCategories()    // 导航栏分类列表
    this.Data["PageTitle"] = fmt.Sprintf("所有话题_%s", beego.AppConfig.String("appname"))

    // 全部标签
    orm := InitDb()
    tags := []Tag{}
    err = orm.FindAll(&tags)
    Check(err)
    this.Data["Tags"] = tags

    this.Data["Sidebar"] = GetSidebar("home", 0)

    this.TplNames = "tags.tpl"    // 模板文件
}
