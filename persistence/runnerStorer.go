package persistence

import (
	"errors"
	"github.com/JanMeckelholt/running/tree/main/running-backend/model"
	log "github.com/sirupsen/logrus"
)

const RunnerNotFound string = "Runner not found."

type RunnerStorage struct {
	storer Storer
}

type RunnerStorer interface {
	CreateRunner(email string, passwordHash string) (*model.Runner, error)
	GetRunner() ([]model.Runner, error)
	GetRunnerByID(ID uint) (*model.Runner, error)
	GetRunnerByEmail(email string) (*model.Runner, error)
	UpdateRunnerByID(ID uint, updatedRunner *model.Runner) error
	DeleteRunnerByID(ID uint) error
}

func NewRunnerStorer(storer Storer) RunnerStorer {
	runnerStorage := RunnerStorage{
		storer: storer,
	}
	runnerStorage.Init()
	return runnerStorage
}

func (rs RunnerStorage) Init() {
	err := rs.storer.InitStorage()
	if err != nil {
		panic(err)
	}
	if err := rs.storer.AutoMigrate(&model.Runner{}); err != nil {
		panic(err)
	}
}

func (rs RunnerStorage) CreateRunner(email string, passwordHash string) (*model.Runner, error) {
	runner := model.Runner{
		Email:        email,
		PasswordHash: passwordHash,
	}
	err := rs.storer.Create(&runner)
	if err != nil {
		return nil, err
	}
	log.Infof("Successfully stored new runner with ID %v in database.", runner.ID)
	log.Tracef("Stored: %v", runner)
	return &runner, nil
}

func (rs RunnerStorage) GetRunner() ([]model.Runner, error) {
	var runner []model.Runner
	_, err := rs.storer.Get("", &runner)
	if err != nil {
		return nil, err
	}
	log.Infof("Successfully received runner from database.")

	return runner, nil
}

func (rs RunnerStorage) GetRunnerByID(ID uint) (*model.Runner, error) {
	var runner *model.Runner
	_, err := rs.storer.GetObjectById("", "", ID, &runner)
	if err != nil {
		return nil, err
	}
	if runner.ID == 0 {
		return nil, nil
	}
	log.Info("Successfully received Runner from database.")
	return runner, nil
}

func (rs RunnerStorage) GetRunnerByEmail(email string) (*model.Runner, error) {
	var runner *model.Runner
	_, err := rs.storer.GetWhere("", &model.Runner{Email: email}, runner)
	if err != nil {
		return nil, err
	}
	if runner.ID == 0 {
		return nil, nil
	}
	log.Info("Successfully received Runner from database.")
	return runner, nil
}

func (rs RunnerStorage) UpdateRunnerByID(ID uint, updatedRunner *model.Runner) error {
	savedRunner, err := rs.GetRunnerByID(ID)
	if err != nil {
		return err
	}
	if savedRunner == nil {
		return errors.New(RunnerNotFound)
	}
	return rs.storer.UpdateObject(&savedRunner, updatedRunner)
}

func (rs RunnerStorage) DeleteRunnerByID(ID uint) error {
	runner, err := rs.GetRunnerByID(ID)
	if err != nil {
		return err
	}
	if runner == nil {
		return errors.New(RunnerNotFound)
	}
	return rs.storer.DeleteObjectById(ID, runner)
}
