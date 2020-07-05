package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"time"
)

type UserClaim struct {
	Username string `json:"username"`
	Sign string `json:"sign"`
	jwt.StandardClaims
}

func duicheng() {
	sec := []byte("123abc")
	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{Username: "zhangsan",Sign:"hahaha"})
	token, _ := token_obj.SignedString(sec)
	fmt.Println(token)
	uc := UserClaim{}
	getToken, _ := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
		return sec, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Username)
	}
}
func main() {

	priKeyBytes, err := ioutil.ReadFile("./pem/private.pem")
	if err != nil {
		log.Fatal(err)
	}
	pubKeyBytes, err := ioutil.ReadFile("./pem/public.pem")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, UserClaim{Username: "zhangsan",Sign:"hahaha"})
	token, _ := token_obj.SignedString(privateKey)
	fmt.Println(token)
	uc := UserClaim{}
	uc.ExpiresAt = time.Now().Add(time.Second * 5).UnixNano()
	getToken, _ := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
		return publicKey, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Username)
		fmt.Println(getToken.Claims.(*UserClaim).Sign)
		fmt.Println(getToken.Claims.(*UserClaim).ExpiresAt)
	}
}
