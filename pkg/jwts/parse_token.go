package jwts

import (
	"GameShopPlatformServer/models"
	cerr "GameShopPlatformServer/pkg/custom-errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (*models.Claims, error) {
	if token == "" {
		return nil, cerr.ErrNoTokenFound
	}

	tokenString := strings.Split(token, " ")[1]

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	tkn, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !tkn.Valid {
		return nil, cerr.ErrInvalidToken
	}

	claims, ok := tkn.Claims.(*models.Claims)
	if !ok {
		return nil, cerr.ErrClaims
	}

	return claims, nil
}
