package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	keyword := "HTTP"

	f, err := os.Open("./seeker/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	var lastKeywordPosition int64
	var currentPosition int64

	for {
		line, err := reader.ReadString('\n')
		currentPosition += int64(len(line))

		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			lastKeywordPosition = (currentPosition - int64(len(line)))
		}

		if err == io.EOF {
			break
			// break out of the while loop if end of file reached
		}

		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

	}

	if lastKeywordPosition != 0 {
		_, err = f.Seek(lastKeywordPosition, io.SeekStart) // this adds a cursor to the file
		if err != nil {
			fmt.Println("error seeking file", err)
			return
		}

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("error reading file: ", err)
			return
		}

		fmt.Println(line)
	} else {
		fmt.Println("Not Found")
	}

	// _, _ = f.Seek(5, io.SeekStart)
	// _, _ = io.CopyN(os.Stdout, f, 10)

}
