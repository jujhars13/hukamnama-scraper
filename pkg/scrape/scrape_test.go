package scrape

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func setupWebServer() {
	fmt.Printf("Starting server at port 8081\n")
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type testCase struct {
	filename string
	date     string
	gurmukhi string
	ang      int
}

var testData = []testCase{
	{"example1.html",
		"[July 21, 2023, Friday 04:45 AM. IST]",
		"ਹੇ ਪ੍ਰਭੂ ਜੀ! ਤੂੰ ਸਾਨੂੰ ਸਭ ਪਦਾਰਥ ਦੇਣ ਵਾਲਾ ਹੈਂ, ਦਾਤਾਂ ਦੇਣ ਵਿਚ ਤੂੰ ਕਦੇ ਖੁੰਝਦਾ ਨਹੀਂ, ਅਸੀ ਤੇਰੇ (ਦਰ ਦੇ) ਮੰਗਤੇ ਹਾਂ । ਮੈਂ ਤੈਥੋਂ ਕੇਹੜੀ ਸ਼ੈ ਮੰਗਾਂ? ਕੋਈ ਸ਼ੈ ਸਦਾ ਟਿਕੀ ਰਹਿਣ ਵਾਲੀ ਨਹੀਂ । (ਹਾਂ, ਤੇਰਾ ਨਾਮ ਸਦਾ-ਥਿਰ ਰਹਿਣ ਵਾਲਾ ਹੈ) ਹੇ ਹਰੀ! ਮੈਨੂੰ ਆਪਣਾ ਨਾਮ ਦੇਹ, ਮੈਂ ਤੇਰੇ ਨਾਮ ਨੂੰ ਪਿਆਰ ਕਰਾਂ ।੧।ਪਰਮਾਤਮਾ ਹਰੇਕ ਸਰੀਰ ਵਿਚ ਵਿਆਪਕ ਹੈ । ਪਾਣੀ ਵਿਚ ਧਰਤੀ ਵਿਚ, ਧਰਤੀ ਦੇ ਉਪਰ ਆਕਾਸ਼ ਵਿਚ ਹਰ ਥਾਂ ਮੌਜੂਦ ਹੈ ਪਰ ਲੁਕਿਆ ਹੋਇਆ ਹੈ । (ਹੇ ਮਨ!) ਗੁਰੂ ਦੇ ਸ਼ਬਦ ਦੀ ਰਾਹੀਂ ਉਸ ਨੂੰ ਵੇਖ ।ਰਹਾਉ।(ਹੇ ਭਾਈ! ਜਿਸ ਮਨੁੱਖ ਉੱਤੇ) ਗੁਰੂ ਨੇ ਸਤਿਗੁਰੂ ਨੇ ਕਿਰਪਾ ਕੀਤੀ ਉਸ ਨੂੰ ਉਸ ਨੇ ਧਰਤੀ ਆਕਾਸ਼ ਪਾਤਾਲ (ਸਾਰਾ ਜਗਤ ਪਰਮਾਤਮਾ ਦੀ ਹੋਂਦ ਨਾਲ ਭਰਪੂਰ) ਵਿਖਾ ਦਿੱਤਾ । ਉਹ ਪਰਮਾਤਮਾ ਜੂਨਾਂ ਵਿਚ ਨਹੀਂ ਆਉਂਦਾ, ਹੁਣ ਭੀ ਮੌਜੂਦ ਹੈ ਅਗਾਂਹ ਨੂੰ ਮੌਜੂਦ ਰਹੇਗਾ, (ਹੇ ਭਾਈ!) ਉਸ ਪ੍ਰਭੂ ਨੂੰ ਤੂੰ ਆਪਣੇ ਹਿਰਦੇ ਵਿਚ (ਵੱਸਦਾ) ਵੇਖ ।੨।ਇਹ ਭਾਗ-ਹੀਣ ਜਗਤ ਜਨਮ ਮਰਨ ਦਾ ਗੇੜ ਸਹੇੜ ਬੈਠਾ ਹੈ ਕਿਉਂਕਿ ਇਸ ਨੇ ਮਾਇਆ ਦੇ ਮੋਹ ਵਿਚ ਪੈ ਕੇ ਪਰਮਾਤਮਾ ਦੀ ਭਗਤੀ ਭੁਲਾ ਦਿੱਤੀ ਹੈ । ਜੇ ਸਤਿਗੁਰੂ ਮਿਲ ਪਏ ਤਾਂ ਗੁਰੂ ਦੇ ਉਪਦੇਸ਼ ਤੇ ਤੁਰਿਆਂ (ਪ੍ਰਭੂ ਦੀ ਭਗਤੀ) ਪ੍ਰਾਪਤ ਹੁੰਦੀ ਹੈ, ਪਰ ਮਾਇਆ-ਵੇੜ੍ਹੇ ਜੀਵ (ਭਗਤੀ ਤੋਂ ਖੁੰਝ ਕੇ ਮਨੁੱਖਾ ਜਨਮ ਦੀ) ਬਾਜ਼ੀ ਹਾਰ ਜਾਂਦੇ ਹਨ ।੩।ਹੇ ਸਤਿਗੁਰੂ! ਮਾਇਆ ਦੇ ਬੰਧਨ ਤੋੜ ਕੇ ਜਿਨ੍ਹਾਂ ਬੰਦਿਆਂ ਨੂੰ ਤੂੰ ਮਾਇਆ ਤੋਂ ਨਿਰਲੇਪ ਕਰ ਦੇਂਦਾ ਹੈਂ, ਉਹ ਮੁੜ ਜਨਮ ਮਰਨ ਦੇ ਗੇੜ ਵਿਚ ਨਹੀਂ ਪੈਂਦਾ । ਹੇ ਨਾਨਕ! (ਗੁਰੂ ਦੀ ਕਿਰਪਾ ਨਾਲ ਜਿਨ੍ਹਾਂ ਦੇ ਅੰਦਰ ਪਰਮਾਤਮਾ ਦੇ) ਗਿਆਨ ਦਾ ਰਤਨ ਚਮਕ ਪੈਂਦਾ ਹੈ, ਉਹਨਾਂ ਦੇ ਮਨ ਵਿਚ ਹਰੀ ਨਿਰੰਕਾਰ (ਆਪ) ਆ ਵੱਸਦਾ ਹੈ ।੪।੮।"
		597},
}

// Ensure that a webserver is running on 8081 with some test files
// this is more of an integration test than a unit test
// docker run -it -p 8081:80 -v ${PWD}:/usr/share/nginx/html nginx
func TestHukamnama(t *testing.T) {

	hukam := GetTodaysHukamnama("http://localhost:8081/example1.html")

	if hukam.date != "2023-07-21" {
		t.Fatalf("Hukamnama date is wrong, got %s,", hukam.date)
	}
	if hukam.gurmukhi != "Satnaam" {
		t.Fatalf("Hukamnama gurmukhi is wrong, got %s,", hukam.gurmukhi)
	}

	fmt.Println(hukam)
}
