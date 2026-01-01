package helpers

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encryptedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptedBytes), err
}

func VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJwt(userId int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userId),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		"iat": jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJwt(tokenString string, secret string) (int64, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return 0, errors.Wrap(err, "invalid token")
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims of token")
	}

	userIdString, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid sub claim of token")
	}
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "invalid sub claim of token")
	}

	return userId, nil
}
