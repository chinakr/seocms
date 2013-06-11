{{$breadcrumb := breadcrumb "用户管理" "添加用户"}}
{{str2html $breadcrumb}}
<form method="post" class="form-horizontal add-user">
    <legend>添加用户</legend>
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
            <input type="password" id="password" name="password" value="{{.Password}}" />
        </div>
    </div>
    <div class="control-group">
        <label class="control-label" for="re-password">请重复输入密码</label>
        <div class="controls">
            <input type="password" id="re-password" name="re-password" />
        </div>
    </div>
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">添加用户</button>
        </div>
    </div>
</form><!-- End .add-user -->
