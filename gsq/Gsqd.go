package gsq

import "sync"

type Gsqd struct {
	sync.RWMutex
	//topics
	topicMap map[string]*Topic
	//消费队列
	inFlightMessages map[string]*Message
}

func (gsqd *Gsqd) Publish(topicName string, msg interface{}) error {
	//校验topic是否存在，如果不存在，先创建这个topic，后续可配置是否自动创建
	topic := gsqd.topicMap[topicName]
	if topic == nil {
		//创建topic
		gsqd.Lock()
		topic = topic.NewTopic(topicName)
		gsqd.Unlock()

	}
	return nil

}

func Subscribe(topicName string, channelName string) error {
	return nil
}
