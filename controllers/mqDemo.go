package controllers

import (
	"strconv"
	"time"
	"ulivideoapi/services/mq"
)

type MqDemoController struct {
	CommonController
}

//简单模式和work工作模式 push方法
// @router /mq/push [*]
func (this *MqDemoController) GetMq() {
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

//订阅模式 push方法
// @router /mq/fanout/push [*]
func (this *MqDemoController) GetFanout() {
	go func() {
		count := 0
		for {
			mq.PublishEx("ulivideo.demo.fanout", "fanout", "", "fanout"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	this.Ctx.WriteString("hello fanout")
}

// 路由模式
// @router /mq/direct/push [*]
func (this *MqDemoController) GetDirect() {
	go func() {
		count := 0
		for {
			route := count % 2
			err := mq.PublishEx("ulivideo.demo.direct", "direct", "routing.queue" + strconv.Itoa(route), "direct" + strconv.Itoa(count))
			if err != nil {
				panic(err)
			}
			count++
			time.Sleep(1 * time.Millisecond)
		}
	}()
	this.Ctx.WriteString("hello direct")
}

// 主题模式
// @router /mq/topic/push [*]
func (this *MqDemoController) GetTopic() {
	go func() {
		count := 0
		var route string
		for {
			c := count % 3
			if c == 0{
				route = "ulivideo.video"
			} else if c == 1 {
				route = "ulivideo.user"
			} else {
				route = "user.video"
			}

			err := mq.PublishEx("ulivideo.demo.topic", "topic", route, route + strconv.Itoa(count))
			if err != nil {
				panic(err)
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	this.Ctx.WriteString("hello topic")
}

// 死信队列push
// @router /mq/dlx/push [*]
func (this *MqDemoController) GetDlx() {
	go func() {
		count := 0
		for {
			err := mq.PublishDlx("ulivideo.demo.dlx", "ulivideo.demo", "dlx" + strconv.Itoa(count))
			if err != nil {
				panic(err)
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	this.Ctx.WriteString("hello dlx")
}

// 死信队列push
// @router /mq/dlx/two/push [*]
func (this *MqDemoController) GetTwoDlx() {
	go func() {
		count := 0
		for {
			err := mq.PublishEx("ulivideo.demo.dlx2", "fanout", "", "dlxtwo" + strconv.Itoa(count))
			if err != nil {
				panic(err)
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	this.Ctx.WriteString("hello dlx2")
}