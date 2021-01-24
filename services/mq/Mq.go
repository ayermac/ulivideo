package mq

import (
	"bytes"
	"fmt"

	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error)  {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

// 发布
func Publish(exchange string, queueName string, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	//创建通道channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// 创建队列
	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(body),
	})

	return err
}

// 接收
func Consumer(exchange string, queueName string, callback Callback) {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	//创建通道channel
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	// 创建队列
	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
		}
	}()

	fmt.Printf("Waiting for message")
	<-forever
}

func BytesToString(b *[]byte) *string  {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}