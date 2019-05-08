package endpointtests

import (
	"path/filepath"
	"rest-demo/models"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

const tokenString = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxZGNkZDI2NC1kMzNhLTRhZDctOGFhZC02OGJmNzU1YzdkNzkiLCJ1c2VybmFtZSI6InRhbmdodWFuIiwidGltZXN0YW1wIjoxNTU3MDQ1ODg2Mjc4LCJpYXQiOjE1NTcwNDU4ODZ9.Tq6z3pVtfAmn5hkRYNb7UZp8e9korcMOmsDfrz8a2GUDO7caO66XVa9wsVMX0w5RI9EZlgWhBAuhO17m8zwXEs8TlAiLQaqdbIiAzbs49TLh2DBCDBUII3L44qez0RlC3u6xwGR0j0UkJWOanqlHDKNrbvrcDSUK51XEBBPaRzvxSUF54viZXnvcVbF8-VV5aLp6CuQGlZ_6zqvy-whmQbsWJd87QGtpNDBvJkNjgK6kkq6eJKOeyt-raxW9JHdt71KvsD0LJVfuH3FuU5-Ff706x-He2A2hv6qekKW2_PqQ36fmZmnuSOtrKiVvKyj2ct_I2DWeZ96Q6JezR7rgeA"

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../../")))
	beego.TestBeegoInit(apppath)
	_, err := getDBConnection()
	if err != nil {
		glog.Errorln(err)
	}
}

func getDBConnection() (db *gorm.DB, e error) {
	dbHost := beego.AppConfig.String("dbHost")
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName")
	dbPass := beego.AppConfig.String("dbPass")
	conn, err := models.Connect(dbHost, dbName, dbUser, dbPass)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
