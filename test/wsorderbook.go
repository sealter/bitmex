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

		client.Subscribe([]string{"orderBookL2:XBTUSD", "orderBookL2:ETHH18"})
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

// func  Analysis() {
// 	if _, ok := rep["success"]; ok {
// 		log.Println("success:", rep["success"])
// 	} else if v, ok := rep["table"]; ok {
// 		if v == "orderBookL2" {
// 			c.commands[fmt.Sprintf("orderBookL2:%v",rep["data"]["symbol"])] <-
// 		}else{
// 			c.commands[] <-
// 		}
// 	}
// }
// func (c *WebSocketClient) Subscribe(commands Command) {
// 	c.commands = commands
// 	var args []string
// 	for command, _ := range commands {
// 		args = append(args, command)

// 	}
// 	c.Write(Msg{"subscribe", args})
// }
