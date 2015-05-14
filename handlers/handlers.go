package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/couchbaselabs/gocb"
	"github.com/daishisystems/payment/couchbase"
	"github.com/gorilla/mux"
)

func Init(b *gocb.Bucket) {

	couchbase.Init(b)
}

func GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	pc := couchbase.GetById(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&pc)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	cards := couchbase.GetAll()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&cards)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
