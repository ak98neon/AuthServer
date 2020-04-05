package model

import (
	"time"
)

type ApplicationForm struct {
	Id              uint      `json:"id"`
	MissionDate     time.Time `json:"mission_date"`
	DepartureCityId uint
}

func (*ApplicationForm) TableName() string {
	return "application_form"
}
