package usecase

import (
	"connector-rate/http/client"
	"connector-rate/redis_jobs"
	"connector-rate/repo"
	"testing"
)

func TestCountRate(t *testing.T) {
	getRateUsecsae := NewGetRateUsecase(client.HttpClient{}, redis_jobs.RedisClient{}, repo.LogRateChangeRepo{})
	expectResult := 14443.195327493726
	actualResult := getRateUsecsae.CountRate(1.056692, 15262.008957)
	t.Log(actualResult)
	if actualResult != expectResult {
		t.Error("hasil belum sesuai expectasi")
	}
	t.Log("OK!")
}
