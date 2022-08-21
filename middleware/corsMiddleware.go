package middleware

import (
	"net/http"
)

// CorsMiddleware Middleware
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, X-Requested-With, Origin, Content-Type, Accept-Encoding, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-CSRF-Token, Authorization, runner_pw")
		w.Header().Set("Access-Control-Max-Age", "86400") //1 day cache

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
