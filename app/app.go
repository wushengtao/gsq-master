package main

import (
	"gsq-master/client"
	"gsq-master/gsq"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//新建一个gsqd
	gsqd := gsq.NewGsqd()

	//发消息
	publish := client.NewGsqClient(gsqd)
	publish.Publish("create", "this is a message")
	wg.Done()

	//接受消息
	subscribe1 := client.NewGsqClient(gsqd)
	subscribe1.Subscribe("create", "c1")

	wg.Wait()

}
