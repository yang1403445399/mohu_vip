package controller

import (
	"example/server/config"
	"example/server/model"
	"example/server/types"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/gofiber/fiber/v2"
)

func (con CommonController) ImageList(c *fiber.Ctx) error {
	imageData := []model.Image{}
	total := int64(0)
	size := c.QueryInt("size", 12)
	current := c.QueryInt("current", 1)
	columnId := c.QueryInt("column_id", 1)

	if size <= 0 || size > 100 {
		size = 12
	}

	if current <= 0 {
		current = 1
	}

	offset := (current - 1) * size

	if err := config.DB.Model(&model.Image{}).Where("column_id = ?", columnId).Order("id desc").Limit(size).Offset(offset).Find(&imageData).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, fiber.Map{
		"list":    imageData,
		"total":   total,
		"size":    size,
		"current": current,
	}, "获取成功")
}

func (con CommonController) ImageUpload(c *fiber.Ctx) error {
	image, err := c.FormFile("image")
	if err != nil {
		return con.ErrorResponse(c, "图片错误："+err.Error())
	}

	columnId, err := strconv.Atoi(c.FormValue("column_id"))
	if err != nil {
		return con.ErrorResponse(c, "栏目错误："+err.Error())
	}

	imageExt := strings.ToLower(filepath.Ext(image.Filename))
	imageTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	if !imageTypes[imageExt] {
		return con.ErrorResponse(c, "只允许上传jpg、jpeg、png、gif、webp格式的图片")
	}
	if image.Size > 2*1024*1024 {
		return con.ErrorResponse(c, "文件大小不能超过2MB")
	}

	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return con.ErrorResponse(c, "用户身份验证失败")
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return con.ErrorResponse(c, "创建上传目录失败："+err.Error())
	}

	fullDir := fmt.Sprintf("%s/%s", uploadDir, carbon.Now().Format("Ymd"))
	if err := os.MkdirAll(fullDir, os.ModePerm); err != nil {
		return con.ErrorResponse(c, "创建上传目录失败："+err.Error())
	}

	imageName := fmt.Sprintf("%d_%s", carbon.Now().Timestamp(), image.Filename)
	savePath := fmt.Sprintf("%s/%s", fullDir, imageName)
	if err := c.SaveFile(image, savePath); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	imageUrl := fmt.Sprintf("%s/uploads/%s/%s", c.BaseURL(), carbon.Now().Format("Ymd"), imageName)

	imageModel := model.Image{
		UserId:   userID,
		ColumnId: uint(columnId),
		Name:     imageName,
		Url:      imageUrl,
	}

	if err := config.DB.Create(&imageModel).Error; err != nil {
		return con.ErrorResponse(c, "上传图片失败："+err.Error())
	}

	return con.SuccessResponse(c, nil, "上传成功")
}

func (con CommonController) ImageDelete(c *fiber.Ctx) error {
	imageBody := []uint{}

	if err := c.BodyParser(&imageBody); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	if len(imageBody) == 0 {
		return con.ErrorResponse(c, "请选择要删除的图片")
	}

	var imageData []model.Image
	if err := config.DB.Model(&model.Image{}).Where("id IN ?", imageBody).Find(&imageData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	for _, image := range imageData {
		imagePath := filepath.Join(".", "uploads", image.CreateAt.Format("Ymd"), image.Name)

		absPath, absErr := filepath.Abs(imagePath)
		if absErr != nil {
			continue
		}

		err := os.Remove(absPath)
		if err != nil {
			go func(path string) {
				const maxRetries = 5

				for i := range make([]int, maxRetries) {
					waitTime := time.Duration(1<<uint(i)) * time.Second
					randomFactor := time.Duration(rand.Intn(500)) * time.Millisecond
					totalWait := waitTime + randomFactor

					time.Sleep(totalWait)

					retryErr := os.Remove(path)
					if retryErr == nil {
						return
					}
				}
			}(absPath)
		}
	}

	if err := config.DB.Model(&model.Image{}).Delete(&model.Image{}, imageBody).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, nil, "删除成功")
}

func (con CommonController) ImageColumnList(c *fiber.Ctx) error {
	imageColumnData := []model.ImageColumn{}

	if err := config.DB.Model(&model.ImageColumn{}).Find(&imageColumnData).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, imageColumnData, "获取成功")
}

func (con CommonController) ImageColumnSubmit(c *fiber.Ctx) error {
	imageColumnBody := types.ImageColumnBody{}

	if err := c.BodyParser(&imageColumnBody); err != nil {
		return con.ErrorResponse(c, "参数格式错误")
	}

	if imageColumnBody.Name == "" {
		return con.ErrorResponse(c, "请输入栏目名称")
	}

	if len(imageColumnBody.Name) > 50 {
		return con.ErrorResponse(c, "栏目名称不能超过50个字符")
	}

	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return con.ErrorResponse(c, "用户身份验证失败")
	}

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if imageColumnBody.Id != nil {
		imageColumnModel := model.ImageColumn{}
		imageColumnCount := int64(0)

		if err := tx.Where("id = ? AND user_id = ?", *imageColumnBody.Id, userID).First(&imageColumnModel).Error; err != nil {
			tx.Rollback()
			return con.ErrorResponse(c, "更新失败："+err.Error())
		}

		tx.Model(&model.ImageColumn{}).Where("name = ? AND user_id = ? AND id != ?", imageColumnBody.Name, userID, *imageColumnBody.Id).Count(&imageColumnCount)
		if imageColumnCount > 0 {
			tx.Rollback()
			return con.ErrorResponse(c, "该栏目名称已存在")
		}

		if err := tx.Model(&model.ImageColumn{}).Where("id = ?", *imageColumnBody.Id).Updates(imageColumnBody).Error; err != nil {
			tx.Rollback()
			return con.ErrorResponse(c, "更新失败："+err.Error())
		}

		tx.Commit()

		return con.SuccessResponse(c, nil, "更新成功")
	}

	imageColumnModel := model.ImageColumn{
		UserId: userID,
		Name:   imageColumnBody.Name,
	}
	imageColumnCount := int64(0)

	tx.Model(&model.ImageColumn{}).Where("name = ? AND user_id = ?", imageColumnBody.Name, userID).Count(&imageColumnCount)
	if imageColumnCount > 0 {
		tx.Rollback()
		return con.ErrorResponse(c, "该栏目名称已存在")
	}

	if err := tx.Model(&model.ImageColumn{}).Create(&imageColumnModel).Error; err != nil {
		tx.Rollback()
		return con.ErrorResponse(c, "新增失败："+err.Error())
	}

	tx.Commit()

	return con.SuccessResponse(c, nil, "新增成功")
}

func (con CommonController) ImageColumnDelete(c *fiber.Ctx) error {
	imageColumnBody := []uint{}

	if err := c.BodyParser(&imageColumnBody); err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	if len(imageColumnBody) == 0 {
		return con.ErrorResponse(c, "请选择要删除的栏目")
	}

	if err := config.DB.Delete(&model.ImageColumn{}, imageColumnBody).Error; err != nil {
		return con.ErrorResponse(c, err.Error())
	}

	return con.SuccessResponse(c, nil, "删除成功")
}
