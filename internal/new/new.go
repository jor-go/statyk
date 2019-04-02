package new

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"statyk/internal/initialize"
	"statyk/internal/things"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	//POST : create new post
	POST = "post"
)

func newPost(title, path string) {
	postURL := strings.ToLower(title)
	postURL = strings.Replace(postURL, " ", "-", -1)

	var pConfig things.PostConfig
	pConfig.Date = time.Now().Format("1/2/2006")
	pConfig.Markdown = postURL + ".md"
	pConfig.Title = title
	pConfig.URL = postURL

	pConfigYaml, err := yaml.Marshal(&pConfig)
	if err != nil {
		log.Fatalln(err)
	}

	configPath := filepath.Join(path, "posts", pConfig.URL+".yml")
	markdownPath := filepath.Join(path, "markdown", pConfig.Markdown)

	ioutil.WriteFile(configPath, pConfigYaml, os.ModePerm)
	ioutil.WriteFile(markdownPath, []byte(initialize.DefaultMarkdown), os.ModePerm)

	fmt.Println("NEW", configPath)
	fmt.Println("NEW", markdownPath)
}

// New Creates a new page*/
func New(args []string) {
	if len(args) < 3 {
		log.Fatalln("Not enough arguements to 'new'")
	}

	newType := args[2]

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch newType {
	case POST:
		if len(args) < 4 {
			newPost("New Post", wd)
		} else {
			newPost(args[3], wd)
		}
	default:
		log.Fatalln("Invalid arguement to 'new")
	}
}
