package ws

import (
	"time"
)

const TFREQ = time.Second * 6

type Heart struct {
	cnt    int
	start  bool
	ticker *time.Ticker
}

func NewHeart() *Heart {
	return &Heart{
		0,
		true,
		time.NewTicker(TFREQ),
	}
}

func (c *Heart) Start() {
	c.start = true
}

func (c *Heart) Reset() {
	if c.start {
		c.cnt = 0
		c.start = false
		c.ticker = time.NewTicker(TFREQ)
	}
}
