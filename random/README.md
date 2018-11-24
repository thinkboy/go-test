主要验证Fastrand和go自带的math/rand的随机算法的散列、计算速度。并不考虑计算CPU消耗

random_mod_test.go: 验证散列情况，test结果如下

```shell
$ go test -v random_mod_test.go 
=== RUN   TestFastrand
Fastrand离散如下:
3 12496901
6 12500717
0 12503554
7 12501752
1 12497935
2 12497935
4 12501751
5 12499455
--- PASS: TestFastrand (2.84s)
=== RUN   TestMathRand
math/rand离散如下:
4 12495015
7 12502722
1 12505881
3 12500407
5 12496796
0 12500899
2 12500294
6 12497986
--- PASS: TestMathRand (4.46s)
PASS
ok  	command-line-arguments	7.306s
```

random_time_test.go 验证2个goroutine下的计算速度，test结果如下

```shell
$ go test -test.bench=".*" random_time_test.go 
goos: darwin
goarch: amd64
BenchmarkFastrand-4   	500000000	         3.73 ns/op
BenchmarkMathRand-4   	20000000	        65.6 ns/op
PASS
ok  	command-line-arguments	3.600s
```

#### 总结

2种随机算法散列效果都还不错，速度方面Fastrand要快更多，并且go自带的math/rand有全局锁，因此并发越高处理越慢，而Fastrand的好处是无锁
