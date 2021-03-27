package client

import "gsq-master/gsq"

type GsqClient struct {
	gsqd *gsq.Gsqd
}

func NewGsqClient(gsqd *gsq.Gsqd) *GsqClient {
	return &GsqClient{
		gsqd: gsqd,
	}
}

func (gsqClient *GsqClient) Publish(topicName string, messgae string) error {
	return gsqClient.gsqd.Publish(topicName, messgae)
}

func (gsqClient *GsqClient) Subscribe(topicName string, channelName string) error {
	return gsqClient.gsqd.Subscribe(topicName, channelName)
}
