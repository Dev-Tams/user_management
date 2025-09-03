package request

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		start := time.Now()
		next.ServeHTTP(w, r)
		now := time.Now()
		fmt.Printf("[%s] %s %v \n", r.Method, r.URL.Path, now.Sub(start))
	})
}