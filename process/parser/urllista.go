package parser

import (
	"awesomeProject/engine"
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"net/url"
	"strings"
)

func ParseCityList(contents []byte, region string) engine.ParseResult {
	doc, err := htmlquery.Parse(strings.NewReader(string(contents)))
	if err != nil{
		log.Printf("htmlquery parse err :", err)
	}
	list := htmlquery.Find(doc, "//a")
	ti := htmlquery.InnerText(htmlquery.FindOne(doc, "//title"))
	fmt.Print(ti)
	result := engine.ParseResult{}
	r := 0
	for _, n := range list {
		uri := string(htmlquery.SelectAttr(n, "href"))
		u, err := url.Parse(uri)
		if err != nil {
			continue
		}

		l := strings.Join(strings.Split(u.Host, ".")[1:],".")
		if  l == region && (u.Scheme == "http" || u.Scheme == "https")   {
			result.Items = ti
			result.Requests = append(result.Requests, engine.Request{
				Url: uri,
				Region: region,
				ParserFunc: ParseTi,
			})
			r++
			fmt.Printf("n: \"%s\",\n", r, uri)
		}
	}
	return result
}