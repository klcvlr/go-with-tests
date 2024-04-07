package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("get the faster server url", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowServer := makeDelayedServer(5 * time.Millisecond)
		defer fastServer.Close()
		defer slowServer.Close()

		actual, err := Racer(fastServer.URL, slowServer.URL)

		if err != nil {
			t.Fatal("Got an error but didn't expect one", err)
		}
		if actual != fastServer.URL {
			t.Errorf("Expected %q but got %q", fastServer.URL, actual)
		}
	})

	t.Run("error when response takes longer than 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(10 * time.Millisecond)
		defer server.Close()

		actual, err := ConfigurableRacer(server.URL, server.URL, 5*time.Millisecond)

		if err == nil {
			t.Fatal("Expected to have an error but didn't get one", err)
		}
		if actual != "" {
			t.Errorf("Expected %q to be empty", actual)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
