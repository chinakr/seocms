<input type="hidden" name="location" value="article-list" />
<ul class="breadcrumb">
    <li><a href="/article/list">管理后台</a> <span class="divider">/</span></li>
    <li><a href="/article/list">文章管理</a> <span class="divider">/</span>
    <li class="active">文章列表</li>
</ul><!-- End .breadcrumb -->
<table class="table table-striped table-hover article-list">
<thead>
    <tr><th>分类</th><th>标题</th><th>发表日期</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Articles}}
    <tr><td>{{id2category .Category}}</td><td>{{.Title}}</td><td>{{dateformat .Pubdate "2006-01-02"}}</td><td><span class="action"><a href="/{{id2category .Category}}/{{.Id}}" target="_blank">查看</a></span><span class="action"><a href="/article/edit/{{.Id}}">修改</a></span><span class="action"><a href="/article/delete/{{.Id}}">删除</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- .article-list -->
