package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFastrand(t *testing.T) {
	randSeed := uint32(time.Now().Unix()) // 随机一个种子
	tmpMap := map[uint32]int{}

	for i := 0; i < 100000000; i++ { // 计算1亿次随机
		r := randSeed
		mx := uint32(int32(r)>>31) & 0xa8888eef
		r = r<<1 ^ mx
		randSeed = r

		tmpMap[r%8]++ // 统计出现的数字的次数
	}

	fmt.Println("Fastrand离散如下:")
	for n, c := range tmpMap {
		fmt.Println(n, c)
	}
}

func TestMathRand(t *testing.T) {
	myRand := rand.New(rand.NewSource(time.Now().Unix())) // 随机一个种子
	tmpMap := map[uint32]int{}

	for i := 0; i < 100000000; i++ { // 计算1亿次随机
		r := uint32(myRand.Int31())

		tmpMap[r%8]++ // 统计出现的数字的次数
	}

	fmt.Println("math/rand离散如下:")
	for n, c := range tmpMap {
		fmt.Println(n, c)
	}
}
