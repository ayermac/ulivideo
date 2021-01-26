package main

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
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
	mq.ConsumerDlx(
		"ulivideo.comment.count",
		"ulivideo.comment.queue",
		"ulivideo.comment.count.dlx",
		"ulivideo.comment.queue.dlx",
		"ulivideo.comment.#",
		10000,//10s
		callback,
		)
}

func callback(s string)  {
	o := orm.NewOrm()
	type Data struct {
		VideoId int
		EpisodesId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	if err == nil {
		// 修改视频总评论数
		o.Raw("UPDATE video SET comment = comment + 1 WHERE id = ?", data.VideoId).Exec()
		// 修改视频剧集评论数
		o.Raw("UPDATE video_episodes SET comment = comment + 1 WHERE id = ?", data.EpisodesId).Exec()
		// 更新Redis排行榜-通过MQ
		// 发布订阅简单模式
		videoObj := map[string]int{
			"VideoId": data.VideoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		_ = mq.PublishEx("ulivideo", "direct", "ulivideo.top", string(videoJson))
	}
	fmt.Println("msg is :%s\n", s)
}