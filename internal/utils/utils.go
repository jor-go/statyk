package utils

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"statyk/internal/things"

	blackfriday "gopkg.in/russross/blackfriday.v2"
	yaml "gopkg.in/yaml.v2"
)

// YamlToPost Convert Yaml file to a postconfig struct
func YamlToPost(path string) (p things.PostConfig) {
	postFile, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln(path, " does not exist")
		} else {
			log.Fatalln(err)
		}
	}

	err = yaml.Unmarshal(postFile, &p)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

// YamlToConfig Convert Yaml file to a siteConfig struct
func YamlToConfig(path string) (s things.SiteConfig) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln(path, " does not exist")
		} else {
			log.Fatalln(err)
		}
	}

	err = yaml.Unmarshal(configFile, &s)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

// MarkdownToHTML Converts markdown file to HTML
func MarkdownToHTML(path string) (t template.HTML) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	htmlBytes := blackfriday.Run(file)
	t = template.HTML(htmlBytes)
	return
}
