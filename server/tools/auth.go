package tools

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(claim Claims, method jwt.SigningMethod, jwtKey []byte) (string, error) {
	// 使用jwt.NewWithClaims创建jwt token的header包含算法与类型的头与body可解码的内容体
	// 使用SignedString来创建签名, 构成完整的jwt
	token, err := jwt.NewWithClaims(method, claim).SignedString(jwtKey)

	return token, err
}

func ParseToken(tokenString string, claims jwt.Claims, encryptToken []byte) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return encryptToken, nil
	})

	if err != nil {
		return nil, err
	}

	if token != nil {
		if token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
