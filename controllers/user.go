/* 用户管理 */

package controllers

import (
    "fmt"
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
)

type UserController struct {
    beego.Controller
}

func (this *UserController) Get() {
    this.Layout = "layout_admin.tpl"    // 后台管理模板布局文件
    action := this.Ctx.Params[":action"]    // 用户的添加、修改或删除
    SiteName := beego.AppConfig.String("appname")    // 网站名称
    ChannelName := "用户管理"    // 频道名称
    switch action {
    case "":    // 用户列表
        this.Data["PageTitle"] = fmt.Sprintf("用户列表_%s_%s", ChannelName, SiteName)
        this.TplNames = "admin/user_list.tpl"    // 页面模板文件
    case "add":    // 添加用户
        this.Data["PageTitle"] = fmt.Sprintf("添加用户_%s_%s", ChannelName, SiteName)    // 页面标题
        this.TplNames = "admin/add_user.tpl"    // 页面模板文件
    case "edit":    // 修改用户
        this.Data["PageTitle"] = fmt.Sprintf("修改用户_%s_%s", ChannelName, SiteName)    // 页面标题
        id := this.Ctx.Params[":id"]    // 用户ID
        Debug("Current user ID is `%s`", id)
        this.TplNames = "admin/edit_user.tpl"    // 页面模板文件
    case "delete":    // 删除用户
        id := this.Ctx.Params[":id"]    // 用户ID
        Debug("Current user ID is `%s`", id)
        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    }
}

func (this *UserController) Post() {
    this.Layout = "layout_admin.tpl"    // 模板布局文件
    action := this.Ctx.Params[":action"]    // 用户的添加或修改
    switch action {
    case "add":    // 添加用户
        email := this.Input().Get("email")    // 用户E-mail
        name := this.Input().Get("name")    // 用户名
        password := this.Input().Get("password")    // 密码
        rePassword := this.Input().Get("re-password")    // 重复输入的密码
        if password != rePassword {
            this.Data["Message"] = "两次输入的密码不一致"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }
        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    case "edit":    // 修改用户
        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    }
}
