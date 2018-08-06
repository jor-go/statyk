package initialize

import "statyk/src/things"

const DefaultHome = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" type="text/css" href="{{.SiteConfig.StyleLocation}}/main.css">
    <title>{{.SiteConfig.Name}}</title>
</head>
<body>
    {{template "header" .}}
    {{range .Posts}}
        {{.Config.Title}}
    {{end}}
    {{template "footer" .}}
</body>
</html>
`

const DefaultGeneral = `
{{define "header"}}

{{end}}

{{define "footer"}}

{{end}}
`

const DefaultPost = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" type="text/css" href="{{.SiteConfig.StyleLocation}}/main.css">
    <title>{{.Config.Title}}</title>
</head>
<body>
	<div class="">{{.HTML}}</div>
</body>
</html>
`

const DefaultStyle = `
body
    margin: 0
`

const DefaultMarkdown = `
# New Post

This is a new post

[Link](http://example.com)

![Some Image](http://example.com/image.jpg)
`

var DefaultPostConfig = things.PostConfig{
	Title:    "New Post",
	URL:      "new-post",
	Date:     "1/2/2006",
	Markdown: "new-post.md",
}
