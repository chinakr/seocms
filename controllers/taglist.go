/* 标签列表页 */
package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
)

type TagListController struct {
    beego.Controller
}

func (this *TagListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName

    tagId := this.Ctx.Params[":tagid"]
    Debug("Current tag id is `%s`.", tagId)

    // 获得当前标签
    orm = InitDb()
    tag := Tag{}
    err = orm.Where("id=?", tagId).Find(&tag)
    Check(err)
    this.Data["Tag"] = tag

    // 获得当前标签下的文章列表
    //this.Data["Articles"] = articles

    // 设置页面标题
    this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", tag.Name, beego.AppConfig.String("appname"))

    this.TplNames = "list/tag_list.tpl"
}
