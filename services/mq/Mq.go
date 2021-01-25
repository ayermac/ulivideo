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
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
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
		true,
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
			_ = d.Ack(false)
		}
	}()

	fmt.Printf("Waiting for message")
	<-forever
}

func PublishEx(exchange string, types string, routingKey string, body string) error {
	// 建立连接
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		return err
	}
	// 创建channel
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return err
	}

	// 创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//创建队列
	_, err = channel.QueueDeclare(
		routingKey,
		true,
		false,
		false,
		false,
		nil,
	)

	//绑定
	err = channel.QueueBind(
		routingKey,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "text/plain",
		Body: []byte(body),
	})

	return err
}

// 接收
func ConsumerEx(exchange string, types string, routingKey string, callback Callback) {
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

	// 创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	//创建队列
	_, err = channel.QueueDeclare(
		routingKey,
		true,
		false,
		false,
		false,
		nil,
	)

	//绑定
	err = channel.QueueBind(
		routingKey,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := channel.Consume(routingKey, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			_ = d.Ack(false)
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