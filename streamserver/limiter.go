package main

import "log"

type LimitConn struct {
	defineConnectNum int
	buffer chan int
}

func NewConnLimiter(cc int) *LimitConn {
	lc := &LimitConn{
		defineConnectNum: cc,
		buffer: make(chan int, cc),
	}
	return lc
}

func (lc *LimitConn) checkConn() bool {
	if len(lc.buffer) >= lc.defineConnectNum {
		log.Printf("Reached the rate limitation.")
		return false
	}

	lc.buffer <- 1
	return true
}

func (lc *LimitConn) release() {
	c := <- lc.buffer
	log.Printf("New connect comming: %d", c)
}