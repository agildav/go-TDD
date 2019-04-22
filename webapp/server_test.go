package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("returns a players score", func(t *testing.T) {
		want := "20"

		request := httptest.NewRequest(http.MethodGet, "/players/Pablo", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		assertResponse(t, got, want)

	})

}

// // // // // // // // // // // // // // // // // // // // // // // // // //

// Helpers
func assertResponse(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		errMsg := fmt.Sprintf("got -> %s, want -> %s", got, want)
		t.Errorf(colorError(errMsg))
	}
}
