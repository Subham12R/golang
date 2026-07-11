package concurrency

import (
	"testing"
	"time"
)

func slowStubber(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "https://www.example.com"
	}

	for b.Loop() {
		CheckWebsites(slowStubber, urls)
	}
}
