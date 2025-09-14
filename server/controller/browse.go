package controller

import (
	"example/server/config"
	"example/server/model"
	"example/server/types"
	"fmt"

	"github.com/dromara/carbon/v2"
	"github.com/gofiber/fiber/v2"
)

func (con CommonController) BrowseCount(c *fiber.Ctx) error {
	now := carbon.Now()

	queryParams := []types.BrowseQueryParams{
		{
			// 今日
			Name: "日访问量",
			Date: now.Format("Y年m月d日"),
			Current: types.BrowseTimeRange{
				Start: now.StartOfDay(),
				End:   now.EndOfDay(),
			},
			Previous: types.BrowseTimeRange{
				Start: now.Copy().StartOfDay().SubDays(1),
				End:   now.Copy().EndOfDay().SubDays(1),
			},
			SelectSQL:   "HOUR(create_at) as label, COUNT(*) as value",
			GroupBy:     "HOUR(create_at)",
			LabelFormat: func(i int) string { return fmt.Sprintf("%02d:00 ~ %02d:00", i, (i+1)%24) },
			Length:      24,
			Offset:      0,
		},
		{
			// 本周
			Name: "周访问量",
			Date: now.StartOfWeek().Format("Y年m月d日") + "~" + now.EndOfWeek().Format("Y年m月d日"),
			Current: types.BrowseTimeRange{
				Start: now.StartOfWeek(),
				End:   now.EndOfWeek(),
			},
			Previous: types.BrowseTimeRange{
				Start: now.Copy().StartOfWeek().SubDays(7),
				End:   now.Copy().EndOfWeek().SubDays(7),
			},
			SelectSQL:   "DAYOFWEEK(create_at) - 1 as label, COUNT(*) as value",
			GroupBy:     "DAYOFWEEK(create_at) - 1",
			LabelFormat: func(i int) string { return now.Copy().StartOfWeek().AddDays(i).Format("Y-m-d") },
			Length:      7,
			Offset:      1,
		},
		{
			// 本月
			Name: "月访问量",
			Date: now.Format("Y年m月"),
			Current: types.BrowseTimeRange{
				Start: now.StartOfMonth(),
				End:   now.EndOfMonth(),
			},
			Previous: types.BrowseTimeRange{
				Start: now.Copy().SubMonths(1).StartOfMonth(),
				End:   now.Copy().SubMonths(1).EndOfMonth(),
			},
			SelectSQL:   "DAY(create_at) as label, COUNT(*) as value",
			GroupBy:     "DAY(create_at)",
			LabelFormat: func(i int) string { return now.Copy().StartOfMonth().AddDays(i).Format("Y-m-d") },
			Length:      now.DaysInMonth(),
			Offset:      1,
		},
		{
			// 本年
			Name: "年访问量",
			Date: now.Format("Y年"),
			Current: types.BrowseTimeRange{
				Start: now.StartOfYear(),
				End:   now.EndOfYear(),
			},
			Previous: types.BrowseTimeRange{
				Start: now.Copy().SubYears(1).StartOfYear(),
				End:   now.Copy().SubYears(1).EndOfYear(),
			},
			SelectSQL:   "MONTH(create_at) as label, COUNT(*) as value",
			GroupBy:     "MONTH(create_at)",
			LabelFormat: func(i int) string { return now.Copy().StartOfYear().AddMonths(i).Format("Y-m") },
			Length:      12,
			Offset:      1,
		},
	}

	browseGroupCount := func(selectClause string, whereClause string, groupBy string, args ...any) (any, error) {
		groupCount := []types.BrowseCroupCount{}
		err := config.DB.Model(&model.Browse{}).Select(selectClause).Where(whereClause, args...).Where("state = ?", 1).Group(groupBy).Find(&groupCount).Error
		return groupCount, err
	}

	BrowseTrendCalculate := func(currentStart, currentEnd, previousStart, previousEnd *carbon.Carbon) string {
		currentCount, previousCount := int64(0), int64(0)

		if err := config.DB.Model(&model.Browse{}).Where("create_at BETWEEN ? AND ?", currentStart, currentEnd).Where("state = ?", 1).Count(&currentCount).Error; err != nil {
			return "0%"
		}

		if err := config.DB.Model(&model.Browse{}).Where("create_at BETWEEN ? AND ?", previousStart, previousEnd).Where("state = ?", 1).Count(&previousCount).Error; err != nil {
			return "0%"
		}

		if previousCount == 0 {
			if currentCount > 0 {
				return "+100%"
			}
			return "0%"
		}

		percentage := float64(currentCount-previousCount) / float64(previousCount) * 100
		return fmt.Sprintf("%.2f%%", percentage)
	}

	browseCountData := make([]types.BrowseCountData, len(queryParams))

	for i, tr := range queryParams {
		browseCountData[i].Name = tr.Name
		browseCountData[i].Date = tr.Date
		browseCountData[i].Label = make([]string, tr.Length)
		browseCountData[i].Value = make([]uint, tr.Length)
		browseCountData[i].Trend = BrowseTrendCalculate(tr.Current.Start, tr.Current.End, tr.Previous.Start, tr.Previous.End)

		for j := 0; j < tr.Length; j++ {
			browseCountData[i].Label[j] = tr.LabelFormat(j)
		}
	}

	for i, tr := range queryParams {
		groupCount, err := browseGroupCount(tr.SelectSQL, "create_at BETWEEN ? AND ?", tr.GroupBy, tr.Current.Start, tr.Current.End)
		if err != nil {
			continue
		}

		groupCountData, ok := groupCount.([]types.BrowseCroupCount)
		if !ok {
			continue
		}

		browseCountData[i].Group = make([]types.BrowseCroupCount, len(groupCountData))

		totalCount := uint(0)
		for key, val := range groupCountData {
			totalCount += val.Value

			valueIdx := int(val.Label) - tr.Offset
			if valueIdx >= 0 && valueIdx < len(browseCountData[i].Value) {
				browseCountData[i].Value[valueIdx] = val.Value
			}

			browseCountData[i].Group[key].Label = val.Label
			browseCountData[i].Group[key].Value = val.Value
		}
		browseCountData[i].Count = totalCount
	}

	return con.SuccessResponse(c, browseCountData, "获取成功")
}

func (con CommonController) BrowseRegion(c *fiber.Ctx) error {
	browseProvinceCount := []types.BrowseProvinceCount{}

	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	query := config.DB.Model(&model.Browse{}).Select("province, COUNT(*) as count").Where("state = ? AND province != ?", 1, "")

	if startTime != "" && endTime != "" {
		query = query.Where("create_at BETWEEN ? AND ?", startTime, endTime)
	}

	if err := query.Group("province").Find(&browseProvinceCount).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	browseRegionData := make([]types.BrowseRegionData, 0, len(browseProvinceCount))

	for _, item := range browseProvinceCount {
		browseRegionData = append(browseRegionData, types.BrowseRegionData{
			Rank:  0,
			Name:  item.Province,
			Value: item.Count,
		})
	}

	return con.SuccessResponse(c, browseRegionData, "获取成功")
}
