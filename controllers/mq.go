package controllers

import (
	"strconv"
	"time"
	"ulivideoapi/services/mq"
)

type MqController struct {
	CommonController
}

//简单模式和work工作模式 push方法
// @router /mq/push [*]
func (this *MqController) GetMq() {
	go func() {
		count := 0
		for {
			mq.Publish("", "fyouku_demo", "hello"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	this.Ctx.WriteString("hello")
}