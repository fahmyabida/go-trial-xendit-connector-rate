package server

import (
	"connector-rate/model"
	"connector-rate/utils"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DefaultMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		traceId := string(ctx.Request().Header.Peek("trace_id"))
		if traceId == "" {
			traceId = uuid.New().String()
			ctx.Locals(model.HEADER_TRACE_ID, traceId)
		}

		request := strings.ReplaceAll(string(ctx.Request().Body()), "\n", "")
		request = strings.ReplaceAll(request, " ", "")
		utils.LogIN(traceId,
			"GATEWAY",
			string(ctx.Request().Header.Method())+" "+ctx.OriginalURL(),
			string(ctx.Request().Body()))
		err := ctx.Next()
		utils.LogOUT(traceId,
			"GATEWAY",
			string(ctx.Request().Header.Method())+" "+ctx.OriginalURL(),
			fmt.Sprintf("http_code:'%v', body_response:'%v'", ctx.Response().StatusCode(), string(ctx.Response().Body())))
		return err
	}
}
