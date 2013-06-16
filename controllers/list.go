/* 分类列表页(默认列表页) */

package controllers

import (
    "fmt"
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
)

type ListController struct {
    beego.Controller
}

const (
    //SiteName = "SEOCMS"    // 网站名称
    ItemsPerPage = 10    // 列表页上每页显示文章数量
)

var SiteName = beego.AppConfig.String("appname")    // 网站名称

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName    // 网站名称
    this.Data["Body"] = GetBody()    // 设置通用body，可以是统计代码
    categoryNameEn := this.Ctx.Params[":category"]
    //Debug("Current category is `%s`.", categoryNameEn)

    // 获取分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

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
    start := (pagenum - 1) * ItemsPerPage

    if categoryNameEn == "" {    // 首页
        this.Layout = "layout_index.tpl"
        // 获取文章总数
        orm = InitDb()
        allArticles := []Article{}
        err = orm.FindAll(&allArticles)
        total := len(allArticles)    // 文章总数
        Debug("There are %d articles in total.", total)

        // 获取当前页的文章列表
        articles := []Article{}
        err = orm.OrderBy("-pubdate, -id").Limit(ItemsPerPage, start).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 获取分页导航HTML代码
        paginator := GetPaginator(total, ItemsPerPage, pagenum)
        this.Data["Paginator"] = paginator

        description := beego.AppConfig.String("description")
        keywords := beego.AppConfig.String("keywords")

        // 设置页面标题
        if keywords != "" {
            this.Data["PageTitle"] = fmt.Sprintf("%s | %s", beego.AppConfig.String("appname"), keywords)
        } else {
            this.Data["PageTitle"] = beego.AppConfig.String("appname")
        }

        // 设置页面描述
        this.Data["PageDescription"] = description

        // 设置页面关键字
        this.Data["PageKeywords"] = keywords

        // 设置边栏
        this.Data["Sidebar"] = GetSidebar("home", 0)

        // 设置首页head
        site := Site{}
        orm = InitDb()
        err = orm.Where("name=?", "head").Find(&site)
        if err == nil {
            this.Data["Head"] = site.Content
        } else {
            this.Data["Head"] = ""
        }

        this.TplNames = "index.tpl"
    } else {    // 分类列表页
        // 获取当前分类
        orm = InitDb()
        category := Category{}
        err = orm.Where("name_en=?", categoryNameEn).Find(&category)
        Check(err)
        this.Data["Category"] = category

        // 获取当前分类文章总数
        allArticles := []Article{}
        err = orm.Where("category=?", category.Id).FindAll(&allArticles)
        total := len(allArticles)    // 文章总数
        Debug("There are %d articles in total in category `%s`.", total, category.Name)


        // 获取当前分类当前页的文章列表
        articles := []Article{}
        err = orm.Where("category=?", category.Id).OrderBy("-pubdate, -id").Limit(ItemsPerPage, start).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 获取分页导航HTML代码
        paginator := GetPaginator(total, ItemsPerPage, pagenum)
        this.Data["Paginator"] = paginator

        // 设置页面标题
        if category.Alias != "" {
            this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s | %s", category.Name, beego.AppConfig.String("appname"), category.Alias)
        } else {
            this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", category.Name, beego.AppConfig.String("appname"))
        }

        // 设置页面描述
        this.Data["PageDescription"] = category.Description

        // 设置页面关键字
        this.Data["PageKeywords"] = category.Alias

        // 设置边栏
        this.Data["Sidebar"] = GetSidebar("category", category.Id)

        this.TplNames = "list/category_list.tpl"
    }
}
