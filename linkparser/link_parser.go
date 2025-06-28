package linkparser

import (
	"fmt"
	"log"
	"os"
	"strings"

	xhtml "golang.org/x/net/html"
)

type LinkElement struct {
	Href string
	Text string
}

func openFileAndParseDoc(fileHtml *string) (*xhtml.Node, error) {
	file, err := os.Open(*fileHtml)
	if err != nil {
		panic(fmt.Sprintf("file %s does not exists", *fileHtml))
	}
	defer file.Close()

	doc, err := xhtml.Parse(file)

	return doc, err
}

func getAttr(n *xhtml.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getText(n *xhtml.Node) string {
	content := strings.TrimSpace(strings.Replace(n.Data, "\n", "", -1))
	if content == "" {
		return ""
	}

	var text strings.Builder

	text.WriteString(content)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text.WriteString(getText(c))
	}

	return text.String()
}

func IsAnchorTag(n *xhtml.Node) bool {
	return n.Type == xhtml.ElementNode && n.Data == "a"
}

func IsTextTag(n *xhtml.Node) bool {
	return n.Type == xhtml.TextNode
}

func getLinks(doc *xhtml.Node) *[]LinkElement {
	links := []LinkElement{}
	for n := range doc.Descendants() {
		href := ""
		text := ""
		if IsAnchorTag(n) {
			href = getAttr(n, "href")

		}
		if IsTextTag(n) {
			text = getText(n)
		}
		fmt.Println("href: " + href)
		fmt.Println("text: " + text)
		// links = append(links, LinkElement{Href: href, Text: text})
	}

	return &links
}

func Parse(fileHtml *string) {
	doc, _ := openFileAndParseDoc(fileHtml)

	log.Println("Init parsing...")

	links := getLinks(doc)

	for _, link := range *links {
		fmt.Printf("%+v\n", link)
	}
}
