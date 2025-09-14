package controller

import (
	"crypto/md5"
	"encoding/hex"
	"example/server/config"
	"example/server/model"
	"example/server/types"
	"example/server/utils"

	"github.com/gofiber/fiber/v2"
)

func (con CommonController) UserLoginSubmit(c *fiber.Ctx) error {
	userLoginBody := types.UserLoginBody{}

	if err := c.BodyParser(&userLoginBody); err != nil {
		return con.ErrorResponse(c, "解析请求数据失败: "+err.Error())
	}

	// 判断登录方式：账号密码登录
	if userLoginBody.Username != nil && userLoginBody.Password != nil {
		if *userLoginBody.Username == "" || *userLoginBody.Password == "" {
			return con.ErrorResponse(c, "用户名或密码不能为空")
		}

		md5Password := md5.Sum([]byte(*userLoginBody.Password))

		userData := model.User{}

		if err := config.DB.Model(&model.User{}).Where("username = ? AND password = ?", *userLoginBody.Username, hex.EncodeToString(md5Password[:])).Take(&userData).Error; err != nil {
			return con.ErrorResponse(c, "用户名或密码错误")
		}

		token, err := utils.GenerateToken(userData.Id)
		if err != nil {
			return con.ErrorResponse(c, "生成token失败: "+err.Error())
		}

		userLoginData := types.UserLoginData{
			Id:       userData.Id,
			RoleId:   userData.RoleId,
			Username: userData.Username,
			Nickname: userData.Nickname,
			Avatar:   userData.Avatar,
			Mobile:   userData.Mobile,
			Uuid:     userData.Uuid,
			CreateAt: userData.CreateAt,
			UpdateAt: userData.UpdateAt,
			State:    userData.State,
			Token:    "Bearer " + token,
		}

		return con.SuccessResponse(c, userLoginData, "登录成功")
	}

	// 判断登录方式：手机号验证码登录
	// if userLoginBody.Mobile != nil && userLoginBody.Code != nil {

	// }

	return con.ErrorResponse(c, "请输入账号密码或手机号验证码进行登录")
}

func (con CommonController) UserLoginVerify(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return con.ErrorResponse(c, "用户ID不存在")
	}

	userData := model.User{}

	if err := config.DB.Model(&model.User{}).Where("id = ?", userID).Take(&userData).Error; err != nil {
		return con.ErrorResponse(c, "用户不存在")
	}

	return con.SuccessResponse(c, userData, "获取成功")
}
