package context

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return http.Header{}
}

func (s *SpyResponseWriter) Write(_ []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(_ int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	resultCh := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("context cancelled")
				return
			default:
				time.Sleep(2 * time.Millisecond)
				result += string(c)
			}
		}
		resultCh <- result
	}()

	select {
	case data := <-resultCh:
		return data, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("Store was not cancelled")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("Store was cancelled")
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, World"
		store := &SpyStore{response: data, t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := &SpyResponseWriter{}
		ctx, cancelFunc := context.WithCancel(request.Context())
		request = request.WithContext(ctx)

		time.AfterFunc(5*time.Millisecond, cancelFunc)
		server.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
