package controllers

import (
    "github.com/astaxie/beedb"
    _ "github.com/ziutek/mymysql/godrv"
    "database/sql"
    "time"
)

type Category struct {    // 文章分类
    Id int    // 分类ID
    Name string    // 分类名称
    NameEn string    // 分类英文名称(包括拼音或缩写)
    Description string    // 分类描述
    Alias string    // 分类别名(别名之间用`, `分隔)
    //Alias []string    // 分类别名(列表)
}

type Article struct {    // 文章
    Id int    // 文章ID
    Title string    // 文章标题
    Abstract string    // 文章摘要，Markdown格式
    AbstractHtml string    // 文章摘要，HTML格式
    Content string    // 文章内容，Markdown格式
    ContentHtml string    // 文章内容，HTML格式
    Pubdate time.Time    // 发布日期
    Updated time.Time    // 最后更新日期
    Category int    // 文章分类(外键)
}

type Tag struct {    // 文章标签
    Id int    // 标签ID
    Name string    // 标签名称
    NameEn string    // 标签英文名称(包括拼音和缩写)
    Description string    // 标签描述
    Alias string    // 标签别名(别名之前用`, `分隔)
}

type ArticleTags struct {    // 文章-标签对应关系
    Id int    // 文章-标签ID
    Article int    // 文章ID
    Tag int    // 标签ID
}

func InitDb() (orm beedb.Model) {
    database := "seocms"
    username := "seocms"
    password := "helloworld"
    db, err := sql.Open("mymysql", database + "/" + username + "/" + password)
    Check(err)
    orm = beedb.New(db)
    return
}

// 在模板中根据分类ID得到分类名称
func Id2category(id int) (category string) {
    orm := InitDb()
    categoryObj := Category{}
    err = orm.Where("id=?", id).Find(&categoryObj)
    Check(err)
    category = categoryObj.NameEn
    return
}

// 如果当前分类被选中则返回` slected`字符串
func IsSelected(categoryName string, categoryId int) (isSelected bool) {
    orm := InitDb()
    category := Category{}
    err = orm.Where("id=?", categoryId).Find(&category)
    Check(err)
    if categoryName == category.Name {
        isSelected = true
    } else {
        isSelected = false
    }
    return
}
