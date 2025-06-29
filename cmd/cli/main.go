package main

import (
	"flag"
	"fmt"
	"htmllinkparser/linkparser"
	"log"
)

func main() {
	log.Println("Start...")

	fileHtml := flag.String("file", "", "Html file path")
	flag.Parse()

	log.Printf("File: %s \n", *fileHtml)

	links := linkparser.Parse(*fileHtml)

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
