package server

import (
	"connector-rate/model"
	"connector-rate/usecase"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ServerHttp struct {
	getRateUsecase usecase.GetRateUsecase
}

func NewServerHttp(getRateUsecase usecase.GetRateUsecase) ServerHttp {
	return ServerHttp{getRateUsecase}
}

func (h *ServerHttp) Init() {
	fiberApp := fiber.New()
	fiberApp.Use(DefaultMiddleware())

	fiberApp.Post("v1/connector/rate-exchange", h.GetRate)

	fmt.Printf("HTTP listen on port :81 \n")
	fiberApp.Listen(fmt.Sprintf(":81"))
}

func (h *ServerHttp) GetRate(ctx *fiber.Ctx) error {
	traceId := fmt.Sprint(ctx.Locals(model.HEADER_TRACE_ID))
	var reqBody model.RequestGetForex
	json.Unmarshal(ctx.Request().Body(), &reqBody)
	response := h.getRateUsecase.GetRate(traceId, reqBody.OriginCurrency, reqBody.DestinationCurrency)
	if !response.Success {
		return ctx.Status(500).JSON(response)
	}
	return ctx.Status(200).JSON(response)
}
