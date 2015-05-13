package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/couchbaselabs/gocb"
	"github.com/daishisystems/payment/todo"
	"github.com/gorilla/mux"
)

type Card struct {
	Number      string    `json:"number"`
	ExpiryMonth time.Time `json:"exp_month"`
	ExpiryYear  time.Time `json:"exp_year"`
}

type Metadata struct {
	PaymentType string `json:"payment_type"`
}

type PaymentCard struct {
	Id       string    `json:"id"`
	Created  time.Time `json:"created"`
	Card     Card      `json:"card"`
	Metadata Metadata  `json:"metadata"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	cluster, _ := gocb.Connect("couchbase://52.17.36.153")
	bucket, _ := cluster.OpenBucket("ryanair", "")

	var paymentCard PaymentCard
	bucket.Get("01b05e21-fc3a-4049-831a-14c8a2ef667d", &paymentCard)

	fmt.Printf("%s is %s\n", paymentCard.Id, paymentCard.Card.Number)
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
