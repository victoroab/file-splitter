package main

import (
	"log"

	"github.com/victoroab/go-file-manipulation/filer/spliter"
)

func main() {
	// spliter.SplitFile("./filer/Concurrency_in_Go.pdf", 1068309)
	if err := spliter.JoinFiles("./filer/Concurrency_in_Go.pdf"); err != nil {
		log.Fatal(err)
	}
}
