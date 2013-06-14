<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<title>{{.PageTitle}}</title>
<meta name="keywords" content="{{.PageKeywords}}" />
<meta name="description" content="{{.PageDescription}}" />
<link rel="stylesheet" href="http://lib.sinaapp.com/js/bootstrap/latest/css/bootstrap.min.css" />
<script src="http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js"></script>
<script src="http://lib.sinaapp.com/js/bootstrap/latest/js/bootstrap.min.js"></script>
<link rel="stylesheet" href="/static/css/global.css" />
<script src="/static/js/global.js"></script>
</head>
<body>
<div class="container admin">
    <div class="header">
        <div class="navbar">
            <ul class="nav pull-right">
                <li class="account"><a>{{.Account}}</a></li>
                <li><a href="/user/logout">退出</a></li>
            </ul>
            <div class="navbar-inner">
                <a class="brand" href="/admin">{{.SiteName}}后台管理</a>
            </div>
        </div><!-- End .navbar -->
    </div><!-- End .header -->
    <div class="content">
        <div class="row">
            <div class="sidebar span2">
                <div class="well">
                    <ul class="nav nav-list">
                        <li class="nav-header">网站管理</li>
                        <li class="edit-head"><a href="/site/head">首页head</a></li>
                        <li class="nav-header">文章管理</li>
                        <li class="article-list"><a href="/article/list">文章列表</a></li>
                        <li class="add-article"><a href="/article/add">添加文章</a></li>
                        <li class="nav-header">分类管理</li>
                        <li class="category-list"><a href="/category/list">分类列表</a></li>
                        <li class="add-category"><a href="/category/add">添加分类</a></li>
                        <li class="nav-header">标签管理</li>
                        <li class="tag-list"><a href="/tag/">标签列表</a></li>
                        <li class="nav-header">友链管理</li>
                        <li class="link-list"><a href="/link/">友链列表</a></li>
                        <li class="add-link"><a href="/link/add">添加友链</a></li>
                        <li class="nav-header">用户管理</li>
                        <li class="user-list"><a href="/user/">用户列表</a></li>
                        <li class="add-user"><a href="/user/add">添加用户</a></li>
                    </ul><!-- End .nav -->
                </div><!-- End .well -->
            </div><!-- End .sidebar -->
            <div class="main span10">
                {{.LayoutContent}}
            </div><!-- End .main -->
        </div><!-- End .row -->
    </div><!-- End .content -->
    <div class="footer">
        <p>&copy; 2013 Developed by <a href="http://weibo.com/chinakr" target="_blank">chinakr</a></p>
    </div><!-- End .footer -->
</div><!-- End .container -->
</body>
</html>
