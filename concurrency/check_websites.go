package concurrency

// WebsiteChecker which takes a single URL and returns a boolean.
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// Alongside the results map we now have a resultChannel, which we make in the same way.
	// chan result is the type of the channel - a channel of result.
	resultChannel := make(chan result)

	for _, url := range urls {
		// we often use anonymous functions when we want to start a goroutine. An anonymous
		// function literal looks just the same as a normal function declaration, but without
		// a name (unsurprisingly).
		// go func() {
		// 	results[url] = wc(url)
		// }()
		// The () at the end will execute the function upon declaration.

		// To properly have our routines work togther however we have to do it this way:
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
		// By giving each anonymous function a parameter for the url - u - and then calling
		// the anonymous function with the url as the argument, we make sure that the value
		// of u is fixed as the value of url for the iteration of the loop that we're launching
		// the goroutine in. u is a copy of the value of url, and so can't be changed.
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
