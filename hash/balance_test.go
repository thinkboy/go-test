package main

import (
	"fmt"
	"testing"

	"github.com/zhenjl/cityhash"
)

func TestHash(t *testing.T) {
	coun := make(map[int]int)
	testMap := make(map[uint32]int)
	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		v := cityhash.CityHash32([]byte(str), uint32(len(str)))
		testMap[v/65536]++
	}
	for n, c := range testMap {
		if n >= 0 && n < 10000 {
			coun[10000] += c
		} else if n >= 10000 && n < 20000 {
			coun[20000] += c
		} else if n >= 20000 && n < 30000 {
			coun[30000] += c
		} else if n >= 30000 && n < 40000 {
			coun[40000] += c
		} else if n >= 40000 && n < 50000 {
			coun[50000] += c
		} else if n >= 50000 && n < 60000 {
			coun[60000] += c
		} else if n >= 60000 && n < 70000 {
			coun[70000] += c
		}
	}
	for n, c := range coun {
		fmt.Println(n, c)
	}
}

