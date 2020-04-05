package main

import (
	"encoding/json"
	"log"

	"github.com/patrickbucher/zmq-playground/correlation/payloads"
	zmq "github.com/zeromq/goczmq"
)

func main() {
	pull := zmq.NewPullChanneler("tcp://0.0.0.0:5555")
	push := zmq.NewPushChanneler("tcp://0.0.0.0:5556")
	defer pull.Destroy()
	defer push.Destroy()

	for {
		select {
		case data := <-pull.RecvChan:
			var payload payloads.RequestPayload
			err := json.Unmarshal(data[0], &payload)
			if err != nil {
				log.Printf("unmarshal %v: %v\n", data[0], err)
				continue
			}
			c := payload.A + payload.B

			response := payloads.ResponsePayload{
				C:             c,
				CorrelationID: payload.CorrelationID,
			}
			responseData, err := json.Marshal(response)
			if err != nil {
				log.Printf("marshal %v: %v\n", response, err)
				continue
			}
			data[0] = responseData
			push.SendChan <- data
		}
	}
}
