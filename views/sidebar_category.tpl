<div class="sidebar">
    <div class="tags-cloud">
        <h2>相关话题</h2>
        <ul>
            {{range .Tags}}
            <li><a href="/t/{{.Id}}/" target="_blank">{{.Name}}</a></li>
            {{end}}
        </ul>
    </div><!-- End .tags-cloud -->
</div><!-- End .sidebar -->
