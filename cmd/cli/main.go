package main

import (
	"flag"
	"fmt"
	linkparser "htmllinkparser"
	"log"
	"time"
)

func main() {
	start := time.Now()

	log.Println("Start...")

	fileHtml := flag.String("file", "", "Html file path")
	flag.Parse()

	log.Printf("File: %s \n", *fileHtml)

	links, _ := linkparser.ParseFile(*fileHtml)

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
	elapsed := time.Since(start)

	log.Printf("Parsing took %s", elapsed)
}
