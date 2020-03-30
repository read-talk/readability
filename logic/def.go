package readability

import (
	"github.com/PuerkitoBio/goquery"
	nurl "net/url"
)

type candidate struct {
	score float64
	node  *goquery.Selection
}

type Document struct {
	html       string
	url        *nurl.URL
	htmlDoc    *goquery.Document
	candidates map[string]candidate

	Title   string
	Content string
}
