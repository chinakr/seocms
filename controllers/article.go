package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
)

type ArticleController struct {
    beego.Controller
}

func (this *ArticleController) Get() {
    // 设置Template布局文件
    this.Layout = "layout.tpl"

    // 设置网站名称
    this.Data["SiteName"] = SiteName

    // 获得URL中包含的分类英文名称和文章ID
    categoryNameEn := this.Ctx.Params[":category"]
    articleId := this.Ctx.Params[":id"]
    Debug("It's article %v of category %s.", articleId, categoryNameEn)

    // 获得分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

    // 获得分类信息，用于导航栏高亮和面包屑
    orm = InitDb()
    category := Category{}
    err = orm.Where("name_en=?", categoryNameEn).Find(&category)
    Check(err)
    this.Data["Category"] = category

    // 获得文章信息
    orm = InitDb()
    article := Article{}
    err = orm.Where("id=?", articleId).Find(&article)
    Check(err)
    this.Data["Article"] = article

    // 设置页面标题、页面描述
    this.Data["PageTitle"] = fmt.Sprintf("%s_%s_%s", article.Title, category.Name, beego.AppConfig.String("appname"))
    this.Data["PageDescription"] = article.Abstract

    // 设置Template文件(默认自动渲染)
    this.TplNames = "article/article.tpl"
}
