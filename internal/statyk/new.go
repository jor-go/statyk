package statyk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"statyk/internal/things"
	"strings"
	"time"

	"github.com/spf13/cobra"
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

	os.WriteFile(configPath, pConfigYaml, os.ModePerm)
	os.WriteFile(markdownPath, []byte(DefaultMarkdown), os.ModePerm)

	fmt.Println("NEW", configPath)
	fmt.Println("NEW", markdownPath)
}

// generateNew Creates a new page*/
func generateNew(args []string) {
	if len(args) < 1 {
		log.Fatalln("Not enough arguements to 'new'")
	}

	newType := args[0]

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch newType {
	case POST:
		if len(args) < 2 {
			newPost("New Post", wd)
		} else {
			newPost(args[1], wd)
		}
	default:
		log.Fatalln("Invalid argument to 'new")
	}
}

var NewCmd = &cobra.Command{
	Use:     "new",
	Short:   "new [type] generates files for a new item",
	Long:    `new [type] generates files for a new item, that type can be a post`,
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		generateNew(args)
	},
}
