/* 用户管理 */

package controllers

import (
    "fmt"
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
    "time"
)

type UserController struct {
    beego.Controller
}

func (this *UserController) Get() {
    this.Layout = "layout_admin.tpl"    // 后台管理模板布局文件
    action := this.Ctx.Params[":action"]    // 用户的添加、修改或删除
    SiteName := beego.AppConfig.String("appname")    // 网站名称
    this.Data["SiteName"] = SiteName
    this.Data["Categories"] = GetCategories()    // 分类列表，用于导航栏
    ChannelName := "用户管理"    // 频道名称
    switch action {
    case "":    // 用户列表
        users := []User{}
        orm = InitDb()
        err = orm.OrderBy("name").FindAll(&users)
        Check(err)
        this.Data["Users"] = users    // 用户列表
        this.Data["PageTitle"] = fmt.Sprintf("用户列表_%s_%s", ChannelName, SiteName)    // 页面标题
        this.TplNames = "admin/user_list.tpl"    // 页面模板文件
    case "add":    // 添加用户
        this.Data["PageTitle"] = fmt.Sprintf("添加用户_%s_%s", ChannelName, SiteName)    // 页面标题
        this.TplNames = "admin/add_user.tpl"    // 页面模板文件
    case "edit":    // 修改用户
        id := this.Ctx.Params[":id"]    // 用户ID
        Debug("Current user ID is `%s`", id)

        orm = InitDb()
        user := User{}
        err = orm.Where("id=?", id).Find(&user)
        Check(err)
        this.Data["User"] = user    // 当前用户

        this.Data["PageTitle"] = fmt.Sprintf("修改用户_%s_%s", ChannelName, SiteName)    // 页面标题
        this.TplNames = "admin/edit_user.tpl"    // 页面模板文件
    case "delete":    // 删除用户
        id := this.Ctx.Params[":id"]    // 用户ID
        Debug("Current user ID is `%s`", id)

        // 删除当前用户
        orm = InitDb()
        user := User{}
        err = orm.Where("id=?", id).Find(&user)
        Check(err)
        orm.Delete(&user)

        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    case "login":    // 用户登录
        this.Layout = "layout_one.tpl"    // 用户登录模板布局文件
        this.TplNames = "admin/login.tpl"    // 页面模板文件
    case "logout":    // 用户退出
        this.Layout = "layout_one.tpl"    // 用户登录模板布局文件
        this.DelSession("account")    // 删除session中的用户登录信息
        this.TplNames = "admin/logout.tpl"    // 页面模板文件
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

        // 检测E-mail或密码是否为空
        if email == "" || name == "" {
            this.Data["Message"] = "E-mail或用户名为空"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }

        // 如果两次输入的密码不一致，需重新填写
        if password != rePassword {
            this.Data["Message"] = "两次输入的密码不一致"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }

        // 检查E-mail或用户名是否已存在
        orm = InitDb()
        user := User{}
        err = orm.Where("email=? or name=?", email, name).Find(&user)
        if err == nil {
            this.Data["Message"] = "E-mail或用户名已存在"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }

        // 保存用户
        orm = InitDb()
        user = User{}
        user.Email = email
        user.Name = name
        user.Password = Sha1(password)
        user.Created = time.Now()
        err = orm.Save(&user)
        Check(err)
        Debug("User `%s` added.", user)

        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    case "edit":    // 修改用户
        id := this.Ctx.Params[":id"]    // 用户ID

        email := this.Input().Get("email")    // 用户E-mail
        name := this.Input().Get("name")    // 用户名
        password := this.Input().Get("password")    // 密码
        rePassword := this.Input().Get("re-password")    // 重复输入的密码

        // 检测E-mail或密码是否为空
        if email == "" || name == "" {
            this.Data["Message"] = "E-mail或用户名为空"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }

        // 如果两次输入的密码不一致，需重新填写
        if password != rePassword {
            this.Data["Message"] = "两次输入的密码不一致"
            this.Data["Email"] = email
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.TplNames = "admin/add_user.tpl"
            return
        }

        // 获得当前用户
        orm = InitDb()
        user := User{}
        err = orm.Where("id=?", id).Find(&user)
        Check(err)

        // 更新用户信息
        user.Email = email
        user.Name = name
        if password != "" {
            user.Password = Sha1(password)
        }
        user.Updated = time.Now()

        // 保存用户信息
        err = orm.Save(&user)
        Check(err)

        this.Ctx.Redirect(302, "/user/")    // 返回用户列表页面
    case "login":    // 用户登录
        name := this.Input().Get("name")    // 用户名
        password := this.Input().Get("password")    // 用户密码

        errFlag := false    // 判断登录是否出错

        // 检测用户名或密码是否为空
        if name == "" || password == "" {
            this.Data["Message"] = "用户名或密码为空"
            errFlag = true
        }

        orm = InitDb()
        user := User{}
        err = orm.Where("name=? and password=?", name, Sha1(password)).Find(&user)
        if err != nil {
            this.Data["Message"] = "用户名或密码错误"
            errFlag = true
        } else {    // 用户名、密码验证成功
            // 保存用户登录信息
            this.SetSession("account", name)

            this.Ctx.Redirect(302, "/admin/")    // 跳转到管理后台首页
        }

        // 显示用户登录页面，再次登录
        if errFlag {
            this.Data["Name"] = name
            this.Data["Password"] = password
            this.Layout = "layout_one.tpl"
            SiteName := beego.AppConfig.String("appname")    // 网站名称
            this.Data["SiteName"] = SiteName
            this.Data["Categories"] = GetCategories()    // 分类列表，用于导航栏
            this.TplNames = "admin/login.tpl"
            return
        }


    }
}
