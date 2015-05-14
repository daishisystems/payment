package couchbase

import (
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
