package main

import (
	http_client "connector-rate/http/client"
	http_server "connector-rate/http/server"
	"connector-rate/redis_jobs"
	"connector-rate/repo"
	"connector-rate/usecase"
	"fmt"
)

func RunApplication() {
	svcConfig := GetConfig()
	fmt.Println(svcConfig)

	redisConn := redisConnect(svcConfig)
	dbConn := databaseConnect(svcConfig)

	httpClient := http_client.NewHttpClient(svcConfig)
	redisClient := redis_jobs.NewRedisClient(redisConn, svcConfig)
	logRateChangeRepo := repo.NewLogRateChangeRepo(dbConn)
	getRateUsecase := usecase.NewGetRateUsecase(httpClient, redisClient, logRateChangeRepo)
	httpServer := http_server.NewServerHttp(getRateUsecase)
	httpServer.Init()
}
