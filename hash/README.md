balance_test.go: 验证用cityhash取模后数据是否均匀

```shell
$ go test -v balance_test.go 
=== RUN   TestHash
3 1527950
4 1524644
5 1526426
2 1525395
1 1525374
0 1525799
6 844412
--- PASS: TestHash (2.93s)
PASS
ok  	command-line-arguments	2.939s
```
cityhash_murmur_test.go 验证cityhash、murmur速度

```shell
$ go test -v cityhash_murmur_test.go 
=== RUN   TestCityHash
--- PASS: TestCityHash (2.00s)
=== RUN   TestMurMur3
--- PASS: TestMurMur3 (1.98s)
PASS
ok  	command-line-arguments	3.985s
```