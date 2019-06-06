package parser

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func ParseTitl(contents []byte, region string) engine.ParseResult {
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


	return result
}