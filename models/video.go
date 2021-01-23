package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
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

type Episodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
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

func GetUserVideo(uid int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count, is_end FROM video WHERE user_id=? ORDER BY add_time DESC", uid).QueryRows(&videos)
	return num, videos, err
}

func GetVideoInfo(videoId int) (Video, error)  {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id=?", videoId).QueryRow(&video)
	return video, err
}

func GetVideoEpisodesList(videoId int) (int64, []Episodes, error)  {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("SELECT * FROM video_episodes WHERE video_id=? AND status =1 ORDER BY num ASC", videoId).QueryRows(&episodes)
	return num, episodes, err
}

//频道排行榜
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return num, videos, err
}

//类型排行榜
func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return num, videos, err
}

func SaveVideo(title string, subTitle string, channelId int, regionId int, typeId int, playUrl string, user_id int, aliyunVideoId string) error {
	o := orm.NewOrm()
	var video Video
	time := time.Now().Unix()
	video.Title = title
	video.SubTitle = subTitle
	video.AddTime = time
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.ChannelId = channelId
	video.Status = 1
	video.RegionId = regionId
	video.TypeId = typeId
	video.EpisodesUpdateTime = time
	video.Comment = 0
	video.UserId = user_id
	videoId, err := o.Insert(&video)
	if err == nil {
		if aliyunVideoId != "" {
			playUrl = ""
		}
		_, err = o.Raw("INSERT INTO video_episodes (title,add_time,num,video_id,play_url,status,comment,aliyun_video_id) VALUES (?,?,?,?,?,?,?,?)", subTitle, time, 1, videoId, playUrl, 1, 0, aliyunVideoId).Exec()
		//fmt.Println(err)
	}
	return err
}

func SaveAliyunVideo(videoId string, log string) error {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO aliyun_video (video_id, log, add_time) VALUES (?,?,?)", videoId, log, time.Now().Unix()).Exec()
	fmt.Println(err)
	return err
}
