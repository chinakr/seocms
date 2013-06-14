<div class="sidebar">
    <div class="tags-cloud">
        <span class="more"><a href="/tags/" target="_blank">更多</a></span><h2>热门话题</h2>
        <ul>
            {{range .Tags}}
            <li><a href="/t/{{.Id}}/" target="_blank">{{.Name}}</a></li>
            {{end}}
        </ul>
    </div><!-- End .tags-cloud -->
    <div class="friend-links">
        <h2>友情链接</h2>
        <ul>
            {{range .Links}}
            <li><a href="{{.Url}}" target="_blank">{{.Name}}</a></li>
            {{end}}
        </ul>
    </div><!-- End .friend-links -->
</div><!-- End .sidebar -->
