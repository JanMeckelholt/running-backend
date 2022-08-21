package handler

import (
	"encoding/json"
	"github.com/JanMeckelholt/running/tree/main/running-backend/model"
	"github.com/JanMeckelholt/running/tree/main/running-backend/persistence"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s Service) CreateRunner(w http.ResponseWriter, r *http.Request) {
	runnerClearPassword, err := requestToRunnerClearPassword(r)
	if err != nil {
		log.Errorf("Can't serialize request body to runnerClearPassword-struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	runner, err := s.runnerService.CreateRunner(runnerClearPassword.Email, runnerClearPassword.Password)
	if err != nil {
		log.Errorf("Error calling service CreateRettungsmittel. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, runner)
}

func (s Service) GetRunnerByID(w http.ResponseWriter, r *http.Request) {
	runnerID, err := getRunnerID(r)
	if err != nil {
		log.Errorf("Error getting runnerID. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entry := log.WithField("runnerID", runnerID)
	runner, err := s.runnerService.GetRunnerByID(runnerID)
	if err != nil {
		entry.Errorf("Error calling service GetRunnerByID. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if runner == nil {
		entry.Errorln(persistence.RunnerNotFound)
		http.Error(w, persistence.RunnerNotFound, http.StatusBadRequest)
		return
	}
	sendJson(w, *runner)
}

func (s Service) GetAllRunner(w http.ResponseWriter, _ *http.Request) {
	allRunner, err := s.runnerService.GetAllRunner()
	if err != nil {
		log.Errorf("Error calling service GetAllRunnern. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, allRunner)
}

func (s Service) UpdateRunnerByID(w http.ResponseWriter, r *http.Request) {
	runnerID, err := getRunnerID(r)
	if err != nil {
		log.Errorf("Error getting runnerID. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entry := log.WithField("runnerID", runnerID)
	runnerClearPassword, err := requestToRunnerClearPassword(r)
	if err != nil {
		entry.Errorf("Can't serialize request body to runner-struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.runnerService.UpdateRunnerByID(runnerID, runnerClearPassword); err != nil {
		entry.Errorf("Error calling service UpdateRunnerByID: %v", err)
		var errorCode int
		switch err.Error() {
		case persistence.RunnerNotFound:
			errorCode = http.StatusBadRequest
		default:
			errorCode = http.StatusInternalServerError
		}
		http.Error(w, err.Error(), errorCode)
		return
	}
	runner, err := s.runnerService.GetRunnerByID(runnerID)
	if err != nil {
		entry.Errorf("Can't get updated RunnerById: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, runner)
}

func (s Service) DeleteRunner(w http.ResponseWriter, r *http.Request) {
	runnerID, err := getRunnerID(r)
	if err != nil {
		log.Errorf("Error getting RunnernID. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entry := log.WithField("runnerID", runnerID)
	toDeleteRunner, err := s.runnerService.GetRunnerByID(runnerID)
	if err != nil {
		entry.Errorf("Can't get to delete RunnerById: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.runnerService.DeleteRunnerByID(runnerID); err != nil {
		entry.Errorf("Error calling service DeleteRunnerByID. %v", err)
		var errorCode int
		switch err.Error() {
		case persistence.RunnerNotFound:
			errorCode = http.StatusBadRequest
		default:
			errorCode = http.StatusInternalServerError
		}
		http.Error(w, err.Error(), errorCode)
		return
	}

	sendJson(w, toDeleteRunner)
}

func requestToRunnerClearPassword(r *http.Request) (*model.RunnerClearPassword, error) {
	var runner model.RunnerClearPassword
	err := json.NewDecoder(r.Body).Decode(&runner)
	if err != nil {
		log.Errorf("Can't serialize request body to runner-struct. %v", err)
		return nil, err
	}
	return &runner, nil
}
