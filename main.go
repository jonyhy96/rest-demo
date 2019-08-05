package main

import (
	"flag"
	"rest-demo/controllers"
	"rest-demo/models"
	_ "rest-demo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	conn, err := models.InitDB()
	if err != nil {
		glog.Errorln(err)
	}
	defer glog.Flush()
	defer conn.Close()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
