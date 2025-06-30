package main

import (
	"fmt"
	linkparser "htmllinkparser"
	"net/http"
)

func main() {

	res, err := http.Get("https://www.calhoun.io/")
	if err != nil {
		panic(err)
	}

	// content, errRead := io.ReadAll(res.Body)
	links, errRead := linkparser.ParseReader(res.Body)
	defer res.Body.Close()
	if errRead != nil {
		panic(fmt.Errorf("could not read page %v", errRead))
	}

	for _, link := range links {
		fmt.Printf("%s -> %s \n", link.Href, link.Content)
	}

	// fmt.Println(string(content))
}
