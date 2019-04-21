package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// Deposit
	t.Run("return the fastest URL check", func(t *testing.T) {
		slowServer := makeTestServer(30 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		// Close the servers
		defer slowServer.Close()
		defer fastServer.Close()

		slowServerURL := slowServer.URL
		fastServerURL := fastServer.URL

		want := fastServerURL
		got := Racer(slowServerURL, fastServerURL)

		assertURL(t, got, want)
	})
}

// // // // // // // // // // // // // // // // // // // // // // // // // //
func makeTestServer(ms time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(ms)
		w.WriteHeader(http.StatusOK)
	}))
}

// Helpers
func assertURL(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
		t.Errorf(colorError(errMsg))
	}
}
