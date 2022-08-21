package model

import (
	"gorm.io/gorm"
	"time"
)

const DateLayout = "2.1.2006"

type Exercise struct {
	gorm.Model
	Type          string `gorm:"default:'running'"`
	Distance      float64
	Duration      time.Duration
	Date          time.Time
	Remarks       string
	Place         string
	avgPace       time.Duration `gorm:"-"`
	RunningWeekID string
	RunnerID      uint
}

func (e *Exercise) AfterFind(tx *gorm.DB) (err error) {
	e.avgPace = time.Duration(float64(e.Duration) / e.Distance)
	return nil
}

type RunningWeek struct {
	gorm.Model
	ID        string     `gorm:"primaryKey"` //RunnerID-Year-CalenderWeek
	Exercises []Exercise `gorm:"foreignKey:RunningWeekID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StartDate time.Time
	KmTotal   float64       `gorm:"-"`
	TimeTotal time.Duration `gorm:"-"`
	avgPace   time.Duration `gorm:"-"`
	Remarks   string
	Goals     string
}

func (rw *RunningWeek) AfterFind(tx *gorm.DB) (err error) {
	var kmTotal float64
	result := tx.Model(&Exercise{}).Select("ifnull(sum(Distance),0)").Where("running_week_id = ?", rw.ID).Scan(&kmTotal)
	if result.Error != nil {
		return result.Error
	}
	rw.KmTotal = kmTotal

	var timeTotal time.Duration
	result = tx.Model(&Exercise{}).Select("ifnull(sum(Duration),0)").Where("running_week_id = ?", rw.ID).Scan(&timeTotal)
	if result.Error != nil {
		return result.Error
	}
	rw.TimeTotal = timeTotal

	rw.avgPace = time.Duration(float64(rw.TimeTotal) / rw.KmTotal)

	return nil
}
