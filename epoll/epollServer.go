package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
	"net/http/pprof"
)

const (
	EPOLLET = 1 << 31
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

type EPoll struct {
	epollFD int
	lnFd    int // listener的fd
	readBuf [200]byte
}

func NewEPoll() *EPoll {
	var (
		err error
	)
	epoll := &EPoll{}

	epoll.epollFD, err = syscall.EpollCreate1(0)
	if err != nil {
		panic(err)
	}
	return epoll
}

func (this *EPoll) AddRead(fd int) {
	if err := syscall.EpollCtl(this.epollFD, syscall.EPOLL_CTL_ADD, fd,
		&syscall.EpollEvent{Fd: int32(fd),
			Events: syscall.EPOLLIN,
		},
	); err != nil {
		panic(err)
	}

	this.lnFd = fd
}

func (this *EPoll) AddReadWrite(fd int) {
	if err := syscall.EpollCtl(this.epollFD, syscall.EPOLL_CTL_ADD, fd,
		&syscall.EpollEvent{Fd: int32(fd),
			Events: syscall.EPOLLIN | syscall.EPOLLOUT | EPOLLET,
		},
	); err != nil {
		panic(err)
	}
}

// 接收建立连接请求
func (this *EPoll) Wait() error {
	events := make([]syscall.EpollEvent, 64)
	for {
		//fmt.Println("epollwait begin", this.epollFD)
		n, err := syscall.EpollWait(this.epollFD, events, -1)
		if err != nil && err != syscall.EINTR {
			panic(err)
		}
		for i := 0; i < n; i++ {
			fd := int(events[i].Fd)
			if fd == this.lnFd { // 如果是一个建立连接事件
				if err := this.Accept(fd); err != nil {
					panic(err)
				}
				fmt.Println("accepted:", this.epollFD, "lnFD:", fd)
			} else { // 其余的都是链接的socket事件
				date := this.LoopRead(fd)
				if date == nil {
					continue
				}
				this.LoopWrite(fd, date)
				atomic.AddInt64(&reqCount, 1)
			}
		}
	}
}

// 从fd里读数据
func (this *EPoll) LoopRead(fd int) []byte {
	n, err := syscall.Read(fd, this.readBuf[:])
	if err != nil {
		if err == syscall.EAGAIN {
			return nil
		}
		panic(err)
	}
	if n == 0 { // socket被关闭了，也就是读到EOF
		syscall.Close(fd)
		fmt.Println(fmt.Sprintf("epollFD:%d socketFD:%d close", this.epollFD, fd))
		return nil
	}
	//fmt.Println("epollFD:", this.epollFD, "socketFD:", fd, "read:", a[:n])
	return this.readBuf[:n]
}

// 往fd里写数据
func (this *EPoll) LoopWrite(fd int, date []byte) {
	n, err := syscall.Write(fd, date)
	if err != nil {
		if err == syscall.EAGAIN {
			return
		}
		panic(err)
	}
	if n == 0 { // socket被关闭了，也就是读到EOF
		syscall.Close(fd)
		fmt.Println(fmt.Sprintf("epollFD:%d socketFD:%d close", this.epollFD, fd))
		return
	}
	//fmt.Println("epollFD:", this.epollFD, "socketFD:", fd, "write:", a[:n])
}

// 获取一个连接并加入到epoll里
func (this *EPoll) Accept(fd int) error {
	nfd, _, err := syscall.Accept(fd)
	if err != nil {
		if err == syscall.EAGAIN {
			return nil
		}
		return err
	}
	if err := syscall.SetNonblock(nfd, true); err != nil {
		return err
	}
	fmt.Println("accept connFD:", nfd)

	// 把socket fd加到epoll里
	this.AddReadWrite(nfd)
	return nil
}

func main() {
	// 开启QPS统计
	go loopCount()

	go prof("0.0.0.0:8009")

	epollNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	// 开启一个本地的监听，并拿到socket fd
	lnFD := createListen("tcp4", "0.0.0.0:8888")
	fmt.Println("lnFD:", lnFD)

	// 创建N个epoll
	for i := int64(0); i < epollNum; i++ {
		epoll := NewEPoll()
		epoll.AddRead(lnFD)
		go epoll.Wait()
	}

	// 卡死，进程不退出
	c := make(chan bool)
	<-c
}

// 本地创建个用于listen的fd
func createListen(network, address string) int {
	ln, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}
	l, ok := ln.(*net.TCPListener)
	if !ok {
		panic("BUG")
	}
	lnFile, err := l.File()
	if err != nil {
		panic(err)
	}
	lnFD := int(lnFile.Fd())
	syscall.SetNonblock(lnFD, true)

	return lnFD
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
