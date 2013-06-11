<form method="post" class="user-login form-horizontal">
    <legend>用户登录</legend>
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
        <div class="controls">
            <button type="submit" class="btn btn-primary">登录</button>
        </div>
    </div>
</form><!-- End .user-login -->
