package jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(sub string, exp time.Duration, secretString string) (string, error) {
	secretKey := []byte(secretString)
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", NewJwtError(fmt.Errorf("signing failed :%w", err))
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string, secretString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, NewJwtError(fmt.Errorf("invalid signing method"))
			}
			return secretString, nil
		},
	)
	if err != nil {
		return nil, NewJwtError(fmt.Errorf("token parsing failed: %w", err))
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, NewJwtError(fmt.Errorf("invalid token"))
}
