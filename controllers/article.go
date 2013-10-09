package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	// 设置Template布局文件
	this.Layout = "layout_article.tpl"

	this.Data["Body"] = GetBody() // 设置通用body，可以是统计代码

	// 设置网站名称
	this.Data["SiteName"] = SiteName

	// 获得URL中包含的分类英文名称和文章ID
	categoryNameEn := this.Ctx.Input.Params(":category")
	articleId := this.Ctx.Input.Params(":id")
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

	// 友言评论
	b, err := ioutil.ReadFile("views/uy.html")
	Check(err)
	uyCode := string(b)
	this.Data["UY"] = uyCode

	// 设置页面描述
	this.Data["PageDescription"] = article.Abstract

	// 设置页面关键字
	tags := FindTagsText(article.Id)
	if tags == "" {
		this.Data["PageKeywords"] = category.Name
	} else {
		this.Data["PageKeywords"] = category.Name + ", " + tags
	}

	// 设置页面标题
	this.Data["PageTitle"] = fmt.Sprintf("%s_%s_%s | %s", article.Title, category.Name, beego.AppConfig.String("appname"), this.Data["PageKeywords"])

	// 设置边栏
	this.Data["Sidebar"] = GetSidebar("article", article.Id)

	// 设置Template文件(默认自动渲染)
	this.TplNames = "article/article.tpl"
}
