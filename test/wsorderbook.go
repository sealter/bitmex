package main

import (
	"log"
	"time"

	"github.com/sunrisedo/bitmex/ws"
)

var DataSignal = ws.NewDataSignal()

func main() {
	loop := func() {
		client, err := ws.NewWsClient(DataSignal)
		if err != nil {
			log.Printf("Websocket connect error:%v", err)
			return
		}

		client.Subscribe([]string{"orderBookL2:XBTUSD"})
		if err := client.Run(); err != nil {
			log.Println("Websocket break:", err)
		}
	}

	go func() {
		for {
			loop()
			time.Sleep(time.Second * 1)
			log.Println("Websocket reconnection.")
		}
	}()

	go deal()
	select {}
}

func deal() {
	for {
		select {
		case in := <-DataSignal:
			log.Println("DataSignal:", in)
		}
	}
}
