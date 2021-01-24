package main

import (
	"fmt"
	"ulivideoapi/services/mq"
)

func main() {
	mq.Consumer("", "fyouku_demo", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
