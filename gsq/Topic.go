package gsq

//topic 定义
type Topic struct {
	name       string
	channelMap map[string]*Channel
}

func NewTopic(name string) *Topic {
	channelMap := make(map[string]*Channel)
	channelMap["c1"] = NewChannel()
	return &Topic{
		name:       name,
		channelMap: channelMap,
	}
}

func (topic *Topic) GetChannelMap() map[string]*Channel {
	return topic.channelMap
}
