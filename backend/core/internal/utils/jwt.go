package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	AccessSecret  []byte
	RefreshSecret []byte
)

func InitSecrets() {
	a := os.Getenv("ACCESS_SECRET")
	r := os.Getenv("REFRESH_SECRET")
	if a == "" || r == "" {
		panic("ACCESS_SECRET and REFRESH_SECRET must be set")
	}
	AccessSecret = []byte(a)
	RefreshSecret = []byte(r)
}

// ------------------ ACCESS TOKEN ------------------

func CreateAccessToken(userID uint) (string, int64, error) {
	exp := time.Now().Add(15 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(AccessSecret)
	if err != nil {
		return "", 0, err
	}
	return signed, exp, nil
}

func ValidateAccessTokenWithExpiry(tokenStr string) (uint, int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return AccessSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, 0, errors.New("invalid or expired access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, 0, errors.New("invalid access token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, 0, errors.New("invalid user_id in access token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, 0, errors.New("invalid exp in access token")
	}

	return uint(userID), int64(exp), nil
}

// ------------------ REFRESH TOKEN ------------------

func CreateRefreshToken(userID uint) (string, int64, error) {
	exp := time.Now().Add(7 * 24 * time.Hour).Unix() // 7 days
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(RefreshSecret)
	if err != nil {
		return "", 0, err
	}
	return signed, exp, nil
}

func ValidateRefreshTokenWithExpiry(tokenStr string) (uint, int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, 0, errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, 0, errors.New("invalid refresh token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, 0, errors.New("invalid user_id in refresh token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, 0, errors.New("invalid exp in refresh token")
	}

	return uint(userID), int64(exp), nil
}
