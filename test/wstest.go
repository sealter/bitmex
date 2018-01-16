package main

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const (
	origin = "http://localhost/"
	url    = "wss://www.bitmex.com/realtime"
)

type Responses map[string]interface{}

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal("Webscoket connect error:%v", err)
	}

	_, err = ws.Write([]byte(`{"op": "subscribe", "args": ["orderBookL2:XBTUSD","orderBookL2:ETHH18"]}`))

	// _, err = ws.Write([]byte("ping"))
	if err != nil {
		log.Printf("Websocket send data error:%v", err)
	}

	go func() {
		for {
			select {
			default:
				var rep Responses
				err := websocket.JSON.Receive(ws, &rep)
				if err == io.EOF {
					log.Printf("Webscoket read no data.")
				} else if err != nil {
					log.Printf("Webscoket read error:%v", err)
				} else {
					log.Printf("Webscoket read data:%v", rep["data"].([0]["symbol"])
					// c.Command(rep.Table) <- rep.Data()
				}
			}
		}

	}()

	select {}
}
