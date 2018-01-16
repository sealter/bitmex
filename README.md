# BITMEX-API

###  REST API EXMPLE :
```
package main

import (
	"log"
	"github.com/sunrisedo/bitmex/bitmex"
)

func main() {
    cfg := bitmex.NewConfigurationWithKey("key", "secret")
    userApi := bitmex.NewUserApiWithConfig(cfg)
    wallet, response, err := userApi.UserGetWallet("")
    if err != nil {
    	log.Println("error: ", err)
        return
    }
    log.Println(response.Status)
    log.Println("wallet: ", *wallet)
}

```


###  WS EXMPLE :
```
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

```