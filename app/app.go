package app

import (
	"log"
	"mongo_rest_api_users/domain"
	"mongo_rest_api_users/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	DB := DBConnect("Users")
	router := mux.NewRouter()
	h := HandlersStruct{service.NewUserRepsitory(domain.NewDatabase(DB))}
	Route(router, &h)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
