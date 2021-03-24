package main

import "gsq-master/client"

func main() {
	gsqClient := client.NewGsqClient()
	gsqClient.Publish("", "")

}
