package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClass struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	mySigningKey := []byte("mynameiszz")
	c := MyClass{
		Username: "zz",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //到期时间
			Issuer:    "zzt",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"exp":      time.Now().Unix() + 5, //到期时间
	//	"iss":      "zy",                  //签发人
	//	"nbf":      time.Now().Unix() - 5,
	//	"username": "my",
	//})

	//加密
	s, e := t.SignedString(mySigningKey)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(s)
	}

	token, err := jwt.ParseWithClaims(s, &MyClass{}, func(token *jwt.Token) (interface{}, error) {
		//token, err := jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("%s", err)
	}
	fmt.Println(token.Claims)
	//fmt.Println(token.Claims.(*MyClass))
	//fmt.Println(token.Claims.(*MyClass).Username)
	//fmt.Println(*(token.Claims.(*jwt.MapClaims)))
	//fmt.Println(token.Claims.(*jwt.MapClaims)["username"].(string))//获取username报错了，暂未找到原因

}
