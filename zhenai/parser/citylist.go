package parser

import (
	"awesomeProject/engine"
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
)

func ParseCityList(contents []byte) engine.ParseResult {
	doc, err := htmlquery.Parse(strings.NewReader(string(contents)))
	if err != nil{
		log.Printf("htmlquery parse err :", err)
	}
	list := htmlquery.Find(doc, "//a")
	result := engine.ParseResult{}

	//limit := 10
	for _, n := range list {
		result.Items = append(result.Items, "City" + htmlquery.OutputHTML(n, false))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(htmlquery.SelectAttr(n, "href")),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}