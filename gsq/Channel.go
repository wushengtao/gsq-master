package gsq

import "sync"

type Channel struct {
	sync.RWMutex
	MemoryMsgChan chan *Message
}

func NewChannel() *Channel {
	return &Channel{
		MemoryMsgChan: make(chan *Message, 100),
	}
}
