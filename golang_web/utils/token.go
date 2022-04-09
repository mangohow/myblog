package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/2/24 18:59
* @Desc:
 */

type Claim struct {
	Username string
	UserId   uint32
	jwt.StandardClaims
}

var jwtKey = []byte("golang-im-server")

func CreateToken(id uint32, username string, expireDuration time.Duration) (string, error) {
	claims := &Claim{
		Username: username,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mgh",
			Subject:   "User_Token",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(token string) (string, uint32, bool) {
	if token == "" {
		return "", 0, false
	}

	tok, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		Logger().Warning("ParseWithClaims error %v", err)
		return "", 0, false
	}

	if claims, ok := tok.Claims.(*Claim); ok && tok.Valid {
		return claims.Username, claims.UserId, true
	} else {
		Logger().Warning("%v", err)
		return "", 0, false
	}

}
