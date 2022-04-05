package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const salt = "footprint"

func GenerateToken(userName string) (string, error) {
	claims := Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*4,
			Issuer:    "beianAPI",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSalt := []byte(salt)
	secretToken, err := tokenClaims.SignedString(jwtSalt)
	return secretToken, err
}

func ParseToken(secretToken string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(secretToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
