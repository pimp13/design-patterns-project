package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pimp13/go-react-project/config"
	"time"
)

func CreateJWT(secret []byte, userID uuid.UUID) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSecond)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
