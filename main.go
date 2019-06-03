package main

import (
	"awesomeProject/engine"
	"awesomeProject/zhenai/parser"
)

//func ParseProfile(contents []byte)  {
//	doc, err := htmlquery.Parse(strings.NewReader(string(contents)))
//	if err != nil {
//		log.Printf("htmlquery parse err :", err)
//	}
//	title := htmlquery.FindOne(doc, "//title")
//	name := htmlquery.FindOne(doc, "//div/h1")
//
//	log.Printf("%s", htmlquery.InnerText(title))
//	log.Printf("%s", htmlquery.InnerText(name))


func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}