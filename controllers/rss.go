/* 文章RSS feed */

package controllers

import (
    "bytes"
    "fmt"
    "github.com/astaxie/beego"
    "text/template"
)

const (
    RSS_ITEMS_COUNT = 10    // RSS feed中包含多少篇文章
)

type RssController struct {
    beego.Controller
}

func (this *RssController) Get() {
    channel := RssChannel{}    // RSS的`channel`元素
    channel.Title = beego.AppConfig.String("appname")    // 网站名称
    channel.Link = beego.AppConfig.String("appurl")    // 网站URL
    channel.Description = beego.AppConfig.String("description")    // 网站简介

    articles := []Article{}
    orm = InitDb()
    err = orm.OrderBy("-pubdate").OrderBy("-id").Limit(RSS_ITEMS_COUNT).FindAll(&articles)
    Check(err)

    for _, article := range(articles) {
        item := RssItem{}    // RSS channel的`item`元素
        item.Title = article.Title
        item.Link = GetArticleFullUrl(article)
        item.Category = Id2category(article.Category)
        //item.Description = "DEBUG"
        //item.Description = article.Content
        item.Description = article.ContentHtml
        //item.Description = beego.Htmlquote(article.ContentHtml)
        item.Pubdate = article.Pubdate
        item.Guid = fmt.Sprintf("%d", article.Id)

        channel.Items = append(channel.Items, item)
    }

    //this.Data["Channel"] = channel    // RSS的`channel`元素

    // 手工渲染RSS模板文件
    t := template.New("rss.tpl")
    t, err = t.ParseFiles("views/rss.tpl")
    Check(err)
    var content bytes.Buffer
    err = t.Execute(&content, channel)
    Check(err)
    rss := content.String()

    this.Ctx.ContentType("application/xml")
    this.Ctx.WriteString(rss)
}
