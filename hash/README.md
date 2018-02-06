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
cityhash_murmur_test.go 验证cityhash、murmur、crc32处理速度

```shell
$ go test -v cityhash_murmur_test.go 
=== RUN   TestCityHash
--- PASS: TestCityHash (5.42s)
	cityhash_murmur_test.go:28: total time:211.114801ms
=== RUN   TestMurMur3
--- PASS: TestMurMur3 (4.68s)
	cityhash_murmur_test.go:37: total time:159.55251ms
=== RUN   TestCRC32
--- PASS: TestCRC32 (5.24s)
	cityhash_murmur_test.go:47: total time:485.60166ms
PASS
ok  	command-line-arguments	15.458s
```