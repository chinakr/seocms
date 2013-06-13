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
    this.CheckLogin()    // 检查用户是否登录
    this.Layout = "layout_admin.tpl"    // 页面模板布局文件
    action := this.Ctx.Params[":action"]    // 标签列表或修改标签
    switch action {
    case "":    // 标签列表
        this.Data["PageTitle"] = fmt.Sprintf("标签列表_标签管理_%s", SiteName)    // 页面标题
        tags := []Tag{}
        orm = InitDb()
        err = orm.OrderBy("name").FindAll(&tags)
        Check(err)
        this.Data["Tags"] = tags    // 标签列表
        this.TplNames = "admin/tag_list.tpl"    // 页面模板文件
    case "edit":    // 修改标签
        this.Data["PageTitle"] = fmt.Sprintf("修改标签_标签管理_%s", SiteName)    // 页面标题
        this.TplNames = "admin/edit_tag.tpl"    // 页面模板文件
    }
}

func (this *AdminTagController) Post() {
    this.CheckLogin()    // 检查用户是否登录
    this.Layout = "layout_admin.tpl"    // 页面模板布局文件
}

// 检测用户是否登录
func (this *AdminTagController) CheckLogin() (flag bool) {
    account := this.GetSession("account")
    Debug("Current user is `%s`.", account)
    if account == nil {    // 用户未登录
        this.Ctx.Redirect(302, "/user/login")    // 跳转到用户登录页面
    } else {
        this.Data["Account"] = account
    }
    return true
}
