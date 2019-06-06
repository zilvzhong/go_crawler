package main

import (
	"awesomeProject/engine"
	"awesomeProject/process/parser"
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
		Url: "http://www.buaa.edu.cn",
		Region: "buaa.edu.cn",
		ParserFunc: parser.ParseCityList,
	},
	)

	//s := []string{"状态码:%!(EXTRA *url.Error=Get http://oa.buaa.edu.cn: dial tcp: lookup oa.buaa.edu.cn: no such host)",
	////	"状态码:%!(EXTRA *errors.errorString=dial 404"}
	//fmt.Printf("%q\n", strings.Split("状态码:%!(EXTRA *errors.errorString=dial 404", "dial "))


}


//for i := 0; i < selection.Length(); i++ {
//d, _ := selection.Eq(i).Attr("href")
//u, err := url.Parse(d)
//if err != nil {
//continue
//}
//l := strings.Join(strings.Split(u.Host, ".")[1:],".")
//if  l == region && (u.Scheme == "http" || u.Scheme == "https")  {
//uri := u.Scheme + "://" + u.Host