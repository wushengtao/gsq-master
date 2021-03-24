package gsq

//topic 定义
type Topic struct {
	name       string
	channelMap map[string]*Channel
}

func NewTopic(name string) *Topic {
	return &Topic{
		name: name,
	}
}

func (topic *Topic) GetChannelMap() map[string]*Channel {
	return topic.channelMap
}
