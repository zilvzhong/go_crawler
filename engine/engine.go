package engine

import (
	"awesomeProject/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.FetchUseHeader(r.Url)
		if err != nil {
			log.Printf("Fetcher: error" + "fetching url %s: %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}



//func (SimpleEngine) worker(r Request) (ParseResult, error) {
//	log.Printf("Fetching %s", r.Url)
//	body, err := fetcher.Fetch(r.Url)
//	if err != nil {
//		log.Printf("Fetcher: Error" + "fetching url %s: %v", r.Url, err)
//		return ParseResult{}, err
//	}
//
//	return r.ParserFunc(body, nil)
//}