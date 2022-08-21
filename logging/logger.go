package logging

import (
	"github.com/JanMeckelholt/running/tree/main/running-backend/config"
	log "github.com/sirupsen/logrus"
)

func SetupLogger(config config.LogConfig) {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(config.Level)
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}
