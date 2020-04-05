package model

import "github.com/jinzhu/gorm"

type ApplicationForm struct {
	gorm.Model

	Id          int    `json:"id"`
	MissionDate string `json:"mission_date"`
}
