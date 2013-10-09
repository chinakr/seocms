/* 后台管理

首页head管理
*/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminSiteController struct {
	beego.Controller
}

var ChannelName = "网站管理" // 频道名称

func (this *AdminSiteController) Get() {
	this.CheckLogin()                          // 检查用户是否登录
	this.Layout = "layout_admin.tpl"           // 页面模板布局文件
	object := this.Ctx.Input.Params(":object") // 当前配置信息
	switch object {
	case "head":
		this.Data["PageTitle"] = fmt.Sprint("首页head管理_%s_%s", ChannelName, SiteName) // 页面标题
		site := Site{}
		orm = InitDb()
		err = orm.Where("name=?", "head").Find(&site)
		//Check(err)
		if err != nil { // 如果head配置信息不存在，则创建这条配置信息
			site = Site{
				Name:    "head",
				Content: "",
			}
			orm = InitDb()
			err = orm.Save(&site)
			Check(err)
		}
		this.Data["Site"] = site              // 当前首页head
		this.TplNames = "admin/edit_head.tpl" // 页面模板文件
	case "body":
		this.Data["PageTitle"] = fmt.Sprint("通用body管理_%s_%s", ChannelName, SiteName) // 页面标题
		site := Site{}
		orm = InitDb()
		err = orm.Where("name=?", "body").Find(&site)
		//Check(err)
		if err != nil { // 如果body配置信息不存在，则创建这条配置信息
			site = Site{
				Name:    "body",
				Content: "",
			}
			orm = InitDb()
			err = orm.Save(&site)
			Check(err)
		}
		this.Data["Site"] = site              // 当前通用body
		this.TplNames = "admin/edit_body.tpl" // 页面模板文件
	case "sitemap":
		sitemaps := []string{}

		sitemapHome := GetSitemapHomeFullUrl() // 站点地图默认页：首页、列表页、最新的100篇文章

		sitemapPage := []string{} // 站点地图内容页：除最新的100篇文章外的所有文章，从老到新，每页最多1000篇文章
		pagenum := 1
		for {
			urls := GetSitemapPage(pagenum)
			if len(urls) > 0 {
				sitemapPage = append(sitemapPage, fmt.Sprintf("%s?page=%d", sitemapHome, pagenum))
			} else {
				break
			}
			pagenum += 1
		}

		sitemaps = append(sitemaps, sitemapHome)
		sitemaps = append(sitemaps, sitemapPage...)

		this.Data["Sitemaps"] = sitemaps
		this.TplNames = "admin/sitemap_list.tpl" // 页面模板文件
	}
}

func (this *AdminSiteController) Post() {
	this.CheckLogin()                          // 检查用户是否登录
	this.Layout = "layout_admin.tpl"           // 页面模板布局文件
	object := this.Ctx.Input.Params(":object") // 当前配置信息
	switch object {
	case "head":
		this.Data["PageTitle"] = fmt.Sprintf("首页head管理_%s_%s", ChannelName, SiteName) // 页面标题
		content := this.Input().Get("content")                                        // 获取表单数据

		// 保存设置
		site := Site{}
		orm = InitDb()
		err = orm.Where("name=?", "head").Find(&site)
		Check(err)
		site.Content = content
		err = orm.Save(&site)
		Check(err)

		this.Data["Site"] = site
		this.TplNames = "admin/edit_head.tpl" // 页面模板文件
	case "body":
		this.Data["PageTitle"] = fmt.Sprintf("通用body管理_%s_%s", ChannelName, SiteName) // 页面标题
		content := this.Input().Get("content")                                        // 获取表单数据

		// 保存设置
		site := Site{}
		orm = InitDb()
		err = orm.Where("name=?", "body").Find(&site)
		Check(err)
		site.Content = content
		err = orm.Save(&site)
		Check(err)

		this.Data["Site"] = site
		this.TplNames = "admin/edit_body.tpl" // 页面模板文件
	}
}

// 检测用户是否登录
func (this *AdminSiteController) CheckLogin() (flag bool) {
	account := this.GetSession("account")
	Debug("Current user is `%s`.", account)
	if account == nil { // 用户未登录
		this.Ctx.Redirect(302, "/user/login") // 跳转到用户登录页面
	} else {
		this.Data["Account"] = account
	}
	return true
}
