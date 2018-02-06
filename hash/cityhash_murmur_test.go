package main

import (
	"fmt"
	"hash/crc32"
	"testing"
	"time"

	"github.com/spaolacci/murmur3"
	"github.com/zhenjl/cityhash"
)

func Init() [][]byte {
	var strArray [][]byte
	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		strArray = append(strArray, []byte(str))
	}
	return strArray
}

func TestCityHash(t *testing.T) {
	array := Init()
	now := time.Now()
	for _, arr := range array {
		cityhash.CityHash32(arr, uint32(len(arr)))
	}
	t.Logf("total time:%s", time.Now().Sub(now).String())
}

func TestMurMur3(t *testing.T) {
	array := Init()
	now := time.Now()
	for _, arr := range array {
		murmur3.Sum32(arr)
	}
	t.Logf("total time:%s", time.Now().Sub(now).String())
}

func TestCRC32(t *testing.T) {
	array := Init()
	hashH := crc32.NewIEEE()
	now := time.Now()
	for _, arr := range array {
		hashH.Sum(arr)
	}
	t.Logf("total time:%s", time.Now().Sub(now).String())
}

