/* 站点地图

XML格式，支持百度和Google。
*/

package controllers

import (
    "bytes"
    "github.com/astaxie/beego"
    "text/template"
)

type SitemapController struct {
    beego.Controller
}

func (this *SitemapController) Get() {
    // 获得XML站点地图中的`url`元素列表
    urls := []SitemapUrl{}
    articles := []Article{}
    orm = InitDb()
    err = orm.OrderBy("-pubdate, -id").FindAll(&articles)
    for _, article := range(articles) {
        url := SitemapUrl{
            Loc: GetArticleFullUrl(article),
            LastMod: article.Pubdate,
        }
        urls = append(urls, url)
    }

    // 手工渲染XML站点地图模板文件
    t := template.New("sitemap.tpl")
    t, err = t.ParseFiles("views/sitemap.tpl")
    Check(err)
    var content bytes.Buffer
    err = t.Execute(&content, urls)
    Check(err)
    sitemap := content.String()
    this.Ctx.ContentType("application/xml")
    this.Ctx.WriteString(sitemap)
}
