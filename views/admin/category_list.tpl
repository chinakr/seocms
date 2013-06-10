<input type="hidden" name="location" value="category-list" />
{{$breadcrumb := breadcrumb "分类管理" "分类列表"}}
{{str2html $breadcrumb}}
<table class="table table-striped table-hover category-list">
<thead>
    <tr><th>分类名称</th><th>英文名称</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Categories}}
    <tr><td>{{.Name}}</td><td>{{.NameEn}}</td><td><span class="action"><a href="/category/edit/{{.Id}}">修改</a></span><span class="action"><a href="/category/delete/{{.Id}}">删除</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- .category-list -->
