package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io"
	"net/url"
	"os"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
	exitOK   = 0
)

func Find(slice []string, val string) (string, error) {
	for _, item := range slice {
		if item == val {
			return val, nil
		}
	}
	return "", errors.New("Unable to find string")
}

func main() {

	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {

	CliIntro()

	siteCallSignType, siteLetterType, siteRegion, filename, err := CommandSwitches()
	if err != nil {
		return err
	}

	visited := []string{}

	params := url.Values{}
	if siteCallSignType != "" {
		params.Add("filterCallsignTypeId_Saved", siteCallSignType)
	}
	if siteLetterType != "" {
		params.Add("filterLetterTypeCode_Saved", siteLetterType)
	}
	if siteRegion != "" {
		params.Add("filterRegionId_Saved", siteRegion)
	}

	params.Add("filterCallsignSuffix_Operator01_Saved", "And")
	params.Add("filterCallsignSuffix_Operator02_Saved", "And")
	params.Add("sortDirection", "Ascending")

	file, err := os.Create(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		fmt.Println("Unable to create file")
		os.Exit(exitFail)
	}
	defer file.Close()
	fileWriter := bufio.NewWriter(file)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("csdb.utas.edu.au"),
	)

	c.OnHTML(".amc-main", func(e *colly.HTMLElement) {
		e.ForEach("table tr", func(_ int, el *colly.HTMLElement) {
			callData := el.ChildTexts("td")
			if len(callData) > 0 {
				fmt.Fprintln(fileWriter, callData[2])
			}
		})
		e.ForEach(".pagination-container .pagination li a[href]", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")

			urlObj, _ := url.Parse(link)
			query := urlObj.RawQuery
			vals, _ := url.ParseQuery(query)
			page := vals.Get("page")

			_, err := Find(visited, page)

			// Hasn't been visited
			if err != nil {
				// Visit link found on page
				// Only those links are visited which are in AllowedDomains
				c.Visit(el.Request.AbsoluteURL(link))
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

		query := r.URL.RawQuery
		vals, _ := url.ParseQuery(query)
		page := vals.Get("page")
		if page != "" {
			visited = append(visited, page)
		} else {
			visited = append(visited, "1")
		}

	})

	params.Set("page", "1")

	c.Visit("https://csdb.utas.edu.au/Callsign/SearchUnallocated?" + params.Encode())

	fileWriter.Flush()
	return nil
}
