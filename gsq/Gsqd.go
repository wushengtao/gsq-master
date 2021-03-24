package gsq

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Gsqd struct {
	exit chan bool
	sync.RWMutex
	//topics
	topicMap map[string]*Topic
	//消费队列
	//inFlightMessages map[string]*Message
}

func (gsqd *Gsqd) Publish(topicName string, msg string) error {
	//判断是否关闭
	select {
	case <-gsqd.exit:
		return errors.New("")
	default:

	}
	//校验topic是否存在，如果不存在，先创建这个topic，后续可配置是否自动创建
	topic := gsqd.topicMap[topicName]
	if topic == nil {
		topic = NewTopic(topicName)
		gsqd.Lock()
		gsqd.topicMap[topicName] = topic
		gsqd.Unlock()
	}
	//广播消息
	channelMap := topic.GetChannelMap()
	//每个channel需要复制一份数据
	for _, channel := range channelMap {
		broadcast(msg, channel)
	}
	return nil

}

func (gsqd *Gsqd) Subscribe(topicName string, channelName string) error {
	//新建channel,持久化之后直接读取
	channel := NewChannel()
	topic := gsqd.topicMap[topicName]
	channelMap := topic.GetChannelMap()

	gsqd.Lock()
	channelMap[channelName] = channel
	gsqd.Unlock()

	for {
		select {
		case msg := <-channel.MemoryMsgChan:
			fmt.Println(msg)
		default:
			fmt.Println("continue")
		}

	}

}

//广播消息
func broadcast(message string, channel *Channel) error {
	msg := NewMessage(time.Now().UnixNano(), message)
	select {
	case channel.MemoryMsgChan <- msg:
	default:
		//超时处理
	}
	return nil
}
