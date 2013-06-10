<input type="hidden" name="location" value="add-category" />
{{$breadcrumb := breadcrumb "分类管理" "添加分类"}}
{{str2html $breadcrumb}}
<div class="alert">
    <button class="close" data-dismiss="alert" type="button">&times;</button>
    <span>{{.Message}}</span>
</div><!-- End .alert -->
<form method="post" class="form-horizontal add-category">
    <legend>添加分类</legend>
    <div class="control-group">
        <label class="control-label" for="name">分类名称</label>
        <div class="controls">
            <input type="text" id="title" name="name" placeholder="请输入分类名称" value="{{.Name}}" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="name_en">英文名称</label>
        <div class="controls">
            <input type="text" id="name_en" name="name_en" placeholder="请输入分类英文名称" value="{{.NameEn}}" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="description">分类描述</label>
        <div class="controls">
            <textarea id="description" name="description" palceholder="请输入分类描述">{{.Description}}</textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="alias">分类别名</label>
        <div class="controls">
            <input type="text" id="alias" name="alias" placeholder="请输入分类别名，别名之间用`, `分隔" class="input-xxlarge" value="{{.Alias}}" />
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">添加分类</button>
        </div>
    </div>
</form><!-- End .add-category -->
