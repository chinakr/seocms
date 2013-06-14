<input type="hidden" name="location" value="edit-body" />
{{$breadcrumb := breadcrumb "网站管理" "通用body管理"}}
{{str2html $breadcrumb}}
<div class="alert">
    <button class="close" data-dismiss="alert" type="button">&times;</button>
    <span>{{.Message}}</span>
</div><!-- End .alert -->
{{with .Site}}
<form method="post" class="edit-body form-horizontal">
    <legend>修改通用body</legend>
    <div class="control-group">
        <label class="control-label" for="content">通用body</label>
        <div class="controls">
            <textarea id="content" name="content" palceholder="请输入通用body">{{.Content}}</textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">修改body</button>
        </div>
    </div>
</form><!-- End .edit-body -->
{{end}}
