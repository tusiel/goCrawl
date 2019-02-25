package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"./manager"
	"./utils"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string
var timeout uint64
var linkManager *manager.LinkManager

func init() {
	startPage := flag.String("startPage", "", "The starting URL that will be crawled")
	reqTimeout := flag.Uint64("timeout", 5, "Timeout for each GET request")

	flag.Parse()

	if *startPage == "" {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	timeout = *reqTimeout

	baseURL = *startPage
	linkManager = manager.NewLinkManager()
}

func main() {
	linkManager.Add(1)
	go crawlPage(baseURL)

	linkManager.Wait()

	for url, suburl := range linkManager.GetReport() {
		fmt.Printf("\nURI: %s\n", url)
		for su, count := range suburl {
			fmt.Printf("\t\t\t\tLinked URL: %s - %d\n", su, count)
		}
	}

}

func crawlPage(uri string) {
	defer linkManager.Done()

	fmt.Print(".") // Give user feedback that something is happening

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	parsed, err := url.Parse(uri)
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}

	response, err := client.Get(parsed.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	toProcess := []string{}

	document.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if !exists {
			return
		}

		if utils.IsExternalDomain(link) {
			linkManager.SetLinkProcessed(link)
			linkManager.SetReportCount(uri, link)
			return
		}

		relativeURL := utils.GetRelativeURL(link, baseURL)

		if !linkManager.IsProcessed(relativeURL) {
			linkManager.SetLinkProcessed(relativeURL)

			toProcess = append(toProcess, relativeURL)

			linkManager.SetReportCount(uri, relativeURL)
		} else {
			linkManager.SetReportCount(uri, relativeURL)
		}
	})

	for _, l := range toProcess {
		linkManager.Add(1)
		go crawlPage(l)
	}
}
