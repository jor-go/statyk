package main

import (
	"statyk/internal/statyk"
)

// APP_VERSION Is the current statyk version
const APP_VERSION = "0.0.1"

const (
	//INIT : Constant for init arg
	INIT = "init"
	//UPLOAD : Constant for upload arg
	UPLOAD = "upload"
	//SERVE : Constant for serve arg
	SERVE = "serve"
	//BUILD : Constant for build arg
	BUILD = "build"
	//NEW : Constant for the new arg
	NEW = "new"
	//VERSION : constant for the version arg
	VERSION = "version"
)

func main() {
	statyk.Execute()
}
