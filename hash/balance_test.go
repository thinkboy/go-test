package main

import (
	"fmt"
	"testing"

	"github.com/zhenjl/cityhash"
)

func TestHash(t *testing.T) {
	coun := make(map[uint32]int)
	testMap := make(map[uint32]int)
	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		v := cityhash.CityHash32([]byte(str), uint32(len(str)))
		testMap[v/65536]++
	}
	for n, c := range testMap {
		coun[n/10000] += c
	}
	for n, c := range coun {
		fmt.Println(n, c)
	}
}

