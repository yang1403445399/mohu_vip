package controller

import (
	"example/server/config"
	"example/server/model"
	"example/server/types"

	"github.com/dromara/carbon/v2"
	"github.com/gofiber/fiber/v2"
)

func (con CommonController) ArticleCount(c *fiber.Ctx) error {
	now := carbon.Now()

	articleGroupCount := []types.ArticleCroupCount{}
	articleCountData := types.ArticleCountData{}

	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	if startTime == "" || endTime == "" {
		startTime = now.Copy().SubMonths(11).StartOfMonth().Format("Y-m-d H:i:s")
		endTime = now.EndOfMonth().Format("Y-m-d H:i:s")
	}

	startDate := carbon.Parse(startTime)
	endDate := carbon.Parse(endTime)

	monthsDiff := startDate.DiffInMonths(endDate) + 1

	articleCountData.Label = make([]string, monthsDiff)
	articleCountData.Value = make([]uint, monthsDiff)

	for i := range monthsDiff {
		articleCountData.Label[i] = startDate.Copy().AddMonths(int(i)).Format("Y-m")
	}

	if err := config.DB.Model(&model.Article{}).
		Select("DATE_FORMAT(create_at, '%Y-%m') as label, COUNT(*) as value").
		Where("state = ?", 1). // 只统计已发布状态的文章
		Where("create_at BETWEEN ? AND ?", startTime, endTime).
		Group("DATE_FORMAT(create_at, '%Y-%m')").
		Find(&articleGroupCount).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	articleGroupCountMap := make(map[string]uint)
	for _, item := range articleGroupCount {
		articleGroupCountMap[item.Label] = item.Value
	}

	for i, item := range articleCountData.Label {
		articleCountData.Value[i] = articleGroupCountMap[item]
	}

	return con.SuccessResponse(c, articleCountData, "获取成功")
}
