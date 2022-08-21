package service

import (
	"github.com/JanMeckelholt/running/tree/main/running-backend/model"
	"github.com/JanMeckelholt/running/tree/main/running-backend/persistence"
)

type Runnerer struct {
	RunnerStorer persistence.RunnerStorer
}

type RunnerService interface {
	CreateRunner(email string, password string) (*model.Runner, error)
	GetRunner() ([]model.Runner, error)
	GetRunnerByID(ID uint) (*model.Runner, error)
	GetRunnerByEmail(email string) (*model.Runner, error)
	UpdateRunnerByID(ID uint, updatedRunner *model.Runner) error
	DeleteRunnerByID(ID uint) error
}

func NewRunnerService(storer persistence.RunnerStorer) Runnerer {
	return Runnerer{
		RunnerStorer: storer,
	}
}

func (rs Runnerer) CreateRunner(email string, password string) (*model.Runner, error) {
	runnerClearPassword := model.RunnerClearPassword{
		Email:    email,
		Password: password,
	}
	runner, err := runnerClearPassword.ToRunner()
	if err != nil {
		return nil, err
	}
	runner, err = rs.RunnerStorer.CreateRunner(runner.Email, runner.PasswordHash)

	return runner, err
}

func (rs Runnerer) GetAllRunner() ([]model.Runner, error) {
	return rs.RunnerStorer.GetRunner()
}

func (rs Runnerer) GetRunnerByID(ID uint) (*model.Runner, error) {
	return rs.RunnerStorer.GetRunnerByID(ID)
}

func (rs Runnerer) GetRunnerByEmail(email string) (*model.Runner, error) {
	return rs.RunnerStorer.GetRunnerByEmail(email)
}

func (rs Runnerer) UpdateRunnerByID(ID uint, updatedRunnerClearPassword *model.RunnerClearPassword) error {
	updatedRunner, err := updatedRunnerClearPassword.ToRunner()
	if err != nil {
		return err
	}
	return rs.RunnerStorer.UpdateRunnerByID(ID, updatedRunner)
}

func (rs Runnerer) DeleteRunnerByID(ID uint) error {
	err := rs.RunnerStorer.DeleteRunnerByID(ID)
	return err
}
