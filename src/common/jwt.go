package common

import (
	//"fmt"
	"time"
	//"crypto/elliptic"
	"github.com/dgrijalva/jwt-go"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego"
	"fmt"
)

type Claims struct {
	Appid string `json:"appid"`
	jwt.StandardClaims
}
var (
	tokenKey []byte = []byte ("-jwt_secrt_LB163@why.com")
)

func SetToken(appid string) (string,error)  {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claim := Claims{
		appid,
		jwt.StandardClaims{
			NotBefore:int64(time.Now().Unix()),
			ExpiresAt: expireToken,
			Issuer:    appid,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	strtoken, err := token.SignedString(tokenKey)
	return strtoken,err
}

func GetToken(tokenString string)(claims *Claims,ok bool){
	//return nil,false
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err!=nil{
		return nil,false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Appid, claims.StandardClaims.ExpiresAt)
		return claims, ok
	} else {
		fmt.Println(err.Error())
		return nil,false
	}
	return nil,false
}

func VerifyToken(tokenString string) bool{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err!= nil{
		return false
	}
	if token.Valid {
		fmt.Println("验证成功")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return false
}

