package racer

import (
	"net/http"
)

func Racer(slowURL, fastURL string) (URL string) {
	select {
	case <-ping(slowURL):
		return slowURL
	case <-ping(fastURL):
		return fastURL
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)

	go func() {
		_, err := http.Get(URL)

		if err != nil {
			ch <- false
		}
		ch <- true
	}()

	return ch
}
