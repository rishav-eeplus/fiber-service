package main

import (
	"log"
	"net/http"
)

// middleware for setting content type json and setting origin
func HeadersCustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if Mode == "prod" {
			w.Header().Set("Access-Control-Allow-Origin", "https://portal.eeplus.com")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		next.ServeHTTP(w, req)
	})
}

// for handeling uncaught exceptions
func ExceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf("something went wrong: %v", err)
				http.Error(w, "Something went very wrong", 500)
			}
		}()
		next.ServeHTTP(w, req)
	})
}