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

    // 获得当前页码
    pagenum, err := this.GetInt("page")
    //Check(err)
    if err != nil {
        //Debug("Can't fetch page num with error `%s`.", err)
        Debug("Page number not specified.")
    } else {
        Debug("Current page number is `%v`.", pagenum)
    }

    // 获得当前标签
    orm = InitDb()
    tag := Tag{}
    err = orm.Where("id=?", tagId).Find(&tag)
    Check(err)
    this.Data["Tag"] = tag

    // 获得当前标签下的文章列表
    articleTagsList := []ArticleTags{}
    err = orm.Where("tag=?", tag.Id).FindAll(&articleTagsList)
    Check(err)
    articles := []Article{}
    for _, articleTags := range(articleTagsList) {
        Debug("Tag ID `%d`, article ID `%d`.", articleTags.Tag, articleTags.Article)
        articleItem := Article{}
        err = orm.Where("id=?", articleTags.Article).Find(&articleItem)
        Debug("Article ID `%d`, article `%s`.", articleItem.Id, articleItem.Title)
        articles = append(articles, articleItem)
    }
    this.Data["Articles"] = articles

    // 获取分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

    // 设置页面标题
    this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", tag.Name, beego.AppConfig.String("appname"))

    this.TplNames = "list/tag_list.tpl"
}
