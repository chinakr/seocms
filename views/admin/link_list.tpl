<input type="hidden" name="location" value="link-list" />
{{$breadcrumb := breadcrumb "友链管理" "友链列表"}}
{{str2html $breadcrumb}}
<table class="link-list table table-striped table-hover">
<thead>
    <tr><th>网站名称</th><th>URL</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Tags}}
    <tr><td><a title="{{.Description}}">{{.Name}}</a></td><td><a href="{{.Url}}" target="_blank">{{.Url}}</a></td><td><span class="action"><a href="/link/edit/{{.Id}}">修改</a></span><span class="action"><a href="/link/delete/{{.Id}}">删除</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- .link-list -->
<div class="pagination">
    <ul>
        {{str2html .Paginator}}
    </ul>
</div><!-- End .pagination -->
