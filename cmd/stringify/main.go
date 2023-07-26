package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := base64.StdEncoding.EncodeToString(b)

	fmt.Println("FILE:", os.Args[1])
	fmt.Println(s)

	//

	// outbytes, err := base64.StdEncoding.DecodeString(s)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(outbytes))
}
