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
		city := &model.City{}
		db.Where("id = ?", (*form)[i].DepartureCityId).First(city)
		(*form)[i].City = *city
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(find)
})
