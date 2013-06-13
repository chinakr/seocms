<input type="hidden" name="location" value="tag-list" />
{{$breadcrumb := breadcrumb "标签管理" "标签列表"}}
{{str2html $breadcrumb}}
<table class="tag-list table table-striped table-hover">
<thead>
    <tr><th>标签名称</th><th>英文名称</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Tags}}
    <tr><td>{{.Name}}</td><td>{{.NameEn}}</td><td><span class="action"><a href="/tag/edit/{{.Id}}">修改</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- .category-list -->
<div class="pagination">
    <ul>
        {{str2html .Paginator}}
    </ul>
</div><!-- End .pagination -->
