package controllers

import (
	"ulivideoapi/models"
)

type BaseController struct {
	CommonController
}

//获取频道地区列表
// @router /channel/region [*]
func (this *BaseController) ChannelRegion() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.ReturnError(4001, "必须指定频道")
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err == nil {
		this.ReturnSuccess("success", regions, num)
	} else {
		this.ReturnError(4001, "没有频道数据")
	}
}

//获取频道类型列表
// @router /channel/type [*]
func (this *BaseController) ChannelType() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.ReturnError(4001, "必须指定频道")
	}

	num, types, err := models.GetChannelType(channelId)
	if err == nil {
		this.ReturnSuccess("success", types, num)
	} else {
		this.ReturnError(4001, "没有频道数据")
	}
}
