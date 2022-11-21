package statyk

import (
	"log"
	"os"
	"path/filepath"
	"statyk/internal/things"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func generateDirs(path string) {
	// generate build directory
	os.Mkdir(filepath.Join(path, "build"), os.ModePerm)

	// generate templates directory and starter templates
	os.Mkdir(filepath.Join(path, "templates"), os.ModePerm)
	os.WriteFile(filepath.Join(path, "templates", "home.html"), []byte(DefaultHome), os.ModePerm)
	os.WriteFile(filepath.Join(path, "templates", "general.html"), []byte(DefaultGeneral), os.ModePerm)
	os.WriteFile(filepath.Join(path, "templates", "post.html"), []byte(DefaultPost), os.ModePerm)

	// generate posts directory
	os.Mkdir(filepath.Join(path, "posts"), os.ModePerm)

	// generate assets directory
	os.Mkdir(filepath.Join(path, "assets"), os.ModePerm)
	os.WriteFile(filepath.Join(path, "assets", "main.scss"), []byte(DefaultStyle), os.ModePerm)

	// generate markdown directory
	os.Mkdir(filepath.Join(path, "markdown"), os.ModePerm)
}

func generateConfigs(path string) {
	// Generate default dev config
	defaultDEVConfig := things.SiteConfig{
		Name:          "New Site",
		StyleLocation: "/main.css",
		HomeLocation:  "/",
		Port:          "8080",
	}
	defaultDEVConfigYAML, err := yaml.Marshal(&defaultDEVConfig)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(filepath.Join(path, "config.dev.yml"), defaultDEVConfigYAML, os.ModePerm)

	// Generate default prod config
	defaultPRODConfig := things.SiteConfig{
		Name:          "New Site",
		StyleLocation: "/main.css",
		HomeLocation:  "/",
	}
	defaultPRODConfigYAML, err := yaml.Marshal(&defaultPRODConfig)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(filepath.Join(path, "config.prod.yml"), defaultPRODConfigYAML, os.ModePerm)
}

func generateExamplePost(path string) {

	defaultPostConfigYAML, err := yaml.Marshal(&DefaultPostConfig)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(filepath.Join(path, "posts", "new-post.yml"), defaultPostConfigYAML, os.ModePerm)
	os.WriteFile(filepath.Join(path, "markdown", "new-post.md"), []byte(DefaultMarkdown), os.ModePerm)
}

// initialize Creates new site in the current directory
func initialize() {
	// get working directory
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// check if project already exists
	configPath := filepath.Join(ex, "config.prod.yml")
	if _, err := os.Stat(configPath); err == nil {
		log.Fatalln("Project already initialized")
	}

	// must be completed first
	generateDirs(ex)

	generateExamplePost(ex)

	generateConfigs(ex)
}

var InitCmd = &cobra.Command{
	Use:     "init",
	Short:   "init starts a new Statyk project in the current directory",
	Long:    `init starts a new Statyk project in the current directory`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}
