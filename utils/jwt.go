package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var (
	secretKey []byte = []byte(viper.GetString("secretkey"))
	expTime   int64  = viper.GetInt64("secretkey")
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func ObtainJWT(username string) (string, error) {

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expTime) * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	if err != nil {
		//beego.Error(err)
		return "", errors.New("token 生成失败")
	}
	return ss, nil
}
func RefreshJWT(strToken string) (string, error) {
	token, err := jwt.ParseWithClaims(strToken, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		//beego.Error(err)
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		println("trst")
		return "", err
	}
	if err := token.Claims.Valid(); err != nil {
		//beego.Error(err)
		return "", err
	}

	newToken, err := ObtainJWT(claims.Username)
	return newToken, err
}
func VerifyJWT(strToken string) (bool, error) {
	token, err := jwt.ParseWithClaims(strToken, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		//beego.Error(err)
		return false, err
	}
	if err := token.Claims.Valid(); err != nil {
		//beego.Error(err)
		return false, err
	}
	return true, nil

}
