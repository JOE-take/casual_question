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

// GenerateAccessToken 指定されたユーザ情報でトークンを生成する生成器
func GenerateAccessToken(u *models.User) (string, error) {

	claims := Claims{
		u.UserID,
		u.UserName,
		u.Email,
		jwt.StandardClaims{
			Issuer:    "cq-api",
			Subject:   "AccessToken",
			Audience:  "cq-front",
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        uuid.NewString(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

// ParseAccessToken アクセストークンを解析してClaimsとエラーを返す解析器
func ParseAccessToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GenerateRefreshToken() (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    "cp-api",
		Subject:   "RefreshToken",
		Audience:  "cp-front",
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        uuid.NewString(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

func ValidateRefreshToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if token != nil {
		_, ok := token.Claims.(*jwt.StandardClaims)
		if ok && token.Valid {
			return true, nil
		}
	}

	return false, err
}
