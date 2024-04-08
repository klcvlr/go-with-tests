package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		_, _ = fmt.Fprintf(w, data)
	}
}
