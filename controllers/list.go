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
    SiteName = "SEOCMS"    // 网站名称
    ItemsPerPage = 10    // 列表页上每页显示文章数量
)

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName
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
    start := (pagenum -1) * ItemsPerPage

    if categoryNameEn == "" {    // 首页
        // 获取文章总数
        orm = InitDb()
        allArticles := []Article{}
        err = orm.FindAll(&allArticles)
        total := len(allArticles)    // 文章总数
        Debug("There are %d articles in total.", total)

        // 获取当前页的文章列表
        articles := []Article{}
        err = orm.OrderBy("-pubdate").Limit(ItemsPerPage, start).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 获取分页导航HTML代码
        paginator := GetPaginator(total, ItemsPerPage, pagenum)
        this.Data["Paginator"] = paginator

        // 设置页面标题
        this.Data["PageTitle"] = beego.AppConfig.String("appname")

        // 设置边栏
        this.Data["Sidebar"] = GetSidebar("home", 0)

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
        err = orm.Where("category=?", category.Id).OrderBy("-pubdate").Limit(ItemsPerPage, start).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 获取分页导航HTML代码
        paginator := GetPaginator(total, ItemsPerPage, pagenum)
        this.Data["Paginator"] = paginator

        // 设置页面标题
        this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", category.Name, beego.AppConfig.String("appname"))

        // 设置边栏
        this.Data["Sidebar"] = GetSidebar("category-list", category.Id)

        this.TplNames = "list/category_list.tpl"
    }
}
