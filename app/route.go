package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route(router *mux.Router, h *HandlersStruct) {
	router.HandleFunc("/users", h.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.GetBuyId).Methods(http.MethodGet)
	router.HandleFunc("/users", h.Insert).Methods(http.MethodPost)
	router.HandleFunc("/users", h.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", h.Delete).Methods(http.MethodDelete)

}
