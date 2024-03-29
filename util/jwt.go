package util

import (
	"casual_question/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
)

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

	jwtSecret := os.Getenv("SECRET_KEY")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

// ParseAccessToken アクセストークンを解析してClaimsとエラーを返す解析器
func ParseAccessToken(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("SECRET_KEY")
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

// GenerateRefreshToken () (tokenString string, expiry, err error)
// expiry はUnix時間
func GenerateRefreshToken() (string, int64, error) {
	exp := time.Now().Add(48 * time.Hour).Unix()
	claims := jwt.StandardClaims{
		Issuer:    "cp-api",
		Subject:   "RefreshToken",
		Audience:  "cp-front",
		ExpiresAt: exp,
		IssuedAt:  time.Now().Unix(),
		Id:        uuid.NewString(),
	}

	jwtSecret := os.Getenv("SECRET_KEY")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, exp, err
}

func ValidateRefreshToken(tokenString string) (bool, error) {
	jwtSecret := os.Getenv("SECRET_KEY")
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
