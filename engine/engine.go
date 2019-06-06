package engine

import (
	"awesomeProject/fetcher"
	"fmt"
	"log"
	"net/url"
	"strings"
)



func Run(seeds ...Request) {
	var requests []Request
	var ti_data []map[string]string
	var urlerr string
	urisetlist := New()

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		do := r.Region
		body, err := fetcher.FetchUseHeader(r.Url)

		if err != nil {
			log.Printf("网址Body: error" + "fetching url %s: %v \n", r.Url, err)
			e := fmt.Sprintf("Error", err)
			urlerr = strings.TrimLeft(strings.Split(e, "http")[1], "s://")
			bundle := make(map[string]string)
			bundle["Url"] = r.Url
			bundle["StatusCode"] = urlerr
			ti_data =append(ti_data, bundle)
			continue
		}else {
			urlerr = fmt.Sprintf("200")
		}
		parseResult := r.ParserFunc(body, do)

		for _, p := range parseResult.Requests {
			uri := strings.TrimRight(p.Url, "/")
			u, _ := url.Parse(uri)
			if u.RawQuery == "" && u.Fragment == "" && !strings.ContainsAny(u.Path, "123456789_#?") &&
				(u.Path == "" || strings.Contains(u.Path, "/index")){
				if urisetlist.Has(uri) {
					continue
				} else {
					urisetlist.Add(uri)
					requests = append(requests, p)
				}
			}

		}

		bundle := make(map[string]string)
		bundle["Url"] = r.Url
		bundle["Title"] = parseResult.Items
		bundle["StatusCode"] = urlerr
		ti_data =append(ti_data, bundle)

	}
	for k, v := range ti_data {
		fmt.Print (k,v)
		fmt.Print("\n")
	}

}