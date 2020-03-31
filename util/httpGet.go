package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	nurl "net/url"
	"strings"
	"time"
)

// httpGet fetch the web page from specified url, check if it's
// readable, then parses the response to find the readable content.
func HttpGet(url string, timeout time.Duration) (string, error) {
	// 1. Make sure URL is valid
	_, err := nurl.ParseRequestURI(url)
	if err != nil {
		return "", fmt.Errorf("1.failed to parse URL: %v", err)
	}
	// 2. Fetch page from URL
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("2.failed to fetch the page: %v", err)
	}
	defer resp.Body.Close()

	// 3. Make sure content type is HTML
	cp := resp.Header.Get("Content-Type")
	if !strings.Contains(cp, "text/html") {
		return "", fmt.Errorf("3.URL is not a HTML document")
	}
	// 4. Check if the page is readable
	var buffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &buffer)

	if _, err = ioutil.ReadAll(tee); err != nil {
		return "", fmt.Errorf("4.the page is not readable")
	}

	b, err := ioutil.ReadAll(&buffer)
	if err != nil {
		return "", fmt.Errorf("4.the page is not readable")
	}

	// 5. go to Parse content
	return string(b), nil
}
