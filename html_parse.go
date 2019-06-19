package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

type HtmlParse struct {
	hotUrl  []string
	hotWord []map[string]string
}

func (c *HtmlParse) GetHotUrl(uftbody *iconv.Reader) {
	doc, err := goquery.NewDocumentFromReader(uftbody)
	if err != nil {
		log.Fatalln("NewDocumentFromReader", err)
	}

	doc.Find(".hblock ul li").Each(func(i int, s *goquery.Selection) {
		if value, isExist := s.Find("a").Attr("href"); isExist {
			c.hotUrl = append(c.hotUrl, value)
		}
	})
}

func (c *HtmlParse) GetDetails(uftbody *iconv.Reader) {
	doc, err := goquery.NewDocumentFromReader(uftbody)
	if err != nil {
		log.Fatalln("NewDocumentFromReader", err)
	}

	doc.Find("tbody  tr").Filter("tr[class!=item-tr]").Each(func(i int, s *goquery.Selection) {
		numTop := s.Find("td[class=first]>span").Text()
		keyword := s.Find("td[class=keyword]>a[class=list-title]").Text()
		hot := s.Find("td[class=last]>span").Text()
		if i != 0 && numTop != "" {
			c.hotWord = append(c.hotWord, map[string]string{
				"numTop":  numTop,
				"keyword": keyword,
				"hot":     hot,
			})
		}
	})

}
