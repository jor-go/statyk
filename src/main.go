package main

import (
	"fmt"
	"log"
	"os"
	"statyk/src/build"
	"statyk/src/initialize"
	"statyk/src/new"
	"statyk/src/serve"
	"statyk/src/upload"
)

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
	default:
		fmt.Println("Command not recognized...")
	}
}
