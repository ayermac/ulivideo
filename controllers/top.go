package controllers

import "ulivideoapi/models"

type TopController struct {
	CommonController
}

//根据频道获取排行榜
// @router /channel/top [*]
func (this *TopController) ChannelTop() {
	//获取频道ID
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.ReturnError(4001, "必须指定频道")
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err == nil {
		this.ReturnSuccess("success", videos, num)
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
}

//根据类型获取排行榜
// @router /type/top [*]
func (this *TopController) TypeTop() {
	typeId, _ := this.GetInt("typeId")
	if typeId == 0 {
		this.ReturnError(4001, "必须指定类型")
	}
	num, videos, err := models.RedisGetTypeTop(typeId)
	if err == nil {
		this.ReturnSuccess("success", videos, num)
	} else {
		this.ReturnError(4004, "没有相关内容")
	}
}