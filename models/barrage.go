package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Barrage struct {
	Id          int
	Content     string
	CurrentTime int
	AddTime     int64
	UserId      int
	Status      int
	EpisodesId  int
	VideoId     int
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init()  {
	orm.RegisterModel(new(Barrage))
}

// 弹幕列表
func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	num, err := o.Raw("SELECT id,content,`current_time` FROM barrage WHERE status=1 AND episodes_id=? AND `current_time`>=? AND `current_time`<? ORDER BY `current_time` ASC", episodesId, startTime, endTime).QueryRows(&barrages)
	return num, barrages, err
}

// 保存弹幕到数据库
func SaveBarrage(episodesId int, videoId int, currentTime int, userId int, content string) error {
	o := orm.NewOrm()
	var barrage Barrage
	barrage.Content = content
	barrage.CurrentTime = currentTime
	barrage.AddTime = time.Now().Unix()
	barrage.UserId = userId
	barrage.Status = 1
	barrage.EpisodesId = episodesId
	barrage.VideoId = videoId
	_, err := o.Insert(&barrage)
	return err
}