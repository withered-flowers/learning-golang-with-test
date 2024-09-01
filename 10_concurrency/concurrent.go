package concurrency

type WebsiteChecker func(string) bool

// ? Now we need to define the channel result struct
type result struct {
	// ? string for the web address
	string
	// ? bool for the validation result
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// ? Concurrency need to be combined with channels
	// ? Channels are the sender and receiver (take it as "room" in socket.io)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// results[u] = wc(u)

			// ? Now instead of assigning to wc directly
			// ? We will send the value via channel
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	// ? Sleep is not needed anymore
	// time.Sleep(2 * time.Second)

	// ? Now we will need to add logic to receive the sended value via channel
	for i := 0; i < len(urls); i++ {
		// ? Now we will receive the channel value, which is result struct
		r := <-resultsChannel

		// ? We will assign the received value to results
		results[r.string] = r.bool
	}

	return results
}
