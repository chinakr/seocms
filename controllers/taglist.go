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
    this.Data["Body"] = GetBody()    // 设置通用body，可以是统计代码

    tagId := this.Ctx.Params[":tagid"]
    Debug("Current tag id is `%s`.", tagId)

    // 获取当前页码
    pagenumInt64, err := this.GetInt("page")
    //Check(err)
    pagenum := int(pagenumInt64)
    if err != nil {
        //Debug("Can't fetch page num with error `%s`.", err)
        Debug("Page number not specified.")
        pagenum = 1
    } else {
        Debug("Current page number is `%v`.", pagenum)
    }

    // 计算起始文章序号
    start := (pagenum -1) * ItemsPerPage

    // 获得当前标签
    orm = InitDb()
    tag := Tag{}
    err = orm.Where("id=?", tagId).Find(&tag)
    Check(err)
    this.Data["Tag"] = tag

    // 获取当前标签下的文章总数
    articleTagsList := []ArticleTags{}
    err = orm.Where("tag=?", tag.Id).FindAll(&articleTagsList)
    Check(err)
    allArticles := []Article{}
    for _, articleTags := range(articleTagsList) {
        Debug("Tag ID `%d`, article ID `%d`.", articleTags.Tag, articleTags.Article)
        articleItem := Article{}
        err = orm.Where("id=?", articleTags.Article).Find(&articleItem)
        Debug("Article ID `%d`, article `%s`.", articleItem.Id, articleItem.Title)
        allArticles = append(allArticles, articleItem)
    }
    total := len(allArticles)    // 文章总数
    Debug("There are %d articles in total with tag `%s`.", total, tag.Name)

    // 获得当前标签当前页下的文章列表
    end := start + 1
    if start + ItemsPerPage > total {
        end = total
    } else {
        end = start + ItemsPerPage
    }
    articles := allArticles[start:end]
    this.Data["Articles"] = SortArticle(articles)    // 文章按日期和ID排序

    // 获取分页导航HTML代码
    paginator := GetPaginator(total, ItemsPerPage, pagenum)
    this.Data["Paginator"] = paginator

    // 获取分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

    keywords := tag.Alias

    // 设置页面标题
    if keywords != "" {
        this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s | %s", tag.Name, beego.AppConfig.String("appname"), keywords)
    } else {
        this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", tag.Name, beego.AppConfig.String("appname"))
    }

    // 设置页面描述
    this.Data["PageDescription"] = tag.Description

    // 设置页面关键字
    this.Data["PageKeywords"] = keywords

    // 设置边栏
    this.Data["Sidebar"] = GetSidebar("tag", tag.Id)

    this.TplNames = "list/tag_list.tpl"
}
