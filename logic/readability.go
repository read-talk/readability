package readability

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/custer-go/readability/util"
	nurl "net/url"
	"strings"
	"time"
)

func NewReadability(url string, timeout time.Duration) (*Document, error) {
	v := &Document{}
	var err error
	v.html, err = util.HttpGet(url, timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch readable content: %v", err)
	}

	v.url, _ = nurl.Parse(url)
	v.candidates = make(map[string]candidate, 0)
	v.html = util.ReplaceBrs.ReplaceAllString(v.html, "</p><p>")
	//v.html = util.ReplaceFonts.ReplaceAllString(v.html, `<\g<1>span>`)

	if v.html == "" {
		return nil, fmt.Errorf("html content is empty")
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(v.html))
	if err != nil {
		return nil, fmt.Errorf("failed return goquery document from io.Reader: %v", err)
	}
	v.htmlDoc = doc
	return v, nil
}
