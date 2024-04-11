package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	alpha := []byte("Hello Victor!")
	f, err := os.Create("./one/alpha.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Write(alpha)

	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Open("./one/alpha.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f2.Close()

	bs, err := io.ReadAll(f2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bs))

	// key := "Victor"
	// text := "My Name is Victor"

	// currentPosition := int64(len(text))
	// var lastKeywordPosition int64

	// if strings.Contains(strings.ToLower(text), strings.ToLower(key)) {
	// 	lastKeywordPosition += (currentPosition - int64(len(text)))
	// }

	// fmt.Println("\n", lastKeywordPosition)
}
