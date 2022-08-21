package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func PasswordProtectionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var passwordHeaderKey = "runner_pw"
		if value, ok := os.LookupEnv("PASSWORD_HEADER_KEY"); ok {
			passwordHeaderKey = value
		} else {
			log.Error("Could not read passwordHeaderKey from env")
		}

		var passwordHeaderValue = ""
		if value, ok := os.LookupEnv("PASSWORD_HEADER_VALUE"); ok {
			passwordHeaderValue = value
		} else {
			log.Error("Could not read passwordHeaderValue from env")
		}

		// check for password header
		passwordHeader := r.Header.Get(passwordHeaderKey)
		if passwordHeader == passwordHeaderValue {
			next.ServeHTTP(w, r)
		} else {
			log.Error("Invalid password header provided")
			log.Error(passwordHeader)
			log.Error(passwordHeaderKey)
			log.Error(passwordHeaderValue)
			http.Error(w, "{'message': 'missing password'}", http.StatusForbidden)
		}
	})
}
