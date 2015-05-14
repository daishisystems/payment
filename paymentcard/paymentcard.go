package paymentcard

import "time"

type Card struct {
	Number      string `json:"number"`
	ExpiryMonth uint8  `json:"exp_month"`
	ExpiryYear  uint16 `json:"exp_year"`
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
