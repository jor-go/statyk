Statyk
======

Statyk is a simple easy to use statyk site generator written in go.

Build
=====

### Get Dependancies
`dep ensure`

### Build
`go build`

Commands
========

## init
run with `statyk init`
This initialized a new project in the current directory

```
/assets
  /main.sass
/build
/config.dev.yml
/config.prod.yml
/markdown
  /new-post.md
/posts
  /new-post.yml
/templates
  /general.html
  /home.html
  /post.html
```

* `/assets` contains your `main.sass` which holds the styles for your site.
* `/build` is where your files will go after they are built.
* You have two config files one for development and one for production.
* `markdown` holds the markdown for your posts and will be converted to html on build.
* `/posts` contains the configuration for each post.
* `templates` holds the html templates.
    - `general.html` holds defined templates that are used in other places like header and footers.
    - `home.html` is the templates for your homepage.
    - `post.html` is the template for each post and where your markdown files will be places once they are converted to html.

## build
Run `statyk build` to build your site with production configuration.
This will place your compiled html files and compiled css file in the `/build` directory.

## serve
Run `statyk serve` to build your site with development configuration and start a file server on the port specified in your `config.dev.yml` file.

## new
Run `statyk new post "Some New Post Name"` to generate the files for a new post with specified title.
If you don't provide a title, it will be called "New Post" by default

## upload
This command is used to upload site files to AWS S3.
`statyk upload` uses your `~/.aws/credentials` and `~/.aws/config` to determine keys and region.
Run `statyk upload "filename" -bucket "bucket name"` to upload file to a certain S3 bucket.
Run `statyk upload -all -bucket "bucket name"` to upload the contents of a directory to a bucket (This will exclude sub-directories).
Run `statyk upload "image.jpg" -image -bucket "bucket name"` to upload image to bucket. This will take the md5 hash of the image as the name. It will print the new name after uploading.
Adding the `-directory "foo"` flag to an upload command will prepend a "directory" to the file name for S3. ex. "images/image.jpg"

### TODO

* Js handling
* Create Manifest.json?