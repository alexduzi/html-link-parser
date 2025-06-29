package linkparser

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	xhtml "golang.org/x/net/html"
)

type LinkElement struct {
	Href    string
	Content string
}

func isTextNode(n *xhtml.Node) bool {
	return n.Type == xhtml.TextNode
}

func isAnchorTag(n *xhtml.Node) bool {
	return n.Type == xhtml.ElementNode && n.Data == "a"
}

func sanitizeText(text string) string {
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(text, " "))
}

func getAttr(n *xhtml.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getContent(doc *xhtml.Node) string {
	var text strings.Builder

	var loopNode func(*xhtml.Node)
	loopNode = func(n *xhtml.Node) {
		if isTextNode(n) {
			text.WriteString(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			loopNode(c)
		}
	}

	loopNode(doc)
	return text.String()
}

func extractLinks(doc *xhtml.Node) []LinkElement {
	var links []LinkElement

	var loopDoc func(*xhtml.Node)
	loopDoc = func(n *xhtml.Node) {
		if isAnchorTag(n) {
			href := getAttr(n, "href")
			content := getContent(n)
			links = append(links, LinkElement{
				Href:    sanitizeText(href),
				Content: sanitizeText(content),
			})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			loopDoc(c)
		}
	}

	loopDoc(doc)
	return links
}

func ParseFile(filePath string) ([]LinkElement, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	return ParseReader(file)
}

func ParseReader(r io.Reader) ([]LinkElement, error) {
	doc, err := xhtml.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	return extractLinks(doc), nil
}
