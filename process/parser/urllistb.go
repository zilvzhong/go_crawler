package parser

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strings"
)


func ParseTi(contents []byte, region string) engine.ParseResult {
	body := strings.NewReader(string(contents))
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("xpath title parse err :", err)
	}
	profile := model.Profile{}
	profile.Title = strings.TrimSpace(doc.Find("title").Text())
	result := engine.ParseResult{
		Items: profile.Title,
	}

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		for i := 0; i < selection.Length(); i++ {
			uri, _ := selection.Eq(i).Attr("href")
			u, err := url.Parse(uri)
			if err != nil {
				continue
			}
			l := strings.Join(strings.Split(u.Host, ".")[1:],".")
			if  l == region && (u.Scheme == "http" || u.Scheme == "https")  {
				result.Requests = append(result.Requests, engine.Request{
					Url:	uri,
					Region: region,
					ParserFunc:		ParseTit,
				})
				fmt.Printf("2: \"%s\",\n", uri)
			}
		}
	})
	return result
}


