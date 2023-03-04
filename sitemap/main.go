package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/EdwinVesga/gophercises/linkparser/parser"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url tha you want to build a site map for.")
	maxDepth := flag.Int("depth", 3, "The maximum number of links deep to traverse.")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	toXml := urlset{
		Urls:  make([]loc, len(pages)),
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	for i, page := range pages {
		toXml.Urls[i] = loc{page}
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "   ")
	if err := enc.Encode(toXml); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

type empty struct{}

func bfs(urlString string, maxDepth int) []string {
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		urlString: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]empty)
		for k, _ := range q {
			if _, ok := seen[k]; ok {
				continue
			}
			seen[k] = empty{}
			for _, link := range get(k) {
				nq[link] = empty{}
			}
		}
	}

	ret := make([]string, 0, len(seen))

	for k := range seen {
		ret = append(ret, k)
	}

	return ret
}

func get(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := parser.Parse(r)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(prefix string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}
