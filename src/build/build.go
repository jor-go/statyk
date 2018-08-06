package build

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"statyk/src/things"
	"statyk/src/utils"
)

var posts []things.Post

func generateHTMLFile(t *template.Template, path, name string, info interface{}) {
	// open new file location
	f, err := os.Create(path)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}
	defer f.Close()

	// create buffer
	fw := bufio.NewWriter(f)

	// run template
	err = t.ExecuteTemplate(fw, name, info)
	if err != nil {
		log.Fatalln(err)
	}

	// flush buffer
	fw.Flush()

	fmt.Println("HTML File Generated: ", path)
}

func generatePostFiles(c things.SiteConfig, wd string) {
	// get template
	tempDir := filepath.Join(wd, "templates")
	postTemplate := template.Must(
		template.ParseFiles(
			filepath.Join(tempDir, "post.html"),
			filepath.Join(tempDir, "general.html")))

	for _, post := range posts {
		generateHTMLFile(postTemplate, filepath.Join(wd, "build", post.Config.URL), "post.html", post)
	}
}

func generateHomepage(c things.SiteConfig, wd string) {
	var info things.MultiPost
	info.SiteConfig = c
	info.Posts = posts

	// get template
	tempDir := filepath.Join(wd, "templates")
	homeTemplate := template.Must(
		template.ParseFiles(
			filepath.Join(tempDir, "home.html"),
			filepath.Join(tempDir, "general.html")))

	generateHTMLFile(homeTemplate, filepath.Join(wd, "build", "home"), "home.html", info)
}

func GetPosts(c things.SiteConfig, wd string) (posts []things.Post) {
	// get all yaml files in posts
	files, err := filepath.Glob(filepath.Join(wd, "/posts/*.yml"))
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		var post things.Post
		post.Config = utils.YamlToPost(file)
		post.HTML = utils.MarkdownToHTML(filepath.Join(wd, "markdown", post.Config.Markdown))
		post.SiteConfig = c
		posts = append(posts, post)
	}
	return
}

/*Build : Builds a site based off config file*/
func Build(isProd bool) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	configFile := "config.prod.yml"
	if !isProd {
		configFile = "config.dev.yml"
	}

	// get config file
	config := utils.YamlToConfig(filepath.Join(workingDirectory, configFile))

	err = os.RemoveAll(filepath.Join(workingDirectory, "build"))
	if err != nil {
		log.Fatalln(err)
	}
	os.Mkdir(filepath.Join(workingDirectory, "build"), os.ModePerm)
	os.Mkdir(filepath.Join(workingDirectory, "build", "assets"), os.ModePerm)

	posts = GetPosts(config, workingDirectory)

	// generate HTML for all posts
	generatePostFiles(config, workingDirectory)

	// generate HTML for Homepage\s
	generateHomepage(config, workingDirectory)

	if isProd {
		utils.GenerateSitemap(posts, config, workingDirectory)
	}

	// css
	inputSass := filepath.Join(workingDirectory, "assets", "main.sass")
	outputSass := filepath.Join(workingDirectory, "build", "assets", "main.css")
	buildPath := inputSass + ":" + outputSass
	sass := exec.Command("sass", buildPath, "--style", "compressed")
	err = sass.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
