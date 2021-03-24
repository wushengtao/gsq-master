package gsq

import (
	"time"
)

//静态变量
const (
	MsgIDLength = 16
)

//type MessageID [MsgIDLength]byte

type Message struct {
	//消息id
	ID int64
	//消息体
	Body string
	//创建时间
	Timestamp int64
}

func NewMessage(id int64, body string) *Message {
	return &Message{
		ID:        id,
		Body:      body,
		Timestamp: time.Now().UnixNano(),
	}
}
