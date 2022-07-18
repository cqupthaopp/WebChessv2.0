package utils

/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}

const TokenTime = time.Hour * 10

var Secret = []byte("RedRockChess Author:Hao_pp")

const Author = "Hao_pp"

func GetToken(username string) (string, error) {

	c := MyClaims{

		username, "AccessToken", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			Issuer:    Author,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(Secret)

}

func GetUserFromToken(token string) (string, error) {

	info, err := PraseToken(token)

	if err != nil {
		return "", errors.New("Token Error")
	} else {
		return info.Username, nil
	}

}

func GetRefreshToken(username string) (string, error) {

	c := MyClaims{

		username, "RefreshToken", jwt.StandardClaims{
			ExpiresAt: 0,
			Issuer:    Author,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(Secret)

}

func PraseToken(tokenStr string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("PraseToken Error")

}

func JudgeAccessToken(msg string) (string, bool) {

	claims, err := PraseToken(msg)

	if err != nil {
		return "NULL", false
	}

	if claims.StandardClaims.ExpiresAt <= time.Now().Unix() || claims.StandardClaims.Issuer != Author || claims.TokenType != "AccessToken" {
		return "NULL", false
	}

	return claims.Username, true
}
