package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type LocationController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
}

func Setup(lc LocationController) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/locations", lc.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/locations/{woeid}", lc.GetOne).Methods(http.MethodGet)

	return r
}
