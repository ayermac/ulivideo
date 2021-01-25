package main

import (
	"fmt"
	"ulivideoapi/services/mq"
)

func main() {
	//mq.Consumer("", "fyouku_demo", callback)
	//mq.ConsumerEx("ulivideo.demo.fanout", "fanout", "", callback)
	mq.ConsumerEx("ulivideo.demo.direct", "direct", "routing.queue0", callback)
	//mq.ConsumerEx("ulivideo.demo.topic", "topic", "ulivideo.*", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
