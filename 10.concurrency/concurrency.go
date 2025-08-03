package concurrency

/*
An operation that does not block in Go will run in a separate process
called a goroutine.

To tell Go to start a new goroutine, we turn a function call into a go statement by putting the keyword go in front of it.
*/

type WebsiteChecker func(string) bool
type result struct {
	url string
	res bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url: url, res: wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.url] = r.res
	}
	return results
}
