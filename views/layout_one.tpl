<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<title>{{.PageTitle}}</title>
<meta name="keywords" content="{{.PageKeywords}}" />
<meta name="description" content="{{.PageDescription}}" />
<link href="/rss/" rel="alternate" title="文章订阅" type="application/rss+xml" />
<link rel="stylesheet" href="/static/css/bootstrap.min.css" />
<script src="/static/js/jquery-2.0.3.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" href="/static/css/global.css" />
<script src="/static/js/global.js"></script>
</head>
<body>
<div class="container">
    <div class="header">
        <div class="navbar navbar-default" role="navigation">
            <div class="navbar-header">
                <a class="navbar-brand" href="/">{{.SiteName}}</a>
            </div><!-- End .navbar-header -->
            <div class="navbar-collapse collapse navbar-ex1-collapse">
                <ul class="nav navbar-nav">
                    <li class="index"><a href="/">首页</a></li>
                    {{range .Categories}}
                    <li class="{{.NameEn}}"><a href="/{{.NameEn}}/">{{.Name}}</a></li>
                    {{end}}
                </ul><!-- End .nav -->
            </div><!-- End .navbar-collapse -->
        </div><!-- End .navbar -->
    </div><!-- End .header -->
    <div class="content">
        <div class="main">
            {{.LayoutContent}}
        </div><!-- End .main -->
    </div><!-- End .content -->
    <div class="footer">
        {{template "footer.tpl"}}
    </div><!-- End .footer -->
</div><!-- End .container -->
</body>
</html>
