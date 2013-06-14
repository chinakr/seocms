<input type="hidden" name="location" value="edit-head" />
{{$breadcrumb := breadcrumb "网站管理" "首页head管理"}}
{{str2html $breadcrumb}}
<div class="alert">
    <button class="close" data-dismiss="alert" type="button">&times;</button>
    <span>{{.Message}}</span>
</div><!-- End .alert -->
{{with .Site}}
<form method="post" class="edit-head form-horizontal">
    <legend>修改首页head</legend>
    <div class="control-group">
        <label class="control-label" for="content">首页head</label>
        <div class="controls">
            <textarea id="content" name="content" palceholder="请输入首页head">{{.Content}}</textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">修改head</button>
        </div>
    </div>
</form><!-- End .edit-head -->
{{end}}
