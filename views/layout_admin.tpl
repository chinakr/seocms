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
</head>
<body>
<div class="container">
    <div class="row">
        <div class="header">
            <div class="navbar">
                <div class="navbar-inner">
                    <a class="brand" href="/admin/">SEOCMS文章管理</a>
                </div>
                <ul class="nav">
                </ul>
            </div><!-- End .navbar -->
        </div><!-- End .header -->
        <div class="content">
            {{.LayoutContent}}
        </div><!-- End .content -->
        <div class="footer">
            <p>&copy; 2013 Developed by <a href="http://weibo.com/chinakr" target="_blank">chinakr</a></p>
        </div><!-- End .footer -->
    </div><!-- End .row -->
</div><!-- End .container -->
</body>
</html>
