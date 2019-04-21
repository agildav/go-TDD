package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increments 1000 times safely concurrently", func(t *testing.T) {
		counter := newCounter()
		const n = 1000

		var wg sync.WaitGroup
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, counter, n)
	})

}

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Helpers
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		errMsg := fmt.Sprintf("got -> %d, want -> %d", got.Value(), want)
		t.Errorf(colorError(errMsg))
	}
}

func newCounter() *Counter {
	return &Counter{}
}
