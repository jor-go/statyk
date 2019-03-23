package main

import (
	"fmt"
	"log"
	"os"
	"statyk/internal/build"
	"statyk/internal/initialize"
	"statyk/internal/new"
	"statyk/internal/serve"
	"statyk/internal/upload"
)

/*APP_VERSION : Is the current statyk version*/
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
	if len(os.Args) < 2 {
		log.Fatalln("Not Enough Arguments to statyk")
	}
	action := os.Args[1]

	switch action {
	case INIT:
		initialize.Initialize()
		break
	case BUILD:
		build.Build(true)
		break
	case SERVE:
		build.Build(false)
		serve.Serve()
		break
	case UPLOAD:
		upload.Upload()
		break
	case NEW:
		new.New(os.Args)
	case VERSION:
		fmt.Println(APP_VERSION)
	default:
		fmt.Println("Command not recognized...")
	}
}
