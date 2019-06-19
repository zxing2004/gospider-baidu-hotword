package main

import (
	"log"
	"net/http"

	iconv "github.com/djimenez/iconv-go"
)

type HtmlDownload struct{}

func (c *HtmlDownload) Downloader(url string) *iconv.Reader {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest", err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	client := &http.Client{}
	resq, err := client.Do(req)
	if err != nil {
		log.Fatalln("client.Do", err)
	}
	if resq.StatusCode != 200 {
		return nil
	}
	defer resq.Body.Close()
	uftbody, err := iconv.NewReader(resq.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatalln("NewReader", err)
	}

	return uftbody
}
