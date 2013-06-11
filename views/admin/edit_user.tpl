<input type="hidden" name="location" value="edit-user" />
{{$breadcrumb := breadcrumb "用户管理" "修改用户"}}
{{str2html $breadcrumb}}
{{with .User}}
<form method="post" class="edit-user form-horizontal">
    <legend>修改用户</legend>
    <div class="control-group">
        <label class="control-label" for="email">E-mail</label>
        <div class="controls">
            <input type="email" id="email" name="email" placeholder="请输入E-mail地址" value="{{.Email}}" />
        </div>
    </div>
    <div class="control-group">
        <label class="control-label" for="name">用户名</label>
        <div class="controls">
            <input type="text" id="name" name="name" placeholder="请输入用户名" value="{{.Name}}" />
        </div>
    </div>
    <div class="control-group">
        <label class="control-label" for="password">密码</label>
        <div class="controls">
            <input type="password" id="password" name="password" value="" />
        </div>
    </div>
    <div class="control-group">
        <label class="control-label" for="re-password">请重复输入密码</label>
        <div class="controls">
            <input type="password" id="re-password" name="re-password" value="" />
        </div>
    </div>
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">修改用户</button>
        </div>
    </div>
</form><!-- End .edit-user -->
{{end}}
