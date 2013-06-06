<div class="sidebar">
    <div class="recommends">
        <h2>推荐阅读</h2>
        <ul>
            {{range .Articles}}
            <li>· <a href="/{{id2categoryEn .Category}}/{{.Id}}" target="_blank">{{.Title}}</a></li>
            {{end}}
        </ul>
    </div><!-- End .recommends-cloud -->
</div><!-- End .sidebar -->
