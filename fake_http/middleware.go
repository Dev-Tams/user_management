package request

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler{
		start := time.Now()
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		now := time.Now()
		fmt.Printf("[%s] %s %v \n", r.Method, r.URL.Path, now.Sub(start))
		next.ServeHTTP(w, r)
	})
}