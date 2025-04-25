package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var jwtKey = []byte("your-secret-key") // đặt ở .env trong thực tế

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// Tạo token khi login thành công
func GenerateToken(userId uint) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Xác minh token từ client gửi lên
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}

// ParseTokenFromHeader nhận HTTP request và trả về userID
func ParseTokenFromHeader(authHeader string) (uint, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return 0, errors.New("invalid auth header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	return claims.UserId, nil
}
