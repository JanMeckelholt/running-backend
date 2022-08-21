package handler

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Healthcheck run")
}
