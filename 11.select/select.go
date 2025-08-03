package racer

import (
	"fmt"
	"net/http"
	"time"
)

// ping returns a channel that will be closed when a GET request to the given URL completes
// use `struct{}` because we are closing and not sending anything on the chan. since `struct{}` is the smallest data type available from a memory prespective, we get no allocation (vs a `bool`)
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// whichever one writes to its channel first will have its code executed in the `select`, which results in its `URL` being returned and being the winner
func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}
