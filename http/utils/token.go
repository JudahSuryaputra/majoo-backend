package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func EncodeAuthToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["ID"] = id
	claims["UserName"] = username
	claims["CreatedAt"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(viper.GetString("MAJOO_SECRET")))
}
