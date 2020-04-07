package model

import (
	"time"
)

type ApplicationForm struct {
	Id                uint      `json:"id"`
	MissionDate       time.Time `json:"mission_date"`
	DepartureCity     City
	DestinationCity   City
	DepartureCityId   uint `json:"-"`
	DestinationCityId uint `json:"-"`
}

func (*ApplicationForm) TableName() string {
	return "application_form"
}
