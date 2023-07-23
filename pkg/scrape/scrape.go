package scrape

import (
	"fmt"
	"log"

	// importing Colly
	"github.com/gocolly/colly"
)

type hukam struct {
	dateTime string
	ang      int
	gurmukhi string
}

func GetTodaysHukamnama(url string) hukam {
	result := hukam{"2023-07-23 04:35:00",
		597,
		"ਮੇਰੈ ਹੀਅਰੈ ਰਤਨੁ ਨਾਮੁ ਹਰਿ ਬਸਿਆ ਗੁਰਿ ਹਾਥੁ ਧਰਿਓ ਮੇਰੈ ਮਾਥਾ ॥ ਜਨਮ ਜਨਮ ਕੇ ਕਿਲਬਿਖ ਦੁਖ ਉਤਰੇ ਗੁਰਿ ਨਾਮੁ ਦੀਓ ਰਿਨੁ ਲਾਥਾ ॥੧॥ ਮੇਰੇ ਮਨ ਭਜੁ ਰਾਮ ਨਾਮੁ ਸਭਿ ਅਰਥਾ ॥ ਗੁਰਿ ਪੂਰੈ ਹਰਿ ਨਾਮੁ ਦ੍ਰਿੜਾਇਆ ਬਿਨੁ ਨਾਵੈ ਜੀਵਨੁ ਬਿਰਥਾ ॥ ਰਹਾਉ ॥ ਬਿਨੁ ਗੁਰ ਮੂੜ ਭਏ ਹੈ ਮਨਮੁਖ ਤੇ ਮੋਹ ਮਾਇਆ ਨਿਤ ਫਾਥਾ ॥ ਤਿਨ ਸਾਧੂ ਚਰਣ ਨ ਸੇਵੇ ਕਬਹੂ ਤਿਨ ਸਭੁ ਜਨਮੁ ਅਕਾਥਾ ॥੨॥ ਜਿਨ ਸਾਧੂ ਚਰਣ ਸਾਧ ਪਗ ਸੇਵੇ ਤਿਨ ਸਫਲਿਓ ਜਨਮੁ ਸਨਾਥਾ ॥ ਮੋ ਕਉ ਕੀਜੈ ਦਾਸੁ ਦਾਸ ਦਾਸਨ ਕੋ ਹਰਿ ਦਇਆ ਧਾਰਿ ਜਗੰਨਾਥਾ ॥੩॥ ਹਮ ਅੰਧੁਲੇ ਗਿਆਨਹੀਨ ਅਗਿਆਨੀ ਕਿਉ ਚਾਲਹ ਮਾਰਗਿ ਪੰਥਾ ॥ ਹਮ ਅੰਧੁਲੇ ਕਉ ਗੁਰ ਅੰਚਲੁ ਦੀਜੈ ਜਨ ਨਾਨਕ ਚਲਹ ਮਿਲੰਥਾ ॥੪॥੧॥"}

	c := colly.NewCollector(
		colly.AllowedDomains("old.sgpc.net", "localhost"),
	)
	// Find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong getting response:", err)
	})

	c.Visit(url)

	return result
}
