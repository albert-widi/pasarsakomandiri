package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
)

type Token struct {
	SecretKey string        	`json:"SecretKey"`
	ExpiredMin time.Duration    `json:"ExpiredMin"`
}

var (
	tokenConfig Token
)

func Configure(t Token) {
	tokenConfig = t
}

func CreateTokenByUserId(id int64) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["id"] = id
	token.Claims["Exp"] = time.Now().Add(time.Minute * tokenConfig.ExpiredMin).Unix()
	tokenString, err := token.SignedString([]byte(tokenConfig.SecretKey))

	return tokenString, err
}


func ClearTokenSession(tokenString string) {
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error){
		return []byte(tokenConfig.SecretKey), nil
	})

	if err != nil {
		log.Println(err)
	}

	//set token be expired immediately
	token.Claims["Exp"] = time.Now().Add(time.Minute * 0).Unix()
}

func IsTokenVerified(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenConfig.SecretKey), nil
	})

	if err != nil && !token.Valid {
		return  false
	}

	return true
}

func NewToken() *jwt.Token {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["Exp"] = time.Now().Add(time.Minute * tokenConfig.ExpiredMin).Unix()
	return token
}
