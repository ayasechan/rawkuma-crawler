package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	urlNeedCrawl = flag.String("url", "", "下载链接")
)

func main() {
	flag.Parse()
	buf, err := HTTPGet(*urlNeedCrawl)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".dload").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})

}

func HTTPGet(rawURL string) ([]byte, error) {
	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"
	req, _ := http.NewRequest("GET", rawURL, nil)
	req.Header.Set("user-agent", ua)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	return io.ReadAll(resp.Body)
}
