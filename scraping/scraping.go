package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetPage(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var text []string
	html := doc.Find("* > section#contentsArea > div.scoreBox")
	html.Each(func(index int, s *goquery.Selection) {

		fmt.Println(index)
		fmt.Println(s)
	})

	fmt.Println(html)
	text = []string{"tes"}

	return text
}

//type baseball
func baseball() []string {

	t := time.Now()
	fmt.Println(t)
	url := "https://www.nikkansports.com/baseball/professional/score/2019/pf-score-20190626.html"
	result := GetPage(url)

	return result

}

func main() {

	//read commandline option
	media := flag.String("m", "baseball", "media_type")
	flag.Parse()

	var text []string
	//select media type
	if *media == "baseball" {
		text = baseball()
	}

	//output
	fmt.Println(text)

}
