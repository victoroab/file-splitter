package main

import (
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("./copyier/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, _ = io.CopyN(os.Stdout, f, 10)

}
