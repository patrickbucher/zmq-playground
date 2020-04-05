package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"
	"github.com/patrickbucher/zmq-playground/correlation/payloads"
	zmq "github.com/zeromq/goczmq"
)

func main() {
	push := zmq.NewPushChanneler("tcp://0.0.0.0:5555")
	pull := zmq.NewPullChanneler("tcp://0.0.0.0:5556")
	defer push.Destroy()
	defer pull.Destroy()

	lock := make(chan struct{}, 1)
	chans := make(map[string]chan payloads.ResponsePayload)

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		aparam := r.URL.Query().Get("a")
		bparam := r.URL.Query().Get("b")
		if aparam == "" || bparam == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		a, aerr := strconv.Atoi(aparam)
		b, berr := strconv.Atoi(bparam)
		if aerr != nil || berr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		correlationID, err := uuid.NewRandom()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		lock <- struct{}{}
		reader := make(chan payloads.ResponsePayload)
		<-lock

		chans[correlationID.String()] = reader
		request := payloads.RequestPayload{
			A:             a,
			B:             b,
			CorrelationID: correlationID.String(),
		}
		payload, err := json.Marshal(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Print(string(payload))
		data := make([][]byte, 1)
		data[0] = payload
		push.SendChan <- data

		response := <-reader
		w.Write([]byte(strconv.Itoa(response.C)))
	})

	go func() {
		for {
			select {
			case data := <-pull.RecvChan:
				var payload payloads.ResponsePayload
				err := json.Unmarshal(data[0], &payload)
				if err != nil {
					log.Printf("unmarshal %v: %v\n", data[0], err)
					continue
				}
				log.Println(payload)
				sink, ok := chans[payload.CorrelationID]
				if ok {
					sink <- payload
				}
			}
		}
	}()

	http.ListenAndServe("0.0.0.0:8080", nil)
}
