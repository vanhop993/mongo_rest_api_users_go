package app

import (
	"encoding/json"
	"mongo_rest_api_users/domain"
	"mongo_rest_api_users/service"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlersStruct struct {
	service service.UserRepository
}

func (s *HandlersStruct) GetAll(w http.ResponseWriter, r *http.Request) {
	result, er1 := s.service.GetAll()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ch *HandlersStruct) GetBuyId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	result, err := ch.service.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ch *HandlersStruct) Insert(w http.ResponseWriter, r *http.Request) {
	var user domain.UserStruct
	er1 := json.NewDecoder(r.Body).Decode(&user)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, er2 := ch.service.Insert(&user)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ch *HandlersStruct) Update(w http.ResponseWriter, r *http.Request) {
	var user domain.UserStruct
	er1 := json.NewDecoder(r.Body).Decode(&user)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, er2 := ch.service.Update(&user)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (ch *HandlersStruct) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	result, err := ch.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
