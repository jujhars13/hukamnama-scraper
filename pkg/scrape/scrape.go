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
		colly.Async(true), colly.AllowedDomains("old.sgpc.net", "localhost"),
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

	// hukamnama Gurmukhi could appear in either JSPath selector
	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > center:nth-child(4) > table > tbody > tr:nth-child(2) > td > p > font > strong > font:nth-child(2) > strong > font > font", func(e *colly.HTMLElement) {
		gurmukhi := strings.TrimSpace(e.Text)
		result.gurmukhi = gurmukhi
	})
	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > center:nth-child(4) > table > tbody > tr:nth-child(3) > td > p > font > strong > font:nth-child(2) > strong > font > font", func(e *colly.HTMLElement) {
		gurmukhi := strings.TrimSpace(e.Text)

		result.gurmukhi = gurmukhi
	})

	// ang raw string looks like this `(AMg: 696)` and sometimes like `(AMg 696)` do a bunch of parsing and stripping
	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > div > table:nth-child(1) > tbody > tr > td:nth-child(2) > font", func(e *colly.HTMLElement) {
		rawAng := strings.Trim(strings.Split(strings.TrimSpace(e.Text), " ")[1], ")")
		ang, err := strconv.Atoi(rawAng)
		if err != nil {
			log.Fatalln("Couldn't get ang number", err)
			result.ang = 0
		}
		result.ang = ang
	})

	c.OnHTML("body > table:nth-child(3) > tbody > tr:nth-child(1) > td:nth-child(6) > blockquote > center:nth-child(3) > table > tbody > tr:nth-child(3) > td > div > font > b > font > font > font", func(e *colly.HTMLElement) {
		dateString := strings.TrimSpace(e.Text)
		parseFormat := "[January 2, 2006, Monday 03:04 AM. IST]"
		dateTime, err := time.Parse(parseFormat, dateString)
		if err != nil {
			log.Fatalln("Couldn't parse date", err)
			result.dateTime = ""
		}
		outputFormat := "2006-01-02 15:04"
		result.dateTime = dateTime.Format(outputFormat)
	})

	c.Visit(url)

	// Display collector's statistics
	log.Println(c)

	if result.gurmukhi == "" {
		log.Fatalln("Couldn't get Gurmukhi using tr:nth-child(3)", result.gurmukhi)
	}

	//c.Wait()

	return result

}
