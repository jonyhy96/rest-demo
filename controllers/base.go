package controllers

import (
	"runtime/debug"
	"rest-demo/models"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type baseController struct {
	beego.Controller
}

// 200
func (bc *baseController) success(resp interface{}) {
	glog.V(4).Infof("success resp[%+v]\n", resp)
	bc.returnJSON(200, resp)
}

// customError custom code
func (bc *baseController) customError(statusCode int, errorType string, resp interface{}) {
	var response = map[string]interface{}{
		"type":    errorType,
		"messgae": resp,
	}
	glog.Errorln(resp)
	bc.returnJSON(statusCode, response)
}

// 500
func (bc *baseController) serviceError(resp interface{}) {
	var response = map[string]interface{}{
		"type":    "DatabaseError",
		"messgae": resp,
	}
	glog.Errorf("serviceError %+v\n", resp)
	bc.returnJSON(500, response)
}

// 400 参数验证失败
func (bc *baseController) badRequest(resp interface{}) {
	var response = map[string]interface{}{
		"error": resp,
	}
	glog.Errorf("Bad Request %+v\n", resp)
	bc.returnJSON(400, response)
}

// 406 验证失败
func (bc *baseController) notAcceptable(resp interface{}) {
	var response = map[string]interface{}{
		"type":   "ValidationError",
		"errors": resp,
	}
	glog.Errorf("notAcceptable %+v\n", resp)
	bc.returnJSON(406, response)
}

func (bc *baseController) catchException() {
	if exp := recover(); exp != nil {
		debug.PrintStack()
		glog.Errorln(string(debug.Stack()))
		glog.Errorln(exp)
		bc.serviceError(models.Error{Error: "panic", Code: 0})
	}
}

func (bc *baseController) returnJSON(status int, resp interface{}) {
	bc.Ctx.Output.SetStatus(status)
	bc.Data["json"] = resp
	bc.ServeJSON()
	bc.Finish()
}
