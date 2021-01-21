package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
}

func init()  {
	orm.RegisterModel(new(Video))
}

func GetChannelHotList(channelId int) (int64, []VideoData, error)  {
	o := orm.NewOrm()
	var hotList []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_hot=1 AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", channelId).QueryRows(&hotList)
	return num, hotList, err
}

func GetChannelRecommendList(channelId int, regionId int) (int64, []VideoData, error)  {
	o := orm.NewOrm()
	var recommendList []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_recommend=1 AND channel_id=? AND region_id = ? ORDER BY episodes_update_time DESC LIMIT 9", channelId, regionId).QueryRows(&recommendList)
	return num, recommendList, err
}

func GetChannelRecommendTypeList(channelId int, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var recommendList []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_recommend=1 AND channel_id=? AND type_id = ? ORDER BY episodes_update_time DESC LIMIT 9", channelId, typeId).QueryRows(&recommendList)
	return num, recommendList, err
}

func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params
	
	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}

	switch sort {
		case "episodesUpdateTime":
			qs = qs.OrderBy("-episodes_update_time")
		case "comment":
			qs = qs.OrderBy("-comment")
		case "addTime":
			qs = qs.OrderBy("-add_time")
		case "default":
			qs = qs.OrderBy("-add_time")	
	}

	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	return nums, videos, err
}