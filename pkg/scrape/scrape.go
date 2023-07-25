package scrape

import (
	"fmt"
	"log"
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

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalln("Something went wrong getting response", err)
	})

	result := hukam{}

	// Find and print all links
	c.OnHTML("", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
		e.Request.Visit(e.Attr("href"))

		result.ang = 597
		result.gurmukhi = "ਮੇਰੈ ਹੀਅਰੈ ਰਤਨੁ ਨਾਮੁ ਹਰਿ ਬਸਿਆ ਗੁਰਿ ਹਾਥੁ ਧਰਿਓ ਮੇਰੈ ਮਾਥਾ ॥ ਜਨਮ ਜਨਮ ਕੇ ਕਿਲਬਿਖ ਦੁਖ ਉਤਰੇ ਗੁਰਿ ਨਾਮੁ ਦੀਓ ਰਿਨੁ ਲਾਥਾ ॥੧॥ ਮੇਰੇ ਮਨ ਭਜੁ ਰਾਮ ਨਾਮੁ ਸਭਿ ਅਰਥਾ ॥ ਗੁਰਿ ਪੂਰੈ ਹਰਿ ਨਾਮੁ ਦ੍ਰਿੜਾਇਆ ਬਿਨੁ ਨਾਵੈ ਜੀਵਨੁ ਬਿਰਥਾ ॥ ਰਹਾਉ ॥ ਬਿਨੁ ਗੁਰ ਮੂੜ ਭਏ ਹੈ ਮਨਮੁਖ ਤੇ ਮੋਹ ਮਾਇਆ ਨਿਤ ਫਾਥਾ ॥ ਤਿਨ ਸਾਧੂ ਚਰਣ ਨ ਸੇਵੇ ਕਬਹੂ ਤਿਨ ਸਭੁ ਜਨਮੁ ਅਕਾਥਾ ॥੨॥ ਜਿਨ ਸਾਧੂ ਚਰਣ ਸਾਧ ਪਗ ਸੇਵੇ ਤਿਨ ਸਫਲਿਓ ਜਨਮੁ ਸਨਾਥਾ ॥ ਮੋ ਕਉ ਕੀਜੈ ਦਾਸੁ ਦਾਸ ਦਾਸਨ ਕੋ ਹਰਿ ਦਇਆ ਧਾਰਿ ਜਗੰਨਾਥਾ ॥੩॥ ਹਮ ਅੰਧੁਲੇ ਗਿਆਨਹੀਨ ਅਗਿਆਨੀ ਕਿਉ ਚਾਲਹ ਮਾਰਗਿ ਪੰਥਾ ॥ ਹਮ ਅੰਧੁਲੇ ਕਉ ਗੁਰ ਅੰਚਲੁ ਦੀਜੈ ਜਨ ਨਾਨਕ ਚਲਹ ਮਿਲੰਥਾ ॥੪॥੧॥"
		result.dateTime = "2023-07-23 04:35:00"
	})

	c.Visit(url)

	c.Wait()

	return result

}
