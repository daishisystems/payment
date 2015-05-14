package couchbase

import (
	"fmt"
	"github.com/couchbaselabs/gocb"
	"github.com/daishisystems/payment/paymentcard"
)

var bucket gocb.Bucket

func Init(b *gocb.Bucket) {
	bucket = *b
}

func GetById(id string) *paymentcard.PaymentCard {

	var pc paymentcard.PaymentCard
	bucket.Get(id, &pc)

	return &pc
}

type viewRow struct {
	Id    string
	Key   string
	Value paymentcard.PaymentCard
}

func GetAll() *[]paymentcard.PaymentCard {

	vq := gocb.NewViewQuery("getallfull", "getallfull")
	rows := bucket.ExecuteViewQuery(vq)

	var cards []paymentcard.PaymentCard

	row := viewRow{}
	for rows.Next(&row) {
		cards = append(cards, row.Value)
	}
	if err := rows.Close(); err != nil {
		fmt.Printf("View Query Error: %s", err)
	}

	return &cards
}
