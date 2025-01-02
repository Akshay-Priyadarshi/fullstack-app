package jwts

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(sub string, exp time.Duration, secretString string) (string, error) {
	secret := []byte(secretString)
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
		"iss": "fullstack.app",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", ErrSigningFailed
	}
	return signedToken, nil
}

func VerifyJWT(signedToken string, secretString string) (jwt.MapClaims, error) {
	secret := []byte(secretString)
	token, err := jwt.Parse(
		signedToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidSigningMethod
			}
			return secret, nil
		},
	)
	if err != nil {
		return nil, ErrTokenParsingFailed
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
