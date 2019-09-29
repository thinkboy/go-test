package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
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

	go prof("0.0.0.0:8009")

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
		var body [200]byte
		n, err := conn.Read(body[:])
		if err != nil {
			break
		}
		//fmt.Printf("收到%s消息: %s\n", addr, string(body[:]))
		// 回包
		_, err = conn.Write(body[:n])
		if err != nil {
			break
		}
		atomic.AddInt64(&reqCount, 1)
		//fmt.Printf("发送给%s: %s\n", addr, string(body[:]))
	}
	fmt.Printf("与%s断开!\n", addr)
}

func prof(addr string) {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/debug/pprof/", pprof.Index)
	serveMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	serveMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serveMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serveMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(addr, serveMux); err != nil {
		fmt.Printf("Failed to StartTcp||err=%s\n", err)
	}
}
