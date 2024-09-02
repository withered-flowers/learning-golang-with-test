package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func WebsiteRacer(a, b string) (winner string, err error) {
	return ConfigurableWebsiteRacer(a, b, tenSecondTimeout)
}

// ? Need to use 2 return variable, since we need to return err
// ? Need to supply timeout duration, for testing purpose
// ? We will rename this to "Configurable", since we need to DI the duration
func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// ! These code are not DRY
	// startA := time.Now()
	// http.Get(a)
	// durationA := time.Since(startA)

	// startB := time.Now()
	// http.Get(b)
	// durationB := time.Since(startB)

	// ? Now usnig new function, this will be DRY
	// durationA := measureResponseTime(a)
	// durationB := measureResponseTime(b)

	// if durationA < durationB {
	// 	return a
	// }

	// return b

	// ? Now we can use "select" for the new logic
	// ? Find which one which return faster
	// ? So we don't need to calculate

	// ? Select will wait on MULTIPLE CHANNEL
	// ? Which one first, will go to the case
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ? To make it DRY we need to make a new function
// ! Since we're using channel, we don't need this anymore
// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// ? Go is good at concurrency isn't it?
// ? Why need to calculate the time?
// ? Change the logic: find which one fastest, return it
// ? -> No need to calculate
// ? -> We can use channel
func ping(url string) chan struct{} {
	// ! chan struct{} IS THE SMALLEST data type
	// ! Not boolean (from memory perspective)
	ch := make(chan struct{})

	// ? We will use goroutine
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
