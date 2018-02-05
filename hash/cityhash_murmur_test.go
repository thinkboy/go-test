package main

import (
	"fmt"
	"testing"

	"github.com/spaolacci/murmur3"
	"github.com/zhenjl/cityhash"
)

func TestCityHash(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		cityhash.CityHash32([]byte(str), uint32(len(str)))
	}
}

func TestMurMur3(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		murmur3.Sum32([]byte(str))
	}
}

