package controller

import (
	"example/server/config"
	"example/server/model"

	"github.com/gofiber/fiber/v2"
)

func (con CommonController) MenuList(c *fiber.Ctx) error {
	menuData := []model.Menu{}

	if err := config.DB.Model(&model.Menu{}).Find(&menuData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, menuData, "获取成功")
}
