package parser

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
)

func ParseProfile(contents []byte) engine.ParseResult {
	doc, err := htmlquery.Parse(strings.NewReader(string(contents)))
	if err != nil {
		log.Printf("htmlquery parse err :", err)
	}
	title := htmlquery.FindOne(doc, "//title")
	name := htmlquery.FindOne(doc, "//div/h1")

	profile := model.Profile{}
	profile.Title = htmlquery.InnerText(title)
	profile.Name = htmlquery.InnerText(name)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}
