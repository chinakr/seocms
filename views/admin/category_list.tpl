<input type="hidden" name="location" value="category-list" />
<ul class="breadcrumb">
    <li><a href="/article/list">管理后台</a> <span class="divider">/</span></li>
    <li><a href="/article/list">文章管理</a> <span class="divider">/</span>
    <li class="active">分类列表</li>
</ul><!-- End .breadcrumb -->
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
