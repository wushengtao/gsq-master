package gsq

//topic 定义
type Topic struct {
	name       string
	channelMap map[string]*Channel
}

func (topic *Topic) NewTopic(name string) *Topic {
	return &Topic{
		name: name,
	}
}
