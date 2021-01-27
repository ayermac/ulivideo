package controllers

import (
	"encoding/json"
	"strconv"
	//beego "github.com/beego/beego/v2/server/web"
	"ulivideoapi/models"
	"ulivideoapi/services/es"
)

type VideoController struct {
	CommonController
}

// 频道页-获取顶部广告
// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert()  {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.ReturnError(4004, "必须指定频道")
	}
	num, adverts, err := models.GetChannelAdvert(channelId)
	if err == nil {
		this.ReturnSuccess("success", adverts, num)
	} else {
		this.ReturnError(4001, "没有频道数据")
	}
}

// 频道页-获取正在热播
// @router /channel/hot [*]
func (this *VideoController) ChannelHotList()  {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.ReturnError(4004, "必须指定频道")
	}
	num, hotList, err := models.GetChannelHotList(channelId)
	if err == nil {
		this.ReturnSuccess("success", hotList, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}

// 频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (this *VideoController) ChannelRecommendList()  {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")

	if channelId == 0 {
		this.ReturnError(4004, "必须指定频道")
	}
	if regionId == 0 {
		this.ReturnError(4004, "必须指定地区")
	}
	num, recommendList, err := models.GetChannelRecommendList(channelId, regionId)
	if err == nil {
		this.ReturnSuccess("success", recommendList, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}

//频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecommendTypeList()  {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.ReturnError(4004, "必须指定频道")
	}
	if typeId == 0 {
		this.ReturnError(4004, "必须指定频道类型")
	}
	num, recommendList, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.ReturnSuccess("success", recommendList, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}

//根据传入参数获取视频列表
// @router /channel/video [*]
func (this *VideoController) ChannelVideo() {
	//获取频道ID
	channelId, _ := this.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := this.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := this.GetInt("typeId")
	//获取状态
	end := this.GetString("end")
	//获取排序
	sort := this.GetString("sort")
	//获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.ReturnError(4004, "必须指定频道")
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoListEs(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.ReturnSuccess("success", videos, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}

//我的视频管理
// @router /user/video [*]
func (this *VideoController) UserVideo() {
	uid, _ := this.GetInt("uid")
	if uid == 0 {
		this.ReturnError(4001, "必须指定用户")
	}
	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		this.ReturnSuccess("success", videos, num)
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
}

// 获取视频详情
// @router /video/info [*]
func (this *VideoController) VideoInfo()  {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.ReturnError(4001, "必须指定视频ID")
	}
	videos, err := models.RedisGetVideoInfo(videoId)
	if err == nil {
		this.ReturnSuccess("success", videos, 1)
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
}

//获取视频剧集列表
// @router /video/episodes/list [*]
func (this *VideoController) VideoEpisodesList()  {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.ReturnError(4001, "必须指定视频ID")
	}
	num, episodes, err := models.RedisGetVideoEpisodesList(videoId)
	if err == nil {
		this.ReturnSuccess("success", episodes, num)
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
}

//保存用户上传视频信息
// @router /video/save [*]
func (this *VideoController) VideoSave() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")
	aliyunVideoId := this.GetString("aliyunVideoId")
	if uid == 0 {
		this.ReturnError(4001, "请先登录")
	}
	if playUrl == "" {
		this.ReturnError(4002, "视频地址不能为空")
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err == nil {
		this.ReturnSuccess( "success", nil, 1)
	} else {
		this.ReturnError(5000, err)
	}
}

//搜索接口
// @router /video/search [*]
func (this *VideoController) Search() {
	//获取搜索关键字
	keyword := this.GetString("keyword")
	//获取翻页信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if keyword == ""{
		this.ReturnError(4001, "关键字不能为空")
	}

	if limit == 0 {
		limit = 12
	}

	sort := []map[string]string{map[string]string{"id": "desc"}}
	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"title": keyword,
				},
			},
		},
	}

	res := es.EsSearch("ulivideo", query, offset, limit, sort)
	total := res.Total.Value
	var data []models.Video

	for _, v := range res.Hits {
		var itemData models.Video
		err := json.Unmarshal([]byte(v.Source), &itemData)
		if err == nil {
			data = append(data, itemData)
		}
	}
	if total > 0 {
		this.ReturnSuccess("success", data, int64(total))
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
	
}

//导入ES脚本
// @router /video/send/es [*]
func (this *VideoController) SendEs() {
	_, data, _ := models.GetAllList()
	for _, v := range data {
		body := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"is_end":               v.IsEnd,
			"channel_id":           v.ChannelId,
			"status":               v.Status,
			"region_id":            v.RegionId,
			"type_id":              v.TypeId,
			"episodes_update_time": v.EpisodesUpdateTime,
			"comment":              v.Comment,
			"user_id":              v.UserId,
			"is_recommend":         v.IsRecommend,
		}
		es.EsAdd("ulivideo", "video-"+strconv.Itoa(v.Id), body)
	}
}
