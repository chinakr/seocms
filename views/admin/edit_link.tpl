<input type="hidden" name="location" value="edit-link" />
{{$breadcrumb := breadcrumb "友链管理" "修改友链"}}
{{str2html $breadcrumb}}
<div class="alert">
    <button class="close" data-dismiss="alert" type="button">&times;</button>
    <span>{{.Message}}</span>
</div><!-- End .alert -->
{{with .Link}}
<form method="post" class="edit-link form-horizontal">
    <legend>修改友链</legend>
    <div class="control-group">
        <label class="control-label" for="name">网站名称</label>
        <div class="controls">
            <input type="text" id="name" name="name" placeholder="请输入网站名称" value="{{.Name}}" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="url">URL</label>
        <div class="controls">
            <input type="text" id="url" name="url" placeholder="请输入友情链接URL" class="input-xxlarge" value="{{.Url}}" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="description">备注</label>
        <div class="controls">
            <textarea id="description" name="description" palceholder="请输入备注信息">{{.Description}}</textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">修改友链</button>
        </div>
    </div>
</form><!-- End .edit-link -->
{{end}}
