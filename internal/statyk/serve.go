package statyk

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"statyk/internal/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var wd string

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["filename"]

	path := filepath.Join(wd, "build", file)

	http.ServeFile(w, r, path)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(wd, "build", "home")

	http.ServeFile(w, r, path)
}

// serve Serves the site locally
func serve() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// get config file
	config := utils.YamlToConfig(filepath.Join(wd, "config.dev.yml"))

	fmt.Println("Now serving", config.Name, "on port :", config.Port)

	r := mux.NewRouter()

	// Main Handlers
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/{filename}", fileServerHandler)

	http.Handle("/", r)
	http.ListenAndServe(":"+config.Port, r)
}

var ServeCmd = &cobra.Command{
	Use:     "serve",
	Short:   "serve runs a service to host the generated static files",
	Long:    `serve runs a service to host the generated static files`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}
