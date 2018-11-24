package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var randSeed = uint32(time.Now().Unix()) // 随机一个种子
func Fastrand() {
	fr := randSeed
	mx := uint32(int32(fr)>>31) & 0xa8888eef
	fr = fr<<1 ^ mx
	randSeed = fr
}

func MathRand() {
	rand.Int31()
}

func BenchmarkFastrand(b *testing.B) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ { // 准备2个线程并发随机
		wg.Add(1)
		go func() {
			<-start
			for i := 0; i < b.N; i++ {
				Fastrand()
			}
			wg.Done()
		}()
	}

	// 开始计时并且开始测试
	b.ResetTimer()
	close(start)
	// 等待结束
	wg.Wait()
}

func BenchmarkMathRand(b *testing.B) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ { // 准备2个线程并发随机
		wg.Add(1)
		go func() {
			<-start
			for i := 0; i < b.N; i++ {
				MathRand()
			}
			wg.Done()
		}()
	}

	// 开始计时并且开始测试
	b.ResetTimer()
	close(start)
	// 等待结束
	wg.Wait()
}
