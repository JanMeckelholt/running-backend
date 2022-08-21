package handler

import (
	"fmt"
	"github.com/JanMeckelholt/running/tree/main/running-backend/config"
	"github.com/JanMeckelholt/running/tree/main/running-backend/middleware"
	"github.com/JanMeckelholt/running/tree/main/running-backend/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Service struct {
	router        *mux.Router
	config        config.RestConfig
	runnerService service.Runnerer
}

func NewAPI(config config.RestConfig, runnerService service.Runnerer) Service {

	return Service{
		router:        mux.NewRouter(),
		config:        config,
		runnerService: runnerService,
	}
}

func (s Service) RegisterMiddleware() {
	if s.config.UsePasswordMiddleware {
		s.router.Use(middleware.PasswordProtectionMiddleware)
	}
	if s.config.UseCorsMiddleware {
		s.router.Use(middleware.CorsMiddleware)
	}
}

func (s Service) RegisterHttpRoutes() {
	s.router.HandleFunc("/health", Health).Methods(http.MethodGet)

	s.router.HandleFunc("/runner", s.CreateRunner).Methods(http.MethodPost)
	s.router.HandleFunc("/runner/{runnerID}", s.GetRunnerByID).Methods(http.MethodGet)
	s.router.HandleFunc("/runner", s.GetAllRunner).Methods(http.MethodGet)
	s.router.HandleFunc("/runner/{runnerID}", s.UpdateRunnerByID).Methods(http.MethodPut)
	s.router.HandleFunc("/runner/{runnerID}", s.DeleteRunner).Methods(http.MethodDelete)

}

func (s Service) ServeREST() error {
	log.WithField("Port", s.config.Port).Info("REST Server running")
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), s.router)
}
