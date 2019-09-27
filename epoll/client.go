package main

import (
	"net"
	"time"
)

func main() {
	for i := 0; i < 300; i++ {
		go client()
	}

	time.Sleep(time.Hour)
}

func client() {
	c, err := net.Dial("tcp4", "10.179.209.77:8888")
	if err != nil {
		panic(err)
	}
	var date [5]byte
	for {
		c.Write([]byte("1"))
		_, err := c.Read(date[:])
		if err != nil {
			panic(err)
		}
		//fmt.Println("response:", date[:n])
	}
}
