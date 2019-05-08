package utils

import (
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

// MakeJSON return a json object for return
// TODO make this function more scalable or deprecate this one
func MakeJSON(baseKey string, content interface{}) map[string]interface{} {
	result := map[string]interface{}{
		baseKey: content,
	}
	return result
}

//GetUserID get userID from token
func GetUserID(ctx *context.Context) (*string, bool) {
	var userID string
	code := strings.SplitN(ctx.Input.Header("Authorization"), " ", 2)
	token, _ := jwt.Parse(code[1], func(token *jwt.Token) (interface{}, error) {
		verifyBytes, _ := ioutil.ReadFile("conf/public.pem")
		res, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		return res, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["sub"]; ok {
			userID = claims["sub"].(string)
			return &userID, true
		}
	}
	return nil, false
}

// Contains verify if a filed is contains in a array
func Contains(name string, fields ...[]string) bool {
	for _, v := range fields[0] {
		if name == v {
			return true
		}
	}
	return false
}
