package linkparser

import (
	"encoding/xml"
	"fmt"
	"net"
	"strings"
)

type UrlSiteMap struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type SiteMap struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	Urls    []UrlSiteMap `xml:"url"`
}

func isValidDomain(s string) bool {
	// Check if it's an IP address
	if net.ParseIP(s) != nil {
		return false // It's an IP, not a domain
	}

	// Attempt a DNS lookup to see if it resolves
	_, err := net.LookupHost(s)
	return err == nil
}

func ParseToSiteMap(links []LinkElement) {
	siteMap := SiteMap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  []UrlSiteMap{},
	}

	for _, link := range links {
		if !strings.Contains(link.Href, "calhoun") || !isValidDomain(link.Href) {
			fmt.Println(link.Href)
			siteMap.Urls = append(siteMap.Urls, UrlSiteMap{Loc: fmt.Sprintf("%s%s", "https://www.calhoun.io", link.Href)})
		} else {
			siteMap.Urls = append(siteMap.Urls, UrlSiteMap{Loc: link.Href})
		}
	}

	output, _ := xml.MarshalIndent(siteMap, "", "  ")

	fmt.Println(xml.Header + string(output))
}
