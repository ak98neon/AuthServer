package controller

import (
	"encoding/json"
	"net/http"
)

var GetAllForms = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	form := &ApplicationForm{}

	find := db.Find(form)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(find)
})
