package main

import (
	"github.com/JanMeckelholt/running/tree/main/running-backend/config"
	"github.com/JanMeckelholt/running/tree/main/running-backend/handler"
	"github.com/JanMeckelholt/running/tree/main/running-backend/logging"
	"github.com/JanMeckelholt/running/tree/main/running-backend/persistence"
	"github.com/JanMeckelholt/running/tree/main/running-backend/service"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting running-backend API server")
	conf := config.NewConfig()
	logging.SetupLogger(conf.Logging)

	storage := persistence.NewStorage(conf.Persistence)
	RunnerStorer := persistence.NewRunnerStorer(storage)
	runnerService := service.NewRunnerService(RunnerStorer)

	api := handler.NewAPI(conf.Rest, runnerService)

	api.RegisterMiddleware()
	api.RegisterHttpRoutes()
	err := api.ServeREST()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Info("Recovered in f", r)
		}
	}()

}
