/* 后台管理

友情链接管理：列表、添加、修改、删除。
*/

package controllers

import (
    "github.com/astaxie/beego"
)

type AdminLinkController struct {
    beego.Controller
}

func (this *AdminLinkController) Get() {
    this.CheckLogin()    // 检查用户是否登录
    this.Layout = "layout_admin.tpl"    // 页面模板布局文件
    action := this.Ctx.Params[":action"]    // 标签列表、添加标签、修改标签或删除标签
    switch action {
    case "":    // 标签列表
        this.TplNames = "admin/link_list.tpl"    // 页面模板文件
    case "add":    // 添加标签
        this.TplNames = "admin/add_link.tpl"    // 页面模板文件
    case "edit":    // 修改标签
        this.TplNames = "admin/edit_link.tpl"    // 页面模板文件
    case "delete":    // 删除标签
        this.TplNames = "admin/delete_link.tpl"    // 页面模板文件
    }
}

func (this *AdminLinkController) Post() {
    this.CheckLogin()    // 检查用户是否登录
    this.Layout = "layout_admin.tpl"    // 页面模板布局文件
    action := this.Ctx.Params[":action"]    // 添加标签或修改标签
    switch action {
    case "add":    // 添加标签
        //
    case "edit":    // 修改标签
        //
    }
}

// 检测用户是否登录
func (this *AdminLinkController) CheckLogin() (flag bool) {
    account := this.GetSession("account")
    Debug("Current user is `%s`.", account)
    if account == nil {    // 用户未登录
        this.Ctx.Redirect(302, "/user/login")    // 跳转到用户登录页面
    } else {
        this.Data["Account"] = account
    }
    return true
}
