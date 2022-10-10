package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

/*
   @Auth: menah3m
   @Desc: 处理Token相关
*/

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	secret, _ := beego.AppConfig.String("TokenSecret")
	return []byte(secret)

}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	tokenExpireStr, _ := beego.AppConfig.String("TokenExpire")
	tokenExpireInt, _ := strconv.Atoi(tokenExpireStr)
	expireTime := nowTime.Add(time.Hour * time.Duration(tokenExpireInt))
	issuer, _ := beego.AppConfig.String("Issuer")
	claims := Claims{
		Username: EncodeMD5(appKey),
		Password: EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
