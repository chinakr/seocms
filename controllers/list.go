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
    ItemsPerPage = 5    // 列表页上每页显示文章数量
)

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName
    categoryNameEn := this.Ctx.Params[":category"]
    //Debug("Current category is `%s`.", categoryNameEn)

    // 获取分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

    // 获得当前页码
    pagenum, err := this.GetInt("page")
    //Check(err)
    if err != nil {
        //Debug("Can't fetch page num with error `%s`.", err)
        Debug("Page number not specified.")
        pagenum = 1
    } else {
        Debug("Current page number is `%v`.", pagenum)
    }

    // 计算起始文章序号
    start := (int(pagenum) -1) * ItemsPerPage

    if categoryNameEn == "" {    // 首页
        // 获得完整的文章列表
        orm = InitDb()
        articles := []Article{}
        err = orm.OrderBy("-pubdate").Limit(ItemsPerPage, start).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 设置页面标题
        this.Data["PageTitle"] = beego.AppConfig.String("appname")

        this.TplNames = "index.tpl"
    } else {    // 分类列表页
        // 获取当前分类
        orm = InitDb()
        category := Category{}
        err = orm.Where("name_en=?", categoryNameEn).Find(&category)
        Check(err)
        this.Data["Category"] = category

        // 获取当前分类文章列表
        orm = InitDb()
        articles := []Article{}
        err = orm.Where("category=?", category.Id).OrderBy("-pubdate").Limit(ItemsPerPage).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 设置页面标题
        this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", category.Name, beego.AppConfig.String("appname"))

        this.TplNames = "list/category_list.tpl"
    }
}
