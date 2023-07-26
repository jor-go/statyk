package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	env := os.Args[1]
	outputFile := os.Args[2]

	val := os.Getenv(env)
	dec, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(dec)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wrote", outputFile)
}
