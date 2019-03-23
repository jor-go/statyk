package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"statyk/internal/utils"

	"github.com/gorilla/mux"
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

/*Serve : Serves the site locally*/
func Serve() {
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
