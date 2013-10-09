/* 后台管理

标签管理：列表、修改
*/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminTagController struct {
	beego.Controller
}

func (this *AdminTagController) Get() {
	this.CheckLogin()                          // 检查用户是否登录
	this.Layout = "layout_admin.tpl"           // 页面模板布局文件
	action := this.Ctx.Input.Params(":action") // 标签列表或修改标签
	switch action {
	case "": // 标签列表
		this.Data["PageTitle"] = fmt.Sprintf("标签列表_标签管理_%s", SiteName) // 页面标题

		allTags := []Tag{}
		orm = InitDb()
		err = orm.FindAll(&allTags)
		Check(err)
		total := len(allTags) // 标签总数

		// 获得当前页码
		pagenumInt64, err := this.GetInt("page")
		pagenum := int(pagenumInt64)
		if err != nil {
			pagenum = 1
		}

		start := (pagenum - 1) * ItemsPerPage // 起始标签序号

		// 获得当前页标签列表
		tags := []Tag{}
		err = orm.OrderBy("name").Limit(ItemsPerPage, start).FindAll(&tags)
		Check(err)
		this.Data["Tags"] = tags // 标签列表

		paginator := GetPaginator(total, ItemsPerPage, pagenum)
		this.Data["Paginator"] = paginator // 分页导航HTML代码

		this.TplNames = "admin/tag_list.tpl" // 页面模板文件
	case "edit": // 修改标签
		this.Data["PageTitle"] = fmt.Sprintf("修改标签_标签管理_%s", SiteName) // 页面标题
		id := this.Ctx.Input.Params(":id")                             // 标签ID
		tag := Tag{}
		err = orm.Where("id=?", id).Find(&tag)
		Check(err)
		this.Data["Tag"] = tag               // 当前标签
		this.TplNames = "admin/edit_tag.tpl" // 页面模板文件
	}
}

func (this *AdminTagController) Post() {
	this.CheckLogin()                  // 检查用户是否登录
	this.Layout = "layout_admin.tpl"   // 页面模板布局文件
	id := this.Ctx.Input.Params(":id") // 标签ID

	// 获得当前标签
	orm = InitDb()
	tag := Tag{}
	err = orm.Where("id=?", id).Find(&tag)
	Check(err)

	// 获得表单数据
	tag.Name = this.Input().Get("name")
	tag.NameEn = this.Input().Get("name_en")
	tag.Description = this.Input().Get("description")
	tag.Alias = this.Input().Get("alias")

	// 检查标签名称是否为空
	if tag.Name == "" {
		this.Data["Message"] = "标签名称不能为空。"
		this.Data["Tag"] = tag
		this.TplNames = "admin/edit_tag.tpl"
		return
	}

	err = orm.Save(&tag)
	Check(err)

	this.Ctx.Redirect(302, "/tag/")
}

// 检测用户是否登录
func (this *AdminTagController) CheckLogin() (flag bool) {
	account := this.GetSession("account")
	Debug("Current user is `%s`.", account)
	if account == nil { // 用户未登录
		this.Ctx.Redirect(302, "/user/login") // 跳转到用户登录页面
	} else {
		this.Data["Account"] = account
	}
	return true
}
