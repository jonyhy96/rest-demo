// Package routers project route
// @APIVersion 1.0.0
// @Title rest-demo
// @Description rest-demo api server write with beego.
// @Contact hy352144278@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"rest-demo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSBefore(auth),
		beego.NSGet("/healthCheck", func(ctx *context.Context) {
			_ = ctx.Output.Body([]byte("{\"status\":200}"))
		}),
	)
	beego.AddNamespace(ns)
}

func auth(ctx *context.Context) {
	var code []string
	if token := ctx.Input.Header("Authorization"); token == "" {
		errorResponse(ctx, http.StatusProxyAuthRequired)
		return
	}
	if code = strings.SplitN(ctx.Input.Header("Authorization"), " ", 2); code[0] != "Bearer" {
		errorResponse(ctx, http.StatusUnauthorized)
		return
	}
	b, err := jwt.Parse(code[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		verifyBytes, err := ioutil.ReadFile("conf/public.pem")
		res, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			panic(err)
		}
		return res, nil
	})
	if err != nil || !b.Valid {
		errorResponse(ctx, http.StatusUnauthorized)
		return
	}
}

// use for handle response error
// TODO make this function more scalable or deprecate this one
func errorResponse(ctx *context.Context, code int) {
	ctx.Output.SetStatus(code)
	var e models.HTTPError
	e.Message = http.StatusText(code)
	e.Status = code
	ctx.Output.JSON(e, false, true)
}
