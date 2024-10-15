package jwts

import (
	"GameShopPlatformServer/models"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string, roles []string) (string, error) {
	claims := &models.Claims{
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			Subject: id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}
