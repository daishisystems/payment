package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/couchbaselabs/gocb"
	"github.com/daishisystems/payment/couchbase"
	"github.com/daishisystems/payment/todo"
	"github.com/gorilla/mux"
)

func Init(b *gocb.Bucket) {

	couchbase.Init(b)
}

func Index(w http.ResponseWriter, r *http.Request) {

	pc := couchbase.GetById("01b05e21-fc3a-4049-831a-14c8a2ef667d")
	fmt.Printf("%s is %s\n", pc.Id, pc.Card.Number)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := todo.Todos{
		todo.Todo{Name: "Write presentation"},
		todo.Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
