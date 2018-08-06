package initialize

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"statyk/src/things"

	"gopkg.in/yaml.v2"
)

func generateDirs(path string) {
	// generate build directory
	os.Mkdir(filepath.Join(path, "build"), os.ModePerm)

	// generate templates directory and starter templates
	os.Mkdir(filepath.Join(path, "templates"), os.ModePerm)
	ioutil.WriteFile(filepath.Join(path, "templates", "home.html"), []byte(DefaultHome), os.ModePerm)
	ioutil.WriteFile(filepath.Join(path, "templates", "general.html"), []byte(DefaultGeneral), os.ModePerm)
	ioutil.WriteFile(filepath.Join(path, "templates", "post.html"), []byte(DefaultPost), os.ModePerm)

	// generate posts directory
	os.Mkdir(filepath.Join(path, "posts"), os.ModePerm)

	// generate assets directory
	os.Mkdir(filepath.Join(path, "assets"), os.ModePerm)
	ioutil.WriteFile(filepath.Join(path, "assets", "main.sass"), []byte(DefaultStyle), os.ModePerm)

	// generate markdown directory
	os.Mkdir(filepath.Join(path, "markdown"), os.ModePerm)
}

func generateConfigs(path string) {
	// Generate default dev config
	defaultDEVConfig := things.SiteConfig{
		Name:          "New Site",
		StyleLocation: "http://localhost:8080/assets/main.sass",
		HomeLocation:  "http://localhost:8080",
		Port:          "8080",
	}
	defaultDEVConfigYAML, err := yaml.Marshal(&defaultDEVConfig)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(filepath.Join(path, "config.dev.yml"), defaultDEVConfigYAML, os.ModePerm)

	// Generate default prod config
	defaultPRODConfig := things.SiteConfig{
		Name:          "New Site",
		StyleLocation: "https://cdn.example.com/main.sass",
		HomeLocation:  "http://example.com",
	}
	defaultPRODConfigYAML, err := yaml.Marshal(&defaultPRODConfig)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(filepath.Join(path, "config.prod.yml"), defaultPRODConfigYAML, os.ModePerm)
}

func generateExamplePost(path string) {

	defaultPostConfigYAML, err := yaml.Marshal(&DefaultPostConfig)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(filepath.Join(path, "posts", "new-post.yml"), defaultPostConfigYAML, os.ModePerm)
	ioutil.WriteFile(filepath.Join(path, "markdown", "new-post.md"), []byte(DefaultMarkdown), os.ModePerm)
}

/*Initialize : Creates new site in the current directory*/
func Initialize() {
	// get working directory
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// must be completed first
	generateDirs(ex)

	generateExamplePost(ex)

	generateConfigs(ex)
}
