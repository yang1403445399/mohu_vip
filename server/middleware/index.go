package middleware

import (
	"fmt"
	"strings"

	"example/server/utils"

	"github.com/gofiber/fiber/v2"
)

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 跳过登录接口的验证
		if c.Path() == "/admin/user/login/submit" {
			return c.Next()
		}

		// 从请求头中获取Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.JSON(fiber.Map{
				"code": 403,
				"msg":  "未提供认证信息",
			})
		}

		// 移除所有可能的Bearer前缀
		tokenParts := strings.Split(authHeader, " ")
		var tokenString string

		// 处理不同的token格式
		if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
			tokenString = tokenParts[1]
		}

		// 解析token
		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || (token != nil && !token.Valid) {
			// 添加详细错误信息用于调试
			errorMsg := "无效的或过期的token"
			if err != nil {
				errorMsg = fmt.Sprintf("token验证失败: %v", err)
			}
			return c.JSON(fiber.Map{
				"code": 403,
				"msg":  errorMsg,
			})
		}

		// 将用户信息存储在上下文中，供后续处理函数使用
		c.Locals("userID", uint((*claims)["user_id"].(float64)))

		// 继续处理请求
		return c.Next()
	}
}
