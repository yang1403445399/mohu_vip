package controller

import (
	"example/server/config"
	"example/server/model"

	"github.com/gofiber/fiber/v2"
)

func (con CommonController) BasicInfo(c *fiber.Ctx) error {
	basicData := []model.Basic{}

	if err := config.DB.Model(&model.Basic{}).Find(&basicData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, basicData, "获取成功")
}

func (con CommonController) BasicSubmit(c *fiber.Ctx) error {
	basicBody := []model.Basic{}

	if err := c.BodyParser(&basicBody); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	if err := config.DB.Model(&model.Basic{}).Save(&basicBody).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, nil, "提交成功")
}
