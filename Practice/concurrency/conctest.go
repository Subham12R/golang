package concurrency

import(
	"reflect"
	"testing"
)

func mockWebsite(url string) bool {
	return url != "waat://test.geds"
}

func TestCheckWebsites(t *testing.T) {
	website := []string{
		"http://google.com",
		"http://wikipedia.com",
		"waat://test.geds",
	}

	want := map[string]bool {
		"http://google.com": true,
		"http://wikipedia.com": true,
		"waat://test.geds": false,
	}

	got := CheckWebsites(mockWebsite, website)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
