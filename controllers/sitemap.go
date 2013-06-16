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
    //urls := GetSitemapPage(1)

    // 获取当前页码
    pagenumInt64, err := this.GetInt("page")
    pagenum := int(pagenumInt64)
    if err != nil {    // 未指定页码，即站点地图默认页
        urls = GetSitemapHome()
    } else {
        urls = GetSitemapPage(pagenum)
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

// 站点地图默认页，包括首页、列表页和最近的100篇文章
func GetSitemapHome() (urls []SitemapUrl) {
    // 首页
    sitemapUrlIndex := SitemapUrl{
        Loc: beego.AppConfig.String("appurl"),
    }

    // 分类列表页
    sitemapUrlCategoryList := []SitemapUrl{}
    categories := []Category{}
    orm = InitDb()
    err = orm.OrderBy("name").FindAll(&categories)
    Check(err)
    for _, category := range(categories) {
        sitemapUrlCategoryList = append(sitemapUrlCategoryList, SitemapUrl{
            Loc: GetCategoryListFullUrl(category),
        })
    }

    // 标签列表页
    sitemapUrlTagList := []SitemapUrl{}
    tags := []Tag{}
    orm = InitDb()
    err = orm.OrderBy("name").FindAll(&tags)
    Check(err)
    for _, tag := range(tags) {
        sitemapUrlTagList = append(sitemapUrlTagList, SitemapUrl{
            Loc: GetTagListFullUrl(tag),
        })
    }

    // 标签云
    sitemapUrlTagCloud := SitemapUrl{
        Loc: GetTagCloudFullUrl(),
    }

    // 最新的100篇文章
    sitemapUrlArticle := []SitemapUrl{}
    articles := []Article{}
    orm = InitDb()
    err = orm.OrderBy("-pubdate, -id").Limit(100).FindAll(&articles)
    Check(err)
    for _, article := range(articles) {
        sitemapUrlArticle = append(sitemapUrlArticle, SitemapUrl{
            Loc: GetArticleFullUrl(article),
        })
    }

    urls = append(urls, sitemapUrlIndex)
    urls = append(urls, sitemapUrlCategoryList...)
    urls = append(urls, sitemapUrlTagList...)
    urls = append(urls, sitemapUrlTagCloud)
    urls = append(urls, sitemapUrlArticle...)
    return
}

// 站点地图内容页，从最早的文章开始，每页包含1000个文章内容页URL
func GetSitemapPage(pagenum int) (urls []SitemapUrl) {
    //itemsPerPage := 5    // DEBUG: 每页包含的文章内容页URL个数
    itemsPerPage := 1000    // 每页包含的文章内容页URL个数
    //reserved := 3    // DEBUG: 站点地图默认页包含的文章数量
    reserved := 100    // 站点地图默认页包含的文章数量

    total := GetArticleCount()    // 文章总数

    // 计算起始文章序号
    start := (pagenum - 1) * itemsPerPage

    if total < reserved {    // 文章总数不足100篇时
        return
    }

    if start > total - reserved {    // 起始文章序号大于剩余文章总数时
        return
    }

    if start + itemsPerPage > total - reserved {    // 文章数不足1000时
        itemsPerPage = total - reserved - start
    }

    articles := []Article{}
    orm = InitDb()
    err = orm.OrderBy("pubdate, id").Limit(itemsPerPage, start).FindAll(&articles)
    for _, article := range(articles) {
        url := SitemapUrl{
            Loc: GetArticleFullUrl(article),
            //LastMod: article.Pubdate,
        }
        urls = append(urls, url)
    }
    return
}
