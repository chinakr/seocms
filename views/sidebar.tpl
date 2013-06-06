<div class="sidebar">
    <div class="tags-cloud">
        <h2>热门话题</h2>
        <ul>
            {{range .Tags}}
            <li><a href="/t/{{.Id}}/" target="_blank">{{.Name}}</a></li>
            {{end}}
        </ul>
    </div><!-- End .tags-cloud -->
    <div class="friend-links">
        <h2>友情链接</h2>
        <ul>
            <li>友情链接1</li>
        </ul>
    </div><!-- End .friend-links -->
</div><!-- End .sidebar -->
