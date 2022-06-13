package repo

import (
	"connector-rate/model"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
)

type LogRateChangeRepo struct {
	db *pg.DB
}

type ILogRateChangeRepo interface {
	InsertLogRateChangeRepo(traceId, origin, destination string, rate float64, expireAt time.Time) error
}

func NewLogRateChangeRepo(db *pg.DB) LogRateChangeRepo {
	return LogRateChangeRepo{db}
}

func (r LogRateChangeRepo) InsertLogRateChangeRepo(traceId, origin, destination string, rate float64, expireAt time.Time) error {
	data := model.LogRateChange{
		TraceId:             traceId,
		OriginCurrency:      origin,
		DestinationCurrency: destination,
		Rate:                rate,
		ExpiredAt:           expireAt,
		CreatedAt:           time.Time{},
	}
	if _, err := r.db.Model(&data).Insert(); err != nil {
		log.Err(fmt.Errorf("%v, %v", traceId, err))
		return err
	}
	return nil
}
