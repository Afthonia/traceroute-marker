package main

import (
	"context"
	"net/http"
)

type key int

const (
	DestKey key = iota
)

func DestAddressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		destAddress := r.URL.Query().Get("destination")

		if destAddress == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"error\":\"please provide a correct destination address\"}"))
			return
		}

		//log.Println(destAddress)

		ctx := r.Context()
		newCtx := context.WithValue(ctx, DestKey, destAddress)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
