<input type="hidden" name="location" value="user-list" />
{{$breadcrumb := breadcrumb "用户管理" "用户列表"}}
{{str2html $breadcrumb}}
<table class="user-list table table-striped table-hover">
<thead>
    <tr><th>用户名</th><th>E-mail</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Users}}
    <tr><td>{{.Name}}</td><td>{{.Email}}</td><td><span class="action"><a href="/user/edit/{{.Id}}">修改</a></span><span class="action"><a href="/user/delete/{{.Id}}">删除</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- End .user-list -->
