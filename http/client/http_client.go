package client

import (
	"connector-rate/model"
	"connector-rate/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type HttpClient struct {
	svcConfig model.SvcConfig
}

func NewHttpClient(svcConfig model.SvcConfig) HttpClient {
	return HttpClient{svcConfig}
}

func (c HttpClient) GetForex(traceId, originCurrency, destinationCurrency string) (response model.ResponseGetForex3rdParty) {
	url := strings.Replace(c.svcConfig.URLGetRate, "{access_key}", c.svcConfig.AccessKeyExchangeRates, 1)
	url = strings.Replace(url, "{origin}", originCurrency, 1)
	url = strings.Replace(url, "{destination}", destinationCurrency, 1)
	method := c.svcConfig.URLGetRateMethod

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Err(fmt.Errorf("'%v', %v", traceId, err))
		return
	}
	utils.LogOUT(traceId, method, url)
	var bodyResponse []byte = []byte("empty body response")
	defer func() {
		utils.LogIN(traceId, method, url, string(bodyResponse))
	}()
	res, err := client.Do(req)
	if err != nil {
		log.Err(fmt.Errorf("'%v', %v", traceId, err))
		return
	}
	defer res.Body.Close()

	bodyResponse, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Err(fmt.Errorf("'%v', %v", traceId, err))
		return
	}
	json.Unmarshal(bodyResponse, &response)
	return response
}
