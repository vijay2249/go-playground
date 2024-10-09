package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"wedonttrack.org/gorest/constants"
)

const SECRET_KEY = "Harry Styles"

func GenerateJWTToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.USER_EMAIL: email,
		constants.USER_ID: userId,
		"exp": time.Now().Add(time.Hour*2).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(interface{}, error) {
		_, isTure := token.Method.(*jwt.SigningMethodHMAC)
		if !isTure {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, isClaims := parsedToken.Claims.(jwt.MapClaims)
	if !isClaims{
		return 0, errors.New("invalid token claims")
	}
	// email, _ := claims[constants.USER_EMAIL].(string)
	userId := int64(claims[constants.USER_ID].(float64))
	return userId, nil
}