package controller

import (
	"example/server/config"
	"example/server/model"

	"github.com/gofiber/fiber/v2"
)

func (con CommonController) LogList(c *fiber.Ctx) error {
	total := int64(0)
	size := c.QueryInt("size", 10)
	current := c.QueryInt("current", 1)
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	if size <= 0 || size > 100 {
		size = 10
	}
	if current <= 0 {
		current = 1
	}

	offset := (current - 1) * size

	logData := []model.Log{}

	query := config.DB.Model(&model.Log{}).Preload("User")

	if startTime != "" && endTime != "" {
		query = query.Where("create_at BETWEEN ? AND ?", startTime, endTime)
	}

	if err := query.Order("id desc").Limit(size).Offset(offset).Find(&logData).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, fiber.Map{
		"list":    logData,
		"total":   total,
		"size":    size,
		"current": current,
	}, "获取成功")
}
