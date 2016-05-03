package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "autosuggest"
	app.Usage = "Autosuggest results from YOUTUBE"
	app.Action = func(c *cli.Context) {

		keywords := strings.Join(c.Args(), "%20")

		url := fmt.Sprintf("https://clients1.google.com/complete/search?client=youtube&hl=en&gl=gb&gs_rn=23&gs_ri=youtube&ds=yt&cp=5&gs_id=9&q=%s&callback=google.sbox.p50&gs_gbg=DCK2xxW6vto2FGXTX8nec69nB8lkn47dM", keywords)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Set("pragma", "no-cache")
		req.Header.Set("accept-language", "en-US,en;q=0.8")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Safari/537.36")
		req.Header.Set("accept", "*/*")
		req.Header.Set("cache-control", "no-cache")
		req.Header.Set("authority", "clients1.google.com")
		req.Header.Set("cookie", "NID=79=GgTh92amgEklcFvTrLEw8rULr0EbEmLwntC6pKRSkRj31-Q3iW-C15huW44cEAIavNVnD2el0woa_ygOkY7Ws7PD66byCBWSSbgPczXRDQTQDsy1d6wKBEAjjCTZlUbM")
		req.Header.Set("referer", "https://www.youtube.com/?gl=GB")

		response, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		bodyString := string(body)

		// The json contains arrays of mixed type :(  Will be simpler to use regex
		r, _ := regexp.Compile(`\["(.*?)"`)
		results := r.FindAllStringSubmatch(bodyString, -1)

		for _, result := range results {
			// subMatches := r2.FindStringSubmatch(item)
			fmt.Println("--")

			if len(result) > 0 {
				fmt.Println(result[1])
			}

		}
	}
	app.Run(os.Args)
}
