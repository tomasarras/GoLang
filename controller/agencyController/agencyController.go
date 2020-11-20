package agencyController

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/tomasarras/GoLang/entity"
	"github.com/tomasarras/GoLang/service/agencyService"
)

var serviceAgency agencyService.ServiceAgency

var find httptreemux.HandlerFunc

func Start(db *sql.DB) {
	serviceAgency, _ = agencyService.New(db)
}

func SaveAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	var a entity.Agency
	_ = json.NewDecoder(r.Body).Decode(&a)

	if !a.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, _ := serviceAgency.Save(a)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func FindByIdAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
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

	json.NewEncoder(w).Encode(a)
}

func FindAllAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceAgency.FindAll())
}

func UpdateAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a entity.Agency
	_ = json.NewDecoder(r.Body).Decode(&a)

	if !a.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ag, _ := serviceAgency.FindByID(id)
	if ag.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.ID = int64(id)
	serviceAgency.Update(a)
	json.NewEncoder(w).Encode(a)
}

func RemoveAgencyHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = serviceAgency.Remove(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
