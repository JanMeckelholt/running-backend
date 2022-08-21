package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Runner struct {
	gorm.Model
	Email        string      `gorm:"notNull"`
	PasswordHash string      `gorm:"notNull"`
	Exercises    []*Exercise `gorm:"foreignKey:RunnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type RunnerClearPassword struct {
	Email    string `gorm:"notNull"`
	Password string `gorm:"notNull"`
}

func (rCP *RunnerClearPassword) ToRunner() (*Runner, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rCP.Password), 8)
	if err != nil {
		return nil, err
	}
	runner := Runner{
		Email:        rCP.Email,
		PasswordHash: string(passwordHash)}
	return &runner, nil
}
