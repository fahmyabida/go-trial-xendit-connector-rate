package usecase

import (
	http_client "connector-rate/http/client"
	"connector-rate/model"
	"connector-rate/redis_jobs"
	"connector-rate/repo"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type GetRateUsecase struct {
	httpClient        http_client.HttpClient
	redisClient       redis_jobs.RedisClient
	logRateChangeRepo repo.LogRateChangeRepo
}

func NewGetRateUsecase(httpClient http_client.HttpClient, redisClient redis_jobs.RedisClient,
	logRateChangeRepo repo.LogRateChangeRepo) GetRateUsecase {
	return GetRateUsecase{httpClient, redisClient, logRateChangeRepo}
}

func (u GetRateUsecase) GetRate(traceId, origin, destination string) (rateData model.RateData) {
	rateData.OriginCurrency = origin
	rateData.DestinationCurrency = destination
	rateData.Rate, rateData.ExpiredAt = u.redisClient.GetCurrentRate(traceId, origin, destination)
	if rateData.Rate > 0 {
		rateData.Success = true
		return rateData
	}
	response := u.httpClient.GetForex(traceId, origin, destination)
	if !response.Success {
		return rateData
	}
	rateInByteJSON, _ := json.Marshal(response.Rates)
	var mapCurrency = make(map[string]float64)
	json.Unmarshal(rateInByteJSON, &mapCurrency)
	rateOrigin, originOk := mapCurrency[origin]
	rateDestination, destiantionOk := mapCurrency[destination]
	if !originOk || !destiantionOk {
		rateData.Message = "currency not valid"
		return rateData
	}
	rateData.Rate = u.CountRate(rateOrigin, rateDestination)
	var err error
	rateData.ExpiredAt, err = time.ParseInLocation("2006-01-02", response.Date, time.UTC)
	if err != nil {
		log.Err(fmt.Errorf("%v, %v", traceId, err))
		return rateData
	}
	rateData.ExpiredAt = rateData.ExpiredAt.AddDate(0, 0, 1)
	err = u.redisClient.SetCurrentRate(traceId, origin, destination, rateData.Rate, rateData.ExpiredAt)
	if err != nil {
		log.Err(fmt.Errorf("%v, %v", traceId, err))
		return rateData
	}
	err = u.logRateChangeRepo.InsertLogRateChangeRepo(traceId, origin, destination, rateData.Rate, rateData.ExpiredAt)
	if err != nil {
		log.Err(fmt.Errorf("%v, %v", traceId, err))
		return rateData
	}
	rateData.Success = true
	return rateData
}

func (u GetRateUsecase) CountRate(originRate, destinationRate float64) (actualRate float64) {
	return (1 / originRate) * destinationRate
}
