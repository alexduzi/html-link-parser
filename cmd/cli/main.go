package main

import (
	"flag"
	"htmllinkparser/linkparser"
	"log"
)

func main() {
	log.Println("Start...")

	fileHtml := flag.String("file", "", "Html file path")
	flag.Parse()

	log.Printf("File: %s \n", *fileHtml)

	linkparser.Parse(fileHtml)
}
