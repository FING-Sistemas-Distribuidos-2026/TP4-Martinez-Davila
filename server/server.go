package main

import (
	"fmt"
	"time"

	"gopkg.in/zeromq/goczmq.v4"
)

const SLEEP = 500 * time.Millisecond

func main() {
	pusher := goczmq.NewSock(goczmq.Push)
	defer pusher.Destroy()

	_, err := pusher.Bind("tcp://*:3000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Producer iniciado!")

	count := 0
	for count < 100 {
		count++
		msg := fmt.Sprintf("Message #%d", count)
		fmt.Printf("Enviando mensaje #%d\n", count)
		err := pusher.SendFrame([]byte(msg), goczmq.FlagNone)
		if err != nil {
			panic(err)
		}
		time.Sleep(SLEEP)
	}
}
