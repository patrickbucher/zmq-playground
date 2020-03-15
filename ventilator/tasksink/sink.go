package main

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

func main() {
	channeler := zmq.NewPullChanneler("tcp://0.0.0.0:5558")
	defer channeler.Destroy()

	for message := range channeler.RecvChan {
		fmt.Println(string(message[0]))
	}
}
