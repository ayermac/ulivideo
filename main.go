package main

import (
	_ "ulivideoapi/routers"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	defaultdb, _ := beego.AppConfig.String("defaultdb")
	err := orm.RegisterDataBase("default", "mysql", defaultdb)
	if err != nil {
		panic(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
