package main

import (
	"fmt"
)

func main() {
	//mq.Consumer("", "fyouku_demo", callback)
	//mq.ConsumerEx("ulivideo.demo.fanout", "fanout", "", callback)
	//mq.ConsumerEx("ulivideo.demo.direct", "direct", "routing.queue0", callback)
	//mq.ConsumerEx("ulivideo.demo.topic", "topic", "ulivideo.*", callback)
	//mq.ConsumerDlx("ulivideo.demo.dlx", "ulivideo_demo_dlx", "ulivideo.demo.dlx2", "ulivideo_demo_dlx2", 1000, callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
