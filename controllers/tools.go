package controllers

import (
    "code.google.com/p/mahonia"
    "github.com/knieriem/markdown"
    "bufio"
    "bytes"
    "fmt"
    "os"
    "runtime"
    "strings"
    "time"
)

var (
    DEBUG = true    // 输出调试信息
    // DEBUG = false    // 不输出调试信息
)

func Check(err error) {
    if err != nil {
        panic(err)
        os.Exit(1)
    }
}

// 当DEBUG开关打开时，输出调试信息
// 使用方法：`debug('格式化字符串', 参数1, 参数2, ...)`
// 输出结果：`DEBUG: 提示信息`(会自动换行)
// 相当于加强版的`fmt.Println()`，支持格式化字符串，输出以`DEBUG: `开头
func Debug(infos ...interface{}) {
    if DEBUG {
        fmt.Printf("DEBUG: " + fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
    }
}

// 占位符，什么也不干，用于临时避免`declared but not used`错误
func NDebug(infos ...interface{}) {
}

// 加强版的`fmt.Printf()`，能识别操作系统，避免终端输出时出现乱码，自动换行
func Info(infos ...interface{}) {
    if runtime.GOOS == "windows" {
        encoder := mahonia.NewEncoder("gbk")
        for i, item := range(infos) {
            item, ok := item.(string)
            if ok {
                infos[i] = encoder.ConvertString(item)
            }
        }
    }
    fmt.Printf(fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
}

// 把用`, `间隔的字符串转换为字符串列表(例如关键字列表)
func Str2slice(str string) []string {
    return strings.Split(str, ", ")
}

// 把字符串转换为time.Time对象
func Str2date(timeStr string) (timeObj time.Time, err error) {
    layout := "2006-01-02"
    timeObj, err = time.Parse(layout, timeStr)
    //Check(err)
    return
}

// 把Markdown文本转换为HTML
func Markdown2html(strMarkdown string) (html string) {
    p := markdown.NewParser(&markdown.Extensions{Smart: true})
    var buf bytes.Buffer
    w := bufio.NewWriter(&buf)
    r := bytes.NewBufferString(strMarkdown)
    p.Markdown(r, markdown.ToHTML(w))
    w.Flush()
    html = string(buf.Bytes())
    return
}

// 判断字符串切片中是否包含某个字符串
func SliceContains(strSlice []string, str string) bool {
    strMap := make(map[string]bool)
    for _, strItem := range(strSlice) {
        strMap[strItem] = true
    }
    if _, ok := strMap[str]; ok {
        return true
    }
    return false
}

// 从数据库获得分类列表
func GetCategories() (categories []Category) {
    orm = InitDb()
    categories = []Category{}
    err = orm.OrderBy("name").FindAll(&categories)
    Check(err)
    return
}
