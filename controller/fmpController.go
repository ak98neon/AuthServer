package controller

import (
	"encoding/json"
	"github.com/ak98neon/authserver/model"
	"net/http"
)

var GetAllForms = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	form := &[]model.ApplicationForm{}

	find := db.Find(form)
	for i, _ := range *form {
		(*form)[i].DepartureCity = *findDepartureCity((*form)[i])
		(*form)[i].DestinationCity = *findDestinationCity((*form)[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(find)
})

func findDepartureCity(form model.ApplicationForm) *model.City {
	city := &model.City{}
	db.Where("id = ?", form.DepartureCityId).First(city)
	return city
}

func findDestinationCity(form model.ApplicationForm) *model.City {
	city := &model.City{}
	db.Where("id = ?", form.DestinationCityId).First(city)
	return city
}
