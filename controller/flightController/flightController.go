package flightController

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tomasarras/GoLang/entity"
	"github.com/tomasarras/GoLang/service/agencyService"
	"github.com/tomasarras/GoLang/service/flightService"
)

var serviceFlight flightService.ServiceFlight
var serviceAgency agencyService.ServiceAgency

func Start(db *sql.DB) {
	serviceFlight, _ = flightService.New(db)
	serviceAgency, _ = agencyService.New(db)
}

func SaveFlightHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["idAgency"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, _ := serviceAgency.FindByID(id)
	if a.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var f entity.Flight
	_ = json.NewDecoder(r.Body).Decode(&f)

	if !f.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f.IdAgency = int64(id)
	response, _ := serviceFlight.Save(f)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func FindByIdFlightHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, _ := serviceFlight.FindByID(id)
	if f.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(f)
}

func FindAllFlightHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceFlight.FindAll())
}

func FindAllFlightByAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["idAgency"]
	id, err := strconv.Atoi(param)
	itsId := err == nil
	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, _ := serviceAgency.FindByID(id)
	if a.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(serviceFlight.FindAllByAgency(id))
}

func UpdateFlightHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil
	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var f entity.Flight
	_ = json.NewDecoder(r.Body).Decode(&f)

	if !f.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fl, _ := serviceFlight.FindByID(id)
	if fl.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f.ID = int64(id)

	response, _ := serviceFlight.Update(f)
	if response.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}

func RemoveFlightHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = serviceFlight.Remove(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
