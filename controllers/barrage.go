package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"ulivideoapi/models"
)

type BarrageController struct {
	CommonController
}

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 获取弹幕websocket
// @router /barrage/ws [*]
func (this *BarrageController) BarrageWs() {
	var (
		conn *websocket.Conn
		err error
		data []byte
		barrages []models.BarrageData
	)
	if conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		_ = json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		//获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err := conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}
	}
ERR:
	conn.Close()
}

//保存弹幕
// @router /barrage/save [*]
func (this *BarrageController) Save() {
	uid, _ := this.GetInt("uid")
	content := this.GetString("content")
	currentTime, _ := this.GetInt("currentTime")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	if content == "" {
		this.ReturnError(4001, "弹幕不能为空")
	}
	if uid == 0 {
		this.ReturnError(4002, "请先登录")
	}
	if episodesId == 0 {
		this.ReturnError(4003, "必须指定剧集ID")
	}
	if videoId == 0 {
		this.ReturnError(4005, "必须指定视频ID")
	}
	if currentTime == 0 {
		this.ReturnError(4006, "必须指定视频播放时间")
	}

	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)
	if err == nil {
		this.ReturnSuccess("success", "", 1)
	} else {
		this.ReturnError(5000, err)
	}
}
