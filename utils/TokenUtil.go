package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("deng")

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	// 设置有效时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "AustinDeng",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}


func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
	})

	if tokenClaims != nil {
			if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
					return claims, nil
			}
	}

	return nil, err
}