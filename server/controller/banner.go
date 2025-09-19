package controller

import (
	"example/server/config"
	"example/server/model"
	"example/server/types"
	"example/server/utils"

	"github.com/gofiber/fiber/v2"
)

func (con CommonController) BannerList(c *fiber.Ctx) error {
	bannerData := []model.Banner{}
	total := int64(0)
	size := c.QueryInt("size", 10)
	current := c.QueryInt("current", 1)
	keyword := c.Query("keyword", "")

	if size <= 0 || size > 100 {
		size = 10
	}
	if current <= 0 {
		current = 1
	}

	offset := (current - 1) * size

	if err := config.DB.Model(&model.Banner{}).Preload("Type").Where("name like ?", "%"+keyword+"%").Order("id desc").Limit(size).Offset(offset).Find(&bannerData).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, fiber.Map{
		"list":    bannerData,
		"total":   total,
		"size":    size,
		"current": current,
		"keyword": keyword,
	}, "获取成功")
}

func (con CommonController) BannerInfo(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	if id <= 0 {
		return con.ErrorResponse(c, "请选择要获取的信息")
	}

	bannerData := model.Banner{}

	if err := config.DB.Model(&model.Banner{}).Where("id = ?", id).First(&bannerData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, bannerData, "获取成功")
}

func (con CommonController) BannerSubmit(c *fiber.Ctx) error {
	bannerBody := types.BannerBody{}

	if err := c.BodyParser(&bannerBody); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	if bannerBody.TypeId == 0 {
		return con.ErrorResponse(c, "请选择类型")
	}

	if bannerBody.Name == "" {
		return con.ErrorResponse(c, "请输入名称")
	}

	if bannerBody.Src == "" {
		return con.ErrorResponse(c, "请选择图片")
	}

	if bannerBody.Id != nil {
		if err := config.DB.Model(&model.Banner{}).Where("id = ?", *bannerBody.Id).Updates(bannerBody).Error; err != nil {
			return con.ErrorResponse(c, "更新失败："+err.Error())
		}

		return con.SuccessResponse(c, nil, "更新成功")
	}

	bannerData := model.Banner{
		TypeId: bannerBody.TypeId,
		Name:   bannerBody.Name,
		Src:    bannerBody.Src,
		Url:    utils.GetDefaultBodyString(bannerBody.Url, ""),
		Sort:   utils.GetDefaultBodyUint(bannerBody.Sort, 0),
	}

	if err := config.DB.Create(&bannerData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, nil, "提交成功")
}

func (con CommonController) BannerDelete(c *fiber.Ctx) error {
	bannerBody := []uint{}

	if err := c.BodyParser(&bannerBody); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	if len(bannerBody) == 0 {
		return con.ErrorResponse(c, "请选择要删除的banner")
	}

	if err := config.DB.Delete(&model.Banner{}, bannerBody).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, nil, "删除成功")
}

func (con CommonController) BannerTypeList(c *fiber.Ctx) error {
	bannerTypeData := []model.BannerType{}

	if err := config.DB.Model(&model.BannerType{}).Find(&bannerTypeData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}
	return con.SuccessResponse(c, bannerTypeData, "获取成功")
}
