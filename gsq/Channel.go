package gsq

import "sync"

type Channel struct {
	sync.RWMutex
	channelName string
}
