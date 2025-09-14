package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = "my-secret-key"

// GenerateToken 生成JWT token
func GenerateToken(userID uint) (string, error) {
	// 设置token过期时间，这里设置为7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 创建claims
	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// 创建token对象，使用HS256算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的编码后的字符串token
	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("意外的签名方法")
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("token解析失败: %w", err)
	}

	// 提取claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("无效的token: claims类型错误")
	}

	// 创建指针类型用于返回
	claimsPtr := &claims

	if !token.Valid {
		return nil, nil, errors.New("无效的token: 验证失败")
	}

	return token, claimsPtr, nil
}

// RefreshToken 刷新JWT token
func RefreshToken(tokenString string) (string, error) {
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return "", err
	}

	// 检查token是否即将过期（例如，剩余时间少于30分钟）
	if exp, ok := (*claims)["exp"].(float64); ok {
		// 计算剩余时间
		remainingTime := time.Until(time.Unix(int64(exp), 0))
		if remainingTime > 30*time.Minute {
			// 不需要刷新
			return tokenString, nil
		}

		// 生成新token
		userID := uint((*claims)["user_id"].(float64))

		return GenerateToken(userID)
	}

	return "", errors.New("无效的token")
}
