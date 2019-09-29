package main

/*
	启动命令例子：./client 50 10 127.0.0.1:8888
*/

import (
	"net"
	"os"
	"strconv"
	"time"
)

var content = "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789" // 100 bytes

func main() {
	// 连接数量
	connNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}
	// 单个连接发消息间隔，单位：毫秒
	intervalMill, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}

	for i := int64(0); i < connNum; i++ {
		go client(intervalMill)
		time.Sleep(10 * time.Millisecond)
	}

	time.Sleep(time.Hour)
}

func client(intervalMill int64) {
	c, err := net.Dial("tcp4", os.Args[3])
	if err != nil {
		panic(err)
	}
	var date [200]byte
	for {
		c.Write([]byte(content))
		_, err := c.Read(date[:])
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(intervalMill) * time.Millisecond)
		//fmt.Println("response:", date[:n])
	}
}
