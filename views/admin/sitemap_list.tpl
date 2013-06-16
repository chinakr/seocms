<input type="hidden" name="location" value="sitemap-list" />
{{$breadcrumb := breadcrumb "网站管理" "站点地图列表"}}
{{str2html $breadcrumb}}
<div class="sitemap well">
    <h2>站点地图列表</h2>
    <ul>
        {{range .Sitemaps}}
        <li><a href="{{.}}" target="_blank">{{.}}</a></li>
        {{end}}
    </ul>
    <h2 class="submit">提交站点地图</h2>
    <p>向Google提交站点地图 <a href="http://support.google.com/webmasters/bin/answer.py?hl=zh-Hans&answer=183669&topic=8476&ctx=topic" target="_blank">说明</a></p>
    <p>向百度提交站点地图 <a href="http://zhanzhang.baidu.com/wiki/93#04" target="_blank">说明</a></p>
</div><!-- End .well -->
