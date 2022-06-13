package model

import "time"

type LogRateChange struct {
	tableName           struct{}  `pg:"public.log_rate_change"`
	TraceId             string    `pg:"trace_id"`
	OriginCurrency      string    `pg:"origin_currency"`
	DestinationCurrency string    `pg:"destination_currency"`
	Rate                float64   `pg:"rate"`
	ExpiredAt           time.Time `pg:"expired_at"`
	CreatedAt           time.Time `pg:"created_at,default:now()"`
}
