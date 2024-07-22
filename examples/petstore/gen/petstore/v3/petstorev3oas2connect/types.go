// Code generated by protoc-gen-oas2connect. DO NOT EDIT.

package petstorev3oas2connect

import (
	"net/http"
)

type ServeMux interface {
	Handle(path string, handler http.Handler)
}

type Middleware func(next http.Handler) http.Handler

func NoopMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}