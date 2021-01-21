package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"ulivideoapi/models"
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

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.ReturnSuccess("success", videos, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}