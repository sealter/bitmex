package ws

import (
	"errors"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const (
	ORIGIN  = "http://localhost/"
	BITWS   = "wss://www.bitmex.com/realtime"
	WBUFFER = 128
	RBUFFER = 512
)

type WebSocketClient struct {
	heart       *Heart
	ws          *websocket.Conn
	readSignal  DataSignal
	writeSignal chan interface{}
	doneSignal  chan bool
	err         error
}

func NewWsClient(data DataSignal) (*WebSocketClient, error) {
	ws, err := websocket.Dial(BITWS, "", ORIGIN)
	return &WebSocketClient{
		NewHeart(),
		ws,
		data,
		make(chan interface{}, WBUFFER),
		make(chan bool, 10),
		nil,
	}, err
}

func NewDataSignal() DataSignal {
	return make(DataSignal, RBUFFER)
}

func (c *WebSocketClient) ListenWrite() {
	for {
		select {
		case <-c.doneSignal:
			c.Done("write.")
			return
		case data := <-c.writeSignal:
			// log.Println("writeSignal:", data)
			if err := websocket.JSON.Send(c.ws, data); err != nil {
				log.Printf("Websocket send data error:%v,data:%v", err, data)
			}
		}
	}
	log.Println("ListenWrite closed.")
}

func (c *WebSocketClient) ListenRead() {
	for {
		select {
		case <-c.doneSignal:
			c.Done("read.")
			log.Println("ListenRead closed.")
			return
		default:
			var rep DataStruct
			err := websocket.JSON.Receive(c.ws, &rep)
			if err == io.EOF {
				c.err = errors.New("Webscoket connection break.")
				c.Done("read.")
			} else if err != nil {
				c.heart.Start()
				var msg string
				if err = websocket.Message.Receive(c.ws, &msg); err != nil {
					log.Printf("Webscoket read error:%v", err)
				}
			} else {
				c.heart.Start()
				c.readSignal <- &rep
			}
		}
	}
}

func (c *WebSocketClient) ListenHeart() {
	for {
		select {
		case <-c.doneSignal:
			c.Done("heart.")
			return
		case <-c.heart.ticker.C:
			c.Write("ping")
			c.heart.cnt++
			log.Println("Send ping times:", c.heart.cnt)
			if c.heart.cnt > 3 {
				c.err = errors.New("Webscoket connection timeout.")
				c.Done("heart.")
				return
			}
		default:
			c.heart.Reset()
		}
	}

}

func (c *WebSocketClient) Run() error {
	go c.ListenWrite()
	go c.ListenRead()
	c.ListenHeart()
	return c.err
}

func (c *WebSocketClient) Write(in interface{}) {
	c.writeSignal <- in
}

func (c *WebSocketClient) Done(id string) {
	c.doneSignal <- true
	log.Println("Close Websocket ", id)
}

func (c *WebSocketClient) Subscribe(args []string) {
	c.Write(Msg{"subscribe", args})
}
