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
    action := this.Ctx.Params[":action"]    // 友情链接列表、添加友情链接、修改友情链接或删除友情链接
    switch action {
    case "":    // 友情链接列表
        links := []Link{}
        orm = InitDb()
        err = orm.OrderBy("name").FindAll(&links)
        Check(err)
        this.Data["Links"] = links    // 友情链接列表
        this.TplNames = "admin/link_list.tpl"    // 页面模板文件
    case "add":    // 添加友情链接
        this.TplNames = "admin/add_link.tpl"    // 页面模板文件
    case "edit":    // 修改友情链接
        id := this.Ctx.Params[":id"]    // 友情链接ID
        link := Link{}
        orm = InitDb()
        err = orm.Where("id=?", id).Find(&link)
        Check(err)
        this.Data["Link"] = link    // 当前友情链接
        this.TplNames = "admin/edit_link.tpl"    // 页面模板文件
    case "delete":    // 删除友情链接
        id :=  this.Ctx.Params[":id"]    // 友情链接ID

        // 获得当前友情链接
        orm = InitDb()
        link := Link{}
        err = orm.Where("id=?", id).Find(&link)
        Check(err)

        orm.Delete(&link)    // 删除友情链接

        // 跳转到友情链接列表页
        this.Ctx.Redirect(302, "/link/")
    }
}

func (this *AdminLinkController) Post() {
    this.CheckLogin()    // 检查用户是否登录
    this.Layout = "layout_admin.tpl"    // 页面模板布局文件
    action := this.Ctx.Params[":action"]    // 添加友情链接或修改友情链接
    switch action {
    case "add":    // 添加友情链接
        // 获取表单数据
        name := this.Input().Get("name")
        url := this.Input().Get("url")
        description := this.Input().Get("description")

        // 保存友情链接
        link := Link{}
        link.Name = name
        link.Url = url
        link.Description = description
        orm = InitDb()
        err = orm.Save(&link)
        Check(err)

        // 跳转到友情链接列表页
        this.Ctx.Redirect(302, "/link/")
    case "edit":    // 修改友情链接
        id := this.Ctx.Params[":id"]    // 友情链接ID

        // 获得当前友情链接
        link := Link{}
        orm = InitDb()
        err = orm.Where("id=?", id).Find(&link)
        Check(err)

        // 获取表单数据
        name := this.Input().Get("name")
        url := this.Input().Get("url")
        description := this.Input().Get("description")

        // 保存友情链接
        link.Name = name
        link.Url = url
        link.Description = description
        err = orm.Save(&link)
        Check(err)

        // 跳转到友情链接列表页
        this.Ctx.Redirect(302, "/link/")
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
