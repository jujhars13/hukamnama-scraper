package scrape

import (
	"fmt"
	"testing"
)

type testCase struct {
	filename string
	dateTime string
	gurmukhi string
	ang      int
}

// get todays hukamnama
// curl https://old.sgpc.net/hukumnama/indexhtml.asp --output "$(date '+%Y-%m-%d').html"
var testData = []testCase{
	{"2023-07-21.html",
		"2023-07-21 04:45",
		"ਤੂ ਪ੍ਰਭ ਦਾਤਾ ਦਾਨਿ ਮਤਿ ਪੂਰਾ ਹਮ ਥਾਰੇ ਭੇਖਾਰੀ ਜੀਉ ॥ ਮੈ ਕਿਆ ਮਾਗਉ ਕਿਛੁ ਥਿਰੁ ਨ ਰਹਾਈ ਹਰਿ ਦੀਜੈ ਨਾਮੁ ਪਿਆਰੀ ਜੀਉ ॥੧॥ ਘਟਿ ਘਟਿ ਰਵਿ ਰਹਿਆ ਬਨਵਾਰੀ ॥ ਜਲਿ ਥਲਿ ਮਹੀਅਲਿ ਗੁਪਤੋ ਵਰਤੈ ਗੁਰ ਸਬਦੀ ਦੇਖਿ ਨਿਹਾਰੀ ਜੀਉ ॥ ਰਹਾਉ ॥ ਮਰਤ ਪਇਆਲ ਅਕਾਸੁ ਦਿਖਾਇਓ ਗੁਰਿ ਸਤਿਗੁਰਿ ਕਿਰਪਾ ਧਾਰੀ ਜੀਉ ॥ ਸੋ ਬ੍ਰਹਮੁ ਅਜੋਨੀ ਹੈ ਭੀ ਹੋਨੀ ਘਟ ਭੀਤਰਿ ਦੇਖੁ ਮੁਰਾਰੀ ਜੀਉ ॥੨॥ ਜਨਮ ਮਰਨ ਕਉ ਇਹੁ ਜਗੁ ਬਪੁੜੋ ਇਨਿ ਦੂਜੈ ਭਗਤਿ ਵਿਸਾਰੀ ਜੀਉ ॥ ਸਤਿਗੁਰੁ ਮਿਲੈ ਤ ਗੁਰਮਤਿ ਪਾਈਐ ਸਾਕਤ ਬਾਜੀ ਹਾਰੀ ਜੀਉ ॥੩॥ ਸਤਿਗੁਰ ਬੰਧਨ ਤੋੜਿ ਨਿਰਾਰੇ ਬਹੁੜਿ ਨ ਗਰਭ ਮਝਾਰੀ ਜੀਉ ॥ ਨਾਨਕ ਗਿਆਨ ਰਤਨੁ ਪਰਗਾਸਿਆ ਹਰਿ ਮਨਿ ਵਸਿਆ ਨਿਰੰਕਾਰੀ ਜੀਉ ॥੪॥੮॥",
		597},
	{"2023-07-23.html",
		"2023-07-23 04:45",
		"ਮੇਰੈ ਹੀਅਰੈ ਰਤਨੁ ਨਾਮੁ ਹਰਿ ਬਸਿਆ ਗੁਰਿ ਹਾਥੁ ਧਰਿਓ ਮੇਰੈ ਮਾਥਾ ॥ ਜਨਮ ਜਨਮ ਕੇ ਕਿਲਬਿਖ ਦੁਖ ਉਤਰੇ ਗੁਰਿ ਨਾਮੁ ਦੀਓ ਰਿਨੁ ਲਾਥਾ ॥੧॥ ਮੇਰੇ ਮਨ ਭਜੁ ਰਾਮ ਨਾਮੁ ਸਭਿ ਅਰਥਾ ॥ ਗੁਰਿ ਪੂਰੈ ਹਰਿ ਨਾਮੁ ਦ੍ਰਿੜਾਇਆ ਬਿਨੁ ਨਾਵੈ ਜੀਵਨੁ ਬਿਰਥਾ ॥ ਰਹਾਉ ॥ ਬਿਨੁ ਗੁਰ ਮੂੜ ਭਏ ਹੈ ਮਨਮੁਖ ਤੇ ਮੋਹ ਮਾਇਆ ਨਿਤ ਫਾਥਾ ॥ ਤਿਨ ਸਾਧੂ ਚਰਣ ਨ ਸੇਵੇ ਕਬਹੂ ਤਿਨ ਸਭੁ ਜਨਮੁ ਅਕਾਥਾ ॥੨॥ ਜਿਨ ਸਾਧੂ ਚਰਣ ਸਾਧ ਪਗ ਸੇਵੇ ਤਿਨ ਸਫਲਿਓ ਜਨਮੁ ਸਨਾਥਾ ॥ ਮੋ ਕਉ ਕੀਜੈ ਦਾਸੁ ਦਾਸ ਦਾਸਨ ਕੋ ਹਰਿ ਦਇਆ ਧਾਰਿ ਜਗੰਨਾਥਾ ॥੩॥ ਹਮ ਅੰਧੁਲੇ ਗਿਆਨਹੀਨ ਅਗਿਆਨੀ ਕਿਉ ਚਾਲਹ ਮਾਰਗਿ ਪੰਥਾ ॥ ਹਮ ਅੰਧੁਲੇ ਕਉ ਗੁਰ ਅੰਚਲੁ ਦੀਜੈ ਜਨ ਨਾਨਕ ਚਲਹ ਮਿਲੰਥਾ ॥੪॥੧॥",
		696},
	{"2023-07-25.html",
		"2023-07-25 04:45",
		"hir isau pRIiq AMqru mnu byiDAw hir ibnu rhxu n jweI ] ijau mCulI ibnu nIrY ibnsY iqau nwmY ibnu mir jweI ]1] myry pRB ikrpw jlu dyvhu hir nweI ] hau AMqir nwmu mMgw idnu rwqI nwmy hI sWiq pweI ] rhwau ] ijau cwiqRku jl ibnu ibllwvY ibnu jl ipAws n jweI ] gurmuiK jlu pwvY suK shjy hirAw Bwie suBweI ]2] mnmuK BUKy dh ids folih ibnu nwvY duKu pweI ] jnim mrY iPir jonI AwvY drgh imlY sjweI ]3] ikRpw krih qw hir gux gwvh hir rsu AMqir pweI ] nwnk dIn dieAwl Bey hY iqRsnw sbid buJweI ]4]8]",
		607},
	{"2023-07-26.html",
		"2023-07-26 05:30",
		"ਏ ਮਨ ਮੇਰਿਆ ਆਵਾ ਗਉਣੁ ਸੰਸਾਰੁ ਹੈ ਅੰਤਿ ਸਚਿ ਨਿਬੇੜਾ ਰਾਮ ॥ ਆਪੇ ਸਚਾ ਬਖਸਿ ਲਏ ਫਿਰਿ ਹੋਇ ਨ ਫੇਰਾ ਰਾਮ ॥ ਫਿਰਿ ਹੋਇ ਨ ਫੇਰਾ ਅੰਤਿ ਸਚਿ ਨਿਬੇੜਾ ਗੁਰਮੁਖਿ ਮਿਲੈ ਵਡਿਆਈ ॥ ਸਾਚੈ ਰੰਗਿ ਰਾਤੇ ਸਹਜੇ ਮਾਤੇ ਸਹਜੇ ਰਹੇ ਸਮਾਈ ॥ ਸਚਾ ਮਨਿ ਭਾਇਆ ਸਚੁ ਵਸਾਇਆ ਸਬਦਿ ਰਤੇ ਅੰਤਿ ਨਿਬੇਰਾ ॥ ਨਾਨਕ ਨਾਮਿ ਰਤੇ ਸੇ ਸਚਿ ਸਮਾਣੇ ਬਹੁਰਿ ਨ ਭਵਜਲਿ ਫੇਰਾ ॥੧॥ ਮਾਇਆ ਮੋਹੁ ਸਭੁ ਬਰਲੁ ਹੈ ਦੂਜੈ ਭਾਇ ਖੁਆਈ ਰਾਮ ॥ ਮਾਤਾ ਪਿਤਾ ਸਭੁ ਹੇਤੁ ਹੈ ਹੇਤੇ ਪਲਚਾਈ ਰਾਮ ॥ ਹੇਤੇ ਪਲਚਾਈ ਪੁਰਬਿ ਕਮਾਈ ਮੇਟਿ ਨ ਸਕੈ ਕੋਈ ॥ ਜਿਨਿ ਸ੍ਰਿਸਟਿ ਸਾਜੀ ਸੋ ਕਰਿ ਵੇਖੈ ਤਿਸੁ ਜੇਵਡੁ ਅਵਰੁ ਨ ਕੋਈ ॥ ਮਨਮੁਖਿ ਅੰਧਾ ਤਪਿ ਤਪਿ ਖਪੈ ਬਿਨੁ ਸਬਦੈ ਸਾਂਤਿ ਨ ਆਈ ॥ ਨਾਨਕ ਬਿਨੁ ਨਾਵੈ ਸਭੁ ਕੋਈ ਭੁਲਾ ਮਾਇਆ ਮੋਹਿ ਖੁਆਈ ॥੨॥ ਏਹੁ ਜਗੁ ਜਲਤਾ ਦੇਖਿ ਕੈ ਭਜਿ ਪਏ ਹਰਿ ਸਰਣਾਈ ਰਾਮ ॥ ਅਰਦਾਸਿ ਕਰੀ ਗੁਰ ਪੂਰੇ ਆਗੈ ਰਖਿ ਲੇਵਹੁ ਦੇਹੁ ਵਡਾਈ ਰਾਮ ॥ ਰਖਿ ਲੇਵਹੁ ਸਰਣਾਈ ਹਰਿ ਨਾਮੁ ਵਡਾਈ ਤੁਧੁ ਜੇਵਡੁ ਅਵਰੁ ਨ ਦਾਤਾ ॥ ਸੇਵਾ ਲਾਗੇ ਸੇ ਵਡਭਾਗੇ ਜੁਗਿ ਜੁਗਿ ਏਕੋ ਜਾਤਾ ॥ ਜਤੁ ਸਤੁ ਸੰਜਮੁ ਕਰਮ ਕਮਾਵੈ ਬਿਨੁ ਗੁਰ ਗਤਿ ਨਹੀ ਪਾਈ ॥ ਨਾਨਕ ਤਿਸ ਨੋ ਸਬਦੁ ਬੁਝਾਏ ਜੋ ਜਾਇ ਪਵੈ ਹਰਿ ਸਰਣਾਈ ॥੩॥ ਜੋ ਹਰਿ ਮਤਿ ਦੇਇ ਸਾ ਊਪਜੈ ਹੋਰ ਮਤਿ ਨ ਕਾਈ ਰਾਮ ॥ ਅੰਤਰਿ ਬਾਹਰਿ ਏਕੁ ਤੂ ਆਪੇ ਦੇਹਿ ਬੁਝਾਈ ਰਾਮ ॥ ਆਪੇ ਦੇਹਿ ਬੁਝਾਈ ਅਵਰ ਨ ਭਾਈ ਗੁਰਮੁਖਿ ਹਰਿ ਰਸੁ ਚਾਖਿਆ ॥ ਦਰਿ ਸਾਚੈ ਸਦਾ ਹੈ ਸਾਚਾ ਸਾਚੈ ਸਬਦਿ ਸੁਭਾਖਿਆ ॥ ਘਰ ਮਹਿ ਨਿਜ ਘਰੁ ਪਾਇਆ ਸਤਿਗੁਰੁ ਦੇਇ ਵਡਾਈ ॥ ਨਾਨਕ ਜੋ ਨਾਮਿ ਰਤੇ ਸੇਈ ਮਹਲੁ ਪਾਇਨਿ ਮਤਿ ਪਰਵਾਣੁ ਸਚੁ ਸਾਈ ॥੪॥੬॥",
		571},
	{"2023-08-01.html",
		"2023-08-01 04:45",
		"ਹਮਰੀ ਗਣਤ ਨ ਗਣੀਆ ਕਾਈ ਅਪਣਾ ਬਿਰਦੁ ਪਛਾਣਿ ॥ ਹਾਥ ਦੇਇ ਰਾਖੇ ਕਰਿ ਅਪੁਨੇ ਸਦਾ ਸਦਾ ਰੰਗੁ ਮਾਣਿ ॥੧॥ ਸਾਚਾ ਸਾਹਿਬੁ ਸਦ ਮਿਹਰਵਾਣੁ ॥ ਬੰਧੁ ਪਾਇਆ ਮੇਰੈ ਸਤਿਗੁਰਿ ਪੂਰੈ ਹੋਈ ਸਰਬ ਕਲਿਆਣ ॥ ਰਹਾਉ ॥ ਜੀਉ ਪਾਇ ਪਿੰਡੁ ਜਿਨਿ ਸਾਜਿਆ ਦਿਤਾ ਪੈਨਣੁ ਖਾਣੁ ॥ ਅਪਣੇ ਦਾਸ ਕੀ ਆਪਿ ਪੈਜ ਰਾਖੀ ਨਾਨਕ ਸਦ ਕੁਰਬਾਣੁ ॥੨॥੧੬॥੪੪॥",
		619},
	{"2023-08-02.html",
		"2023-08-02 04:45",
		"ਪੂਰਬਿ ਲਿਖਿਆ ਕਮਾਵਣਾ ਜਿ ਕਰਤੈ ਆਪਿ ਲਿਖਿਆਸੁ ॥ ਮੋਹ ਠਗਉਲੀ ਪਾਈਅਨੁ ਵਿਸਰਿਆ ਗੁਣਤਾਸੁ ॥ ਮਤੁ ਜਾਣਹੁ ਜਗੁ ਜੀਵਦਾ ਦੂਜੈ ਭਾਇ ਮੁਇਆਸੁ ॥ ਜਿਨੀ ਗੁਰਮੁਖਿ ਨਾਮੁ ਨ ਚੇਤਿਓ ਸੇ ਬਹਣਿ ਨ ਮਿਲਨੀ ਪਾਸਿ ॥ ਦੁਖੁ ਲਾਗਾ ਬਹੁ ਅਤਿ ਘਣਾ ਪੁਤੁ ਕਲਤੁ ਨ ਸਾਥਿ ਕੋਈ ਜਾਸਿ ॥ ਲੋਕਾ ਵਿਚਿ ਮੁਹੁ ਕਾਲਾ ਹੋਆ ਅੰਦਰਿ ਉਭੇ ਸਾਸ ॥ ਮਨਮੁਖਾ ਨੋ ਕੋ ਨ ਵਿਸਹੀ ਚੁਕਿ ਗਇਆ ਵੇਸਾਸੁ ॥ ਨਾਨਕ ਗੁਰਮੁਖਾ ਨੋ ਸੁਖੁ ਅਗਲਾ ਜਿਨਾ ਅੰਤਰਿ ਨਾਮ ਨਿਵਾਸੁ ॥੧॥ ਮਃ ੩ ॥ ਸੇ ਸੈਣ ਸੇ ਸਜਣਾ ਜਿ ਗੁਰਮੁਖਿ ਮਿਲਹਿ ਸੁਭਾਇ ॥ ਸਤਿਗੁਰ ਕਾ ਭਾਣਾ ਅਨਦਿਨੁ ਕਰਹਿ ਸੇ ਸਚਿ ਰਹੇ ਸਮਾਇ ॥ ਦੂਜੈ ਭਾਇ ਲਗੇ ਸਜਣ ਨ ਆਖੀਅਹਿ ਜਿ ਅਭਿਮਾਨੁ ਕਰਹਿ ਵੇਕਾਰ ॥ ਮਨਮੁਖ ਆਪ ਸੁਆਰਥੀ ਕਾਰਜੁ ਨ ਸਕਹਿ ਸਵਾਰਿ ॥ ਨਾਨਕ ਪੂਰਬਿ ਲਿਖਿਆ ਕਮਾਵਣਾ ਕੋਇ ਨ ਮੇਟਣਹਾਰੁ ॥੨॥ ਪਉੜੀ ॥ ਤੁਧੁ ਆਪੇ ਜਗਤੁ ਉਪਾਇ ਕੈ ਆਪਿ ਖੇਲੁ ਰਚਾਇਆ ॥ ਤ੍ਰੈ ਗੁਣ ਆਪਿ ਸਿਰਜਿਆ ਮਾਇਆ ਮੋਹੁ ਵਧਾਇਆ ॥ ਵਿਚਿ ਹਉਮੈ ਲੇਖਾ ਮੰਗੀਐ ਫਿਰਿ ਆਵੈ ਜਾਇਆ ॥ ਜਿਨਾ ਹਰਿ ਆਪਿ ਕ੍ਰਿਪਾ ਕਰੇ ਸੇ ਗੁਰਿ ਸਮਝਾਇਆ ॥ ਬਲਿਹਾਰੀ ਗੁਰ ਆਪਣੇ ਸਦਾ ਸਦਾ ਘੁਮਾਇਆ ॥੩॥",
		643},
}

// Ensure that a webserver is running on 8081 with some test files
// this is more of an integration test than a unit test, but not gonna be strict about it as my skills are not there yet
// docker run -it -p 8081:80 -v ${PWD}:/usr/share/nginx/html nginx
func TestHukamnama(t *testing.T) {

	for _, wants := range testData {
		hukam := GetTodaysHukamnama(fmt.Sprintf("http://localhost:8081/%s", wants.filename))

		//fmt.Printf("hukam object %v", hukam)
		if hukam.dateTime != wants.dateTime {
			t.Fatalf("Hukamnama date for %s is wrong, wants %s got %s,", wants.filename, wants.dateTime, hukam.dateTime)
		}
		if hukam.ang != wants.ang {
			t.Fatalf("Hukamnama ang for %s is wrong, wants %d got %d,", wants.filename, hukam.ang, wants.ang)
		}
		if hukam.gurmukhi != wants.gurmukhi {
			t.Fatalf("Hukamnama Gurmukhi for %s is wrong, wants %s got %s,", wants.filename, hukam.gurmukhi, wants.gurmukhi)
		}
	}
}
