package linkparser

import (
	"fmt"
	"log"
	"os"
	"strings"

	xhtml "golang.org/x/net/html"
)

type LinkElement struct {
	Href    string
	Content string
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

func isAnchorTag(n *xhtml.Node) bool {
	return n.Type == xhtml.ElementNode && n.Data == "a"
}

func isElementNode(n *xhtml.Node) bool {
	return n.Type == xhtml.ElementNode && (n.Data == "strong" || n.Data == "div" || n.Data == "span")
}

func isTextTag(n *xhtml.Node) bool {
	return n.Type == xhtml.TextNode
}

func sanitizeText(data string) string {
	content := strings.TrimSpace(strings.Replace(data, "\n", "", -1))
	if content == "" {
		return ""
	}
	return content
}

func getAttr(n *xhtml.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getText(n *xhtml.Node, key string, isNearElement bool) string {
	parentHref := getAttr(n.Parent, "href")
	if parentHref != key && !isNearElement {
		return ""
	}

	var text strings.Builder

	text.WriteString(sanitizeText(n.Data))

	if n.NextSibling != nil && isElementNode(n.NextSibling) && n.NextSibling.LastChild != nil {
		if !isNearElement {
			text.WriteString(" ")
		}
		text.WriteString(getText(n.NextSibling.LastChild, key, true))
	}

	return text.String()
}

func setHrefKeys(doc *xhtml.Node) map[string]string {
	mapLinks := make(map[string]string, 0)
	for n := range doc.Descendants() {
		if isAnchorTag(n) {
			href := getAttr(n, "href")
			mapLinks[href] = ""
		}
	}
	return mapLinks
}

func setContent(doc *xhtml.Node, m map[string]string) {
	for n := range doc.Descendants() {
		for key, _ := range m {
			if isTextTag(n) {
				text := getText(n, key, false)
				if value, ok := m[key]; ok && text != "" {
					newValue := value + text
					m[key] = newValue
				}
			}
		}
	}
}

func convertToLinkElementSlice(m map[string]string) []LinkElement {
	links := make([]LinkElement, 0)
	for href, content := range m {
		links = append(links, LinkElement{
			Href:    href,
			Content: content,
		})
	}
	return links
}

func getLinks(doc *xhtml.Node) []LinkElement {

	links := []LinkElement{}

	mapLinks := setHrefKeys(doc)

	setContent(doc, mapLinks)

	links = convertToLinkElementSlice(mapLinks)

	return links
}

func Parse(fileHtml *string) []LinkElement {
	doc, _ := openFileAndParseDoc(fileHtml)

	log.Println("Init parsing...")

	return getLinks(doc)
}
