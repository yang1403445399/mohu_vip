package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type UseCommonController struct {
	CommonController // 嵌入基础控制器
}

func (con UseCommonController) getIPLocation(c *fiber.Ctx) error {
	agent := fiber.Get("https://restapi.amap.com/v3/ip?key=13f18176eb98b6757c452a64743d3c99&ip=111.178.57.181")

	_, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return con.ErrorResponse(c, errs[0].Error())
	}

	ipLocationData := fiber.Map{}
	err := json.Unmarshal(body, &ipLocationData)
	if err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, ipLocationData, "获取成功")
}
