package main

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"ulivideoapi/models"
	"ulivideoapi/services/mq"
	redisClient "ulivideoapi/services/redis"
)

func main()  {
	beego.LoadAppConfig("ini", "../../../conf/app.conf")
	defaultdb, _ := beego.AppConfig.String("defaultdb")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", defaultdb)

	if err != nil {
		panic(err)
	}
	mq.ConsumerEx("ulivideo", "direct", "ulivideo.top", callback)
}

func callback(s string)  {
	type Data struct {
		VideoId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	videoInfo, err := models.RedisGetVideoInfo(data.VideoId)
	if err == nil {
		conn := redisClient.PoolConnect()
		defer conn.Close()
		// 更新排行榜
		redisChannelKey := "video:top:channel:channelId:" + strconv.Itoa(videoInfo.ChannelId)
		redisTypeKey := "video:top:channel:typeId:" + strconv.Itoa(videoInfo.TypeId)

		conn.Do("zincrby", redisChannelKey, 1, data.VideoId)
		conn.Do("zincrby", redisTypeKey, 1, data.VideoId)
	}
	fmt.Println("msg is :%s\n", s)
}