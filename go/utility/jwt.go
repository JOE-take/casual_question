package utility

import (
	"casual_question/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const jwtSecret = "secret_key?"

type Claims struct {
	UserID   string
	UserName string
	Email    string
	jwt.StandardClaims
}

// GenerateToken 指定されたユーザ情報でトークンを生成する生成器
func GenerateToken(u *models.User) (string, error) {

	claims := Claims{
		u.UserID,
		u.UserName,
		u.Email,
		jwt.StandardClaims{
			Issuer:    "cq-api",
			Subject:   "auth",
			Audience:  "cq-front",
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        uuid.NewString(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken トークンを解析してClaimsとエラーを返す解析器
func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
