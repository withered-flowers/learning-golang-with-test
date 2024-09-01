package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// ? Only simple mock checker
func mockWebsiteChecker(url string) bool {
	return url != "waat://furtherwe.geds"
}

// ? Slow (sleep) mock delay web checking
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.halodunia.com",
		"waat://furtherwe.geds",
	}

	want := map[string]bool{
		"http://google.com":         true,
		"http://blog.halodunia.com": true,
		"waat://furtherwe.geds":     false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// ? Since we're using goroutines, we need to be careful
// ! Can error "concurrent map writes" / "data race"
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
