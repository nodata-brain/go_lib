package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("").Each(func(i int, s *goquery.Selection) {
		text := s.Find("").Text()
		fmt.Printf(text)

	})
}

func main() {
	url := ""
	GetPage(url)
}
