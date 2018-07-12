package main

import (
	"log"
	"net/http"
	"web/test-graphql/graphql"
)

func main() {
	http.Handle("/api/graphql", graphql.Schema)
	log.Print("ready: listening...\n")
	http.ListenAndServe(":6666", nil)
}

// func disableCors(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
// 		if r.Method == "OPTIONS" {
// 			w.Header().Set("Access-Control-Max-Age", "86400")
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }
