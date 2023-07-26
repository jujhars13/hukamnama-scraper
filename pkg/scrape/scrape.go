package scrape

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// importing Colly
	"github.com/gocolly/colly"
)

type hukam struct {
	dateTime string
	ang      int
	gurmukhi string
}

func GetTodaysHukamnama(url string) hukam {

	c := colly.NewCollector(
	// colly.Async(true), colly.AllowedDomains("old.sgpc.net", "localhost"),
	)

	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	result := hukam{}
	result.dateTime = "2023-07-23 04:35:00"

	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > center:nth-child(4) > table > tbody > tr:nth-child(3) > td > p > font > strong > font:nth-child(2) > strong > font > font", func(e *colly.HTMLElement) {
		result.gurmukhi = strings.TrimSpace(e.Text)
	})
	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > div > table:nth-child(1) > tbody > tr > td:nth-child(2) > font", func(e *colly.HTMLElement) {
		// ang raw string looks like this `(AMg: 696)`, do a bunch of parsing and stripping
		rawAng := strings.TrimRight(strings.Split(strings.TrimSpace(e.Text), ": ")[1], ")")
		ang, err := strconv.Atoi(rawAng)
		if err != nil {
			log.Fatalln("Couldn't get ang number", err)
			result.ang = 0
		}
		result.ang = ang
	})

	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > center:nth-child(3) > table > tbody > tr:nth-child(3) > td > div > font > b > font > font > font", func(e *colly.HTMLElement) {
		fmt.Println("date is %i", e.Text)

		//result.dateTime = e.Text
	})

	c.Visit(url)

	// Display collector's statistics
	log.Println(c)

	//c.Wait()

	return result

}
