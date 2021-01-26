package main

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"ulivideoapi/models"
	"ulivideoapi/services/mq"
)

func main()  {
	beego.LoadAppConfig("ini", "../../../conf/app.conf")
	defaultdb, _ := beego.AppConfig.String("defaultdb")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", defaultdb)

	if err != nil {
		panic(err)
	}
	mq.ConsumerEx("ulivideo", "direct", "ulivideo.message", callback)
}

func callback(s string)  {
	type Data struct {
		UserId int
		MessageId int64
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		_ = models.SendMessageUser(data.UserId, data.MessageId)
	}
	fmt.Println("msg is :%s\n", s)
}