package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	// ? Case 1 - Comparing 2 website, which one rendered faster
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {

		// slowUrl := "http://facebook.com"
		// fastUrl := "http://www.quii.dev"

		// want := fastUrl

		// got := WebsiteRacer(slowUrl, fastUrl)

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		// ? Since we're using test
		// ? and test should not use the http directly
		// ? Now we will mock the http get using the net/http/httptest

		// ! This code is not DRY and it's not good in golang, so we need to refactor it
		// slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	time.Sleep(20 * time.Millisecond)
		// 	w.WriteHeader(http.StatusOK)
		// }))

		// fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	w.WriteHeader(http.StatusOK)
		// }))

		// slowUrl := slowServer.URL
		// fastUrl := fastServer.URL

		// want := fastUrl
		// got := WebsiteRacer(slowUrl, fastUrl)

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		// slowServer.Close()
		// fastServer.Close()

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// ? defer - run this at the end of the function
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		// ? This one will use WebsiteRacer
		got, err := WebsiteRacer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// ? Case 2 - Will return error if more than 10s
	t.Run("return error if doesn't respond within 10 second", func(t *testing.T) {
		// ? Since everything is now just a mockup, we won't make 2 of it
		// serverA := makeDelayedServer(10 * time.Millisecond)
		// serverB := makeDelayedServer(25 * time.Millisecond)

		// defer serverA.Close()
		// defer serverB.Close()

		server := makeDelayedServer(25 * time.Millisecond)

		// This one will use ConfigurableWebsiteRacer
		_, err := ConfigurableWebsiteRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected an error, didn't get one")
		}
	})
}

// ? We will make the code DRY by making a new function - to use the delayed server
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
