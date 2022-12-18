package handler

import (
	"fmt"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			fmt.Println(123)
			w.WriteHeader(http.StatusOK)
			return
		}
		fmt.Println("ok")
		// Next
		next.ServeHTTP(w, r)
		return
	})
}
