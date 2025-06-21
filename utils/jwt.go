package utils

import (
	"os"
	"strconv"
	"time"
	"vcm-medical-platform/models"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	expiresHours := 72 // Default 3 days
	if hours := os.Getenv("JWT_EXPIRES_HOURS"); hours != "" {
		if h, err := strconv.Atoi(hours); err == nil {
			expiresHours = h
		}
	}

	expirationTime := time.Now().Add(time.Duration(expiresHours) * time.Hour)
	
	claims := &Claims{
		UserID:   user.CdUser,
		Email:    user.Email,
		UserType: user.TyUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "vcm-medical-platform",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-change-in-production"
	}
	
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-change-in-production"
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
