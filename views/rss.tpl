<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
    <channel>
        {{with .}}
        <title>{{.Title}}</title>
        <link>{{.Link}}</link>
        <description>{{.Description}}</description>
        <generator>SEOCMS</generator>
        {{range .Items}}
        <item>
            <title>{{.Title}}</title>
            <link>{{.Link}}</link>
            <category>{{.Category}}</category>
            <description><![CDATA[{{.Description}}]]></description>
            <pubDate>{{.Pubdate}}</pubDate>
            <guid>{{.Guid}}</guid>
        </item>
        {{end}}
        {{end}}
    </channel>
</rss>
