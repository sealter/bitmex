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

type BitmexDepth struct {
	Id     int64   `json:"id,omitempty"`
	Symbol string  `json:"symbol,omitempty"`
	Side   string  `json:"side,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Size   float64 `json:"size,omitempty"`
}

func deal() {
	for {
		select {
		case in := <-DataSignal:
			if in.Data == nil {
			} else {
				log.Println("DataSignal:", in)
				// for _, v := range in.Data.([]map[string]interface{}) {
				// 	log.Println("v:", v)
				// }

			}

		}
	}
}
