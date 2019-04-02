package initialize

import "statyk/internal/things"

// DefaultHome is the content of the default home template
const DefaultHome = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="description" content="">
    <link rel="canonical" href="{{.SiteConfig.HomeLocation}}">
    <link rel="stylesheet" type="text/css" href="{{.SiteConfig.StyleLocation}}">
    <title>{{.SiteConfig.Name}}</title>
</head>
<body>
    {{template "header" .}}
    {{range .Posts}}
        {{.Config.Title}}
        {{index .Config.Custom "main-img"}}
    {{end}}
    {{template "footer" .}}
</body>
</html>
`

// DefaultGeneral is the contents of the default general html file
const DefaultGeneral = `
{{define "header"}}

{{end}}

{{define "footer"}}

{{end}}
`

// DefaultPost is the contents of the default html template
const DefaultPost = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="description" content="{{.Config.Description}}">
    <link rel="stylesheet" type="text/css" href="{{.SiteConfig.StyleLocation}}">
    <title>{{.Config.Title}}</title>
</head>
<body>
	<div class="">{{.HTML}}</div>
</body>
</html>
`

// DefaultStyle is the content for the default .sass file
const DefaultStyle = `
body
    margin: 0
`

// DefaultMarkdown is the content for the default markdown file
const DefaultMarkdown = `
# New Post

This is a new post

[Link](http://example.com)

![Some Image](http://example.com/image.jpg)
`

// DefaultPostConfig is the default config file
var DefaultPostConfig = things.PostConfig{
	Title:       "New Post",
	URL:         "new-post",
	Date:        "1/2/2006",
	Markdown:    "new-post.md",
	Description: "This is an example of a post description. This will show up in meta tags",
	Custom: map[string]string{
		"main-img": "https://example.com/image.jpg",
	},
}
