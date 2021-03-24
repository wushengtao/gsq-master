package client

import "gsq-master/gsq"

type GsqClient struct {
	gsqd *gsq.Gsqd
}

func NewGsqClient() *GsqClient {
	return &GsqClient{
		gsqd: &gsq.Gsqd{},
	}
}

func (gsqClient *GsqClient) Publish(topicName string, messgae string) error {
	return gsqClient.gsqd.Publish(topicName, messgae)
}
