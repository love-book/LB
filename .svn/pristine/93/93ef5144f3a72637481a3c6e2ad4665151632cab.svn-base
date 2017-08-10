package common

import (
	//"fmt"
	"time"
	//"crypto/elliptic"
	"github.com/dgrijalva/jwt-go"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego"
	//"fmt"
)

type Claims struct {
	Appid string `json:"appid"`
	// recommended having
	jwt.StandardClaims
}
var (
	key []byte = []byte ("-jwt_secrt_LB163@why.com")
)

func GetToken(appid string) (string,error)  {
	expireToken := time.Now().Add(time.Hour * 2).Unix()
	claim := Claims{
		appid,
		jwt.StandardClaims{
			NotBefore:int64(time.Now().Unix()),
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	strtoken, err := token.SignedString(key)

	return strtoken,err

}

func VerifyToken(token string) bool  {
	_ , err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {

		return false
	}
	return true
}