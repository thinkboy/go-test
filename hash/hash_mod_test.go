package main

import (
	"fmt"
	"hash/crc32"
	"hash/fnv"
	"testing"

	"github.com/spaolacci/murmur3"
	"github.com/zhenjl/cityhash"
)

func TestCityHash32(t *testing.T) {
	testMap := make(map[uint32]int)

	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		v := cityhash.CityHash32([]byte(str), uint32(len(str)))
		testMap[v%8]++ // 取模统计下每个模的分布
	}
	for n, c := range testMap { // 打出来看下分布情况
		fmt.Println(n, c)
	}
}

func TestMurmur3(t *testing.T) {
	testMap := make(map[uint32]int)

	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		v := murmur3.Sum32([]byte(str))
		testMap[v%8]++ // 取模统计下每个模的分布
	}
	for n, c := range testMap { // 打出来看下分布情况
		fmt.Println(n, c)
	}
}

func TestCRC32(t *testing.T) {
	testMap := make(map[uint32]int)

	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		hashH := crc32.NewIEEE()
		hashH.Write([]byte(str))
		v := hashH.Sum32()
		testMap[v%8]++ // 取模统计下每个模的分布
	}
	for n, c := range testMap { // 打出来看下分布情况
		fmt.Println(n, c)
	}
}

func TestFNV(t *testing.T) {
	testMap := make(map[uint32]int)

	for i := 0; i < 10000000; i++ {
		str := fmt.Sprintf("test-%09d", i)
		hashH := fnv.New32()
		hashH.Write([]byte(str))
		v := hashH.Sum32()
		testMap[v%8]++ // 取模统计下每个模的分布
	}
	for n, c := range testMap { // 打出来看下分布情况
		fmt.Println(n, c)
	}
}
