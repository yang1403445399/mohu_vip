package controller

import (
	"github.com/gofiber/fiber/v2"
)

type CommonController struct {
	// 可以在这里添加共享的字段
}

// SuccessResponse 返回成功响应，自动处理data为nil的情况
func (c *CommonController) SuccessResponse(ctx *fiber.Ctx, data any, msg string) error {
	responseData := fiber.Map{
		"code": 200,
		"msg":  msg,
	}

	if data != nil {
		responseData["data"] = data
	}

	return ctx.JSON(responseData)
}

// ErrorResponse 返回错误响应
func (c *CommonController) ErrorResponse(ctx *fiber.Ctx, errMsg string) error {
	return ctx.JSON(fiber.Map{
		"code": 400,
		"msg":  errMsg,
	})
}
