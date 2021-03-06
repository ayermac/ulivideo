package models

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"time"
	"ulivideoapi/services/mq"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

//保存通知消息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	messageId, err := o.Insert(&message)
	return messageId, err
}

//保存消息接收人
func SendMessageUser(userId int, messageId int64) error {
	o := orm.NewOrm()
	var messageUser MessageUser
	messageUser.UserId = userId
	messageUser.MessageId = messageId
	messageUser.Status = 1
	messageUser.AddTime = time.Now().Unix()
	_, err := o.Insert(&messageUser)
	return err
}

// 保存消息接收人到Mq
func SendMessageUserToMq(userId int, messageId int64) error {
	type Data struct {
		UserId int
		MessageId int64
	}
	var data Data
	data.UserId = userId
	data.MessageId = messageId
	dataJson, _ := json.Marshal(data)
	err := mq.PublishEx("ulivideo", "direct", "ulivideo.message", string(dataJson))
	return err
}
