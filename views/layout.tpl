<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<title>{{.PageTitle}}</title>
<meta name="keywords" content="{{.PageKeywords}}" />
<meta name="description" content="{{.PageDescription}}" />
<link href="/rss/" rel="alternate" title="文章订阅" type="application/rss+xml" />
<link rel="stylesheet" href="http://lib.sinaapp.com/js/bootstrap/latest/css/bootstrap.min.css" />
<script src="http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js"></script>
<script src="http://lib.sinaapp.com/js/bootstrap/latest/js/bootstrap.min.js"></script>
<link rel="stylesheet" href="/static/css/global.css" />
<script src="/static/js/global.js"></script>
</head>
<body>
<div class="container">
    <div class="header">
        <div class="navbar">
            <div class="navbar-inner">
                <a class="brand" href="/">{{.SiteName}}</a>
                <ul class="nav">
                    <li class="index"><a href="/">首页</a></li>
                    {{range .Categories}}
                    <li class="{{.NameEn}}"><a href="/{{.NameEn}}/">{{.Name}}</a></li>
                    {{end}}
                </ul>
            </div><!-- End .navbar-inner -->
        </div><!-- End .navbar -->
    </div><!-- End .header -->
    <div class="content">
        <div class="row">
            <div class="main span9">
                {{.LayoutContent}}
            </div><!-- End .main -->
            <div class="sidebar span3">
                {{str2html .Sidebar}}
            </div><!-- End .sidebar -->
        </div><!-- End .row -->
    </div><!-- End .content -->
    <div class="footer">
        {{template "footer.tpl"}}
    </div><!-- End .footer -->
</div><!-- End .container -->
{{str2html .Body}}
</body>
</html>
