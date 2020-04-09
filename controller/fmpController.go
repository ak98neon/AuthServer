package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ak98neon/authserver/model"
	"io/ioutil"
	"net/http"
)

type FormApprove struct {
	ID string `json:"id"`
}

var ApproveForm = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	approveForm := &FormApprove{}
	err := json.NewDecoder(r.Body).Decode(approveForm)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid form id"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	url := "http://localhost:9090/forms/approve"
	fmt.Println("URL:>", url)

	jsonBody, _ := json.Marshal(approveForm)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := sendRequest(req)
	json.NewEncoder(w).Encode(resp)
	return
})

var RejectForm = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	approveForm := &FormApprove{}
	err := json.NewDecoder(r.Body).Decode(approveForm)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid form id"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	url := "http://localhost:9090/forms/reject"
	fmt.Println("URL:>", url)

	jsonBody, _ := json.Marshal(approveForm)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := sendRequest(req)
	json.NewEncoder(w).Encode(resp)
	return
})

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

func sendRequest(req *http.Request) map[string]interface{} {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"status": false, "message": "Main server return error"}
	}
	return map[string]interface{}{"status": true, "message": response}
}
