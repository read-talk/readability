package main

import (
	"flag"
	"fmt"
	readability "github.com/custer-go/readability/logic"
	"time"
)

var url string

func init() {
	flag.StringVar(&url, "url", "http://wd.leiting.com/home/news/news_detail.php?id=599", "input url to parse")
}

func main() {
	flag.Parse()
	fmt.Println("url: ", url)
	test, err := readability.NewReadability(url, 30*time.Second)
	if err!=nil {
		fmt.Printf("failed test: %v\n", err)
	}
	test.Parse()
	fmt.Println(test.Title)
}
