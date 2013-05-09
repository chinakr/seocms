<input type="hidden" name="location" value="add-article" />
<ul class="breadcrumb">
    <li><a href="/article/list">管理后台</a> <span class="divider">/</span></li>
    <li><a href="/article/list">文章管理</a> <span class="divider">/</span>
    <li class="active">添加文章</li>
</ul><!-- End .breadcrumb -->
<form method="post" class="form-horizontal add-article">
    <legend>添加文章</legend>
    <div class="control-group">
        <label class="control-label" for="title">文章标题</label>
        <div class="controls">
            <input type="text" id="title" name="title" placeholder="请输入文章标题" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="pubdate">发布日期</label>
        <div class="controls">
            <input type="date" id="pubdate" name="pubdate" placeholder="请输入发布日期" />
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="abstract">文章摘要</label>
        <div class="controls">
            <textarea id="abstract" name="abstract" placeholder="请输入文章摘要"></textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="content">文章内容</label>
        <div class="controls">
            <textarea id="content" name="content" placeholder="请输入文章内容" rows="10"></textarea>
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="category">文章分类</label>
        <div class="controls">
            <select>
                {{range .Categories}}
                <option>{{.}}</option>
                {{end}}
            </select>
            (必填)
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <label class="control-label" for="tags">文章标签</label>
        <div class="controls">
            <input type="text" id="tags" name="tags" placeholder="请输入文章标签，标签之间用`, `分隔" class="input-xxlarge" />
        </div>
    </div><!-- End .control-group -->
    <div class="control-group">
        <div class="controls">
            <button type="submit" class="btn btn-primary">添加文章</button>
        </div>
    </div>
</form><!-- End .add-article -->
