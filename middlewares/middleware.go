package middlewares

import (
	"fmt"
	"net/http"
)

func LogReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
