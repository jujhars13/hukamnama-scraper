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
	dateTime string
	gurmukhi string
	ang      int
}

// get todays hukamnama
// curl https://old.sgpc.net/hukumnama/indexhtml.asp --output "$(date '+%Y-%m-%d').html"
var testData = []testCase{
	{"2023-07-23.html",
		"2023-07-23 04:35:00",
		"ਮੇਰੈ ਹੀਅਰੈ ਰਤਨੁ ਨਾਮੁ ਹਰਿ ਬਸਿਆ ਗੁਰਿ ਹਾਥੁ ਧਰਿਓ ਮੇਰੈ ਮਾਥਾ ॥ ਜਨਮ ਜਨਮ ਕੇ ਕਿਲਬਿਖ ਦੁਖ ਉਤਰੇ ਗੁਰਿ ਨਾਮੁ ਦੀਓ ਰਿਨੁ ਲਾਥਾ ॥੧॥ ਮੇਰੇ ਮਨ ਭਜੁ ਰਾਮ ਨਾਮੁ ਸਭਿ ਅਰਥਾ ॥ ਗੁਰਿ ਪੂਰੈ ਹਰਿ ਨਾਮੁ ਦ੍ਰਿੜਾਇਆ ਬਿਨੁ ਨਾਵੈ ਜੀਵਨੁ ਬਿਰਥਾ ॥ ਰਹਾਉ ॥ ਬਿਨੁ ਗੁਰ ਮੂੜ ਭਏ ਹੈ ਮਨਮੁਖ ਤੇ ਮੋਹ ਮਾਇਆ ਨਿਤ ਫਾਥਾ ॥ ਤਿਨ ਸਾਧੂ ਚਰਣ ਨ ਸੇਵੇ ਕਬਹੂ ਤਿਨ ਸਭੁ ਜਨਮੁ ਅਕਾਥਾ ॥੨॥ ਜਿਨ ਸਾਧੂ ਚਰਣ ਸਾਧ ਪਗ ਸੇਵੇ ਤਿਨ ਸਫਲਿਓ ਜਨਮੁ ਸਨਾਥਾ ॥ ਮੋ ਕਉ ਕੀਜੈ ਦਾਸੁ ਦਾਸ ਦਾਸਨ ਕੋ ਹਰਿ ਦਇਆ ਧਾਰਿ ਜਗੰਨਾਥਾ ॥੩॥ ਹਮ ਅੰਧੁਲੇ ਗਿਆਨਹੀਨ ਅਗਿਆਨੀ ਕਿਉ ਚਾਲਹ ਮਾਰਗਿ ਪੰਥਾ ॥ ਹਮ ਅੰਧੁਲੇ ਕਉ ਗੁਰ ਅੰਚਲੁ ਦੀਜੈ ਜਨ ਨਾਨਕ ਚਲਹ ਮਿਲੰਥਾ ॥੪॥੧॥",
		696},
	// {"2023-07-25.html",
	// 	"2023-07-25 04:35:00",
	// 	"hir isau pRIiq AMqru mnu byiDAw hir ibnu rhxu n jweI ] ijau mCulI ibnu nIrY ibnsY iqau nwmY ibnu mir jweI ]1] myry pRB ikrpw jlu dyvhu hir nweI ] hau AMqir nwmu mMgw idnu rwqI nwmy hI sWiq pweI ] rhwau ] ijau cwiqRku jl ibnu ibllwvY ibnu jl ipAws n jweI ] gurmuiK jlu pwvY suK shjy hirAw Bwie suBweI ]2] mnmuK BUKy dh ids folih ibnu nwvY duKu pweI ] jnim mrY iPir jonI AwvY drgh imlY sjweI ]3] ikRpw krih qw hir gux gwvh hir rsu AMqir pweI ] nwnk dIn dieAwl Bey hY iqRsnw sbid buJweI ]4]8]",
	// 	607},
}

// Ensure that a webserver is running on 8081 with some test files
// this is more of an integration test than a unit test, but not gonna be strict about it as my skills are not there
// docker run -it -p 8081:80 -v ${PWD}:/usr/share/nginx/html nginx
func TestHukamnama(t *testing.T) {

	for _, wants := range testData {
		hukam := GetTodaysHukamnama(fmt.Sprintf("http://localhost:8081/%s", wants.filename))

		//fmt.Println(hukam)
		if hukam.dateTime != wants.dateTime {
			t.Fatalf("Hukamnama date is wrong, wants %s got %s,", hukam.dateTime, wants.dateTime)
		}
		if hukam.ang != wants.ang {
			t.Fatalf("Hukamnama ang is wrong, wants %d got %d,", hukam.ang, wants.ang)
		}
		if hukam.gurmukhi != wants.gurmukhi {
			t.Fatalf("Hukamnama Gurmukhi is wrong, wants %s got %s,", hukam.gurmukhi, wants.gurmukhi)
		}
	}
}
