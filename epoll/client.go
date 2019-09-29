package main

import (
	"net"
	"strconv"
	"time"
	"os"
)

var content = "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"// 100 bytes

func main() {
	num, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < num; i++ {
		go client()
		time.Sleep(100*time.Millisecond)
	}

	time.Sleep(time.Hour)
}

func client() {
	c, err := net.Dial("tcp4", os.Args[2])
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
		//fmt.Println("response:", date[:n])
	}
}
