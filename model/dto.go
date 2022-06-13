package model

import "time"

type RequestGetForex struct {
	OriginCurrency      string `json:"origin_currency"`
	DestinationCurrency string `json:"destination_currency"`
}

type ResponseGetForex3rdParty struct {
	Success   bool        `json:"success"`
	Timestamp int         `json:"timestamp"`
	Base      string      `json:"base"`
	Date      string      `json:"date"`
	Rates     interface{} `json:"rates"`
}

type RateData struct {
	Success             bool      `json:"success"`
	Message             string    `json:"message"`
	OriginCurrency      string    `json:"origin_currency"`
	DestinationCurrency string    `json:"destination_currency"`
	Rate                float64   `json:"rate"`
	ExpiredAt           time.Time `json:"expired_at"`
}
