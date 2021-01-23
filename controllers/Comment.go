package controllers

import "ulivideoapi/models"

type CommentController struct {
	CommonController
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

// 获取评论列表
// @router /comment/list [*]
func (this *CommentController) List()  {
	episodesId, _ := this.GetInt("episodesId")
	//获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if episodesId == 0 {
		this.ReturnError(4001, "必须指定集数")
	}
	if limit == 0 {
		limit = 12
	}

	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err == nil {
		var data []CommentInfo
		var commentInfo CommentInfo

		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			// 用户信息
			commentInfo.UserInfo, _ = models.GetUserInfo(v.UserId)
			data = append(data, commentInfo)
		}

		this.ReturnSuccess("success", data, num)
	} else {
		this.ReturnError(4001, "没有相关内容")
	}
}

//保存评论
// @router /comment/save [*]
func (this *CommentController) Save() {
	content := this.GetString("content")
	uid, _ := this.GetInt("uid")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")

	if content == "" {
		this.ReturnError(4001, "内容不能为空")
	}
	if uid == 0 {
		this.ReturnError(4002, "请先登录")
	}
	if episodesId == 0 {
		this.ReturnError(4003, "必须指定评论剧集ID")
	}
	if videoId == 0 {
		this.ReturnError(4005, "必须指定视频ID")
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err == nil {
		this.ReturnSuccess("success", "", 1)
	} else {
		this.ReturnError(5000, err)
	}
}