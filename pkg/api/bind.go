package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) bind(w http.ResponseWriter, r *http.Request) {
	instanceID := mux.Vars(r)["instance_id"]
	log.Println(instanceID)
	bindingID := mux.Vars(r)["biding_id"]
	log.Println(bindingID)
	// TODO: Kick off synchronous binding
	// TODO: There are actually a lot of different response codes required for
	// different circumstances; also see spec for response body format
	w.WriteHeader(http.StatusCreated)
}