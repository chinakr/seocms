<input type="hidden" name="location" value="edit-tag" />
{{$breadcrumb := breadcrumb "标签管理" "修改标签"}}
{{str2html $breadcrumb}}
<div class="alert">
    <button class="close" data-dismiss="alert" type="button">&times;</button>
    <span>{{.Message}}</span>
</div><!-- End .alert -->
{{with .Tag}}
<form method="post" class="edit-tag form-horizontal">
    <legend>修改标签</legend>
    <div class="control-group">
        <label class="control-label" for="name">标签名称</label>
        <div class="controls">
            <input type="text" id="title" name="name" placeholder="请输入标签名称" value="{{.Name}}" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="name_en">英文名称</label>
        <div class="controls">
            <input type="text" id="name_en" name="name_en" placeholder="请输入标签英文名称" value="{{.NameEn}}" />
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="description">标签描述</label>
        <div class="controls">
            <textarea id="description" name="description" palceholder="请输入标签描述">{{.Description}}</textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="alias">标签别名</label>
        <div class="controls">
            <input type="text" id="alias" name="alias" placeholder="请输入标签别名，别名之间用`, `分隔" class="input-xxlarge" value="{{.Alias}}" />
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">修改标签</button>
        </div>
    </div>
</form><!-- End .edit-tag -->
{{end}}
