package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var (
	reqCount  int64
	lastCount int64
)

// 定时统计输出当前QPS
func loopCount() {
	for {
		qps := (reqCount - lastCount) / 10
		lastCount = reqCount
		fmt.Printf("qps:%d totle:%d\n", qps, reqCount)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go loopCount()

	ln, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// 每个Client一个Goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	for {
		// 读取客户端消息
		var body [5]byte
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		//fmt.Printf("收到%s消息: %s\n", addr, string(body[:]))
		// 回包
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
		atomic.AddInt64(&reqCount, 1)
		//fmt.Printf("发送给%s: %s\n", addr, string(body[:]))
	}
	fmt.Printf("与%s断开!\n", addr)
}