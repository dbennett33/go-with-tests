package testserver

import (
	"context"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data ,
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
