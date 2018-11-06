本片主要验证CityHash32、Murmur3、CRC32、FNV四种hash算法的散列均匀程度、计算速度。并不考虑计算CPU消耗

hash_mod_test.go: 验证散列情况，test结果如下

```shell
=== RUN   TestCityHash32
5 1249995
4 1249768
1 1249681
7 1253270
3 1249529
6 1249045
2 1248789
0 1249923
--- PASS: TestCityHash32 (2.54s)
=== RUN   TestMurmur3
4 1252014
7 1249278
5 1250732
0 1250782
6 1247427
2 1251659
3 1249871
1 1248237
--- PASS: TestMurmur3 (2.54s)
=== RUN   TestCRC32
3 1250000
0 1250000
6 1250000
4 1250000
2 1250000
1 1250000
7 1250000
5 1250000
--- PASS: TestCRC32 (3.53s)
=== RUN   TestFNV
2 1250004
3 1250004
0 1250000
1 1250000
6 1249996
7 1249996
4 1250000
5 1250000
--- PASS: TestFNV (2.69s)
PASS
ok  	command-line-arguments	11.305s
```

hash_time_test.go 验证计算速度，test结果如下

```shell
$ go test -v hash_time_test.go 
=== RUN   TestCityHash
--- PASS: TestCityHash (6.79s)
	hash_time_test.go:30: total time:224.754705ms
=== RUN   TestMurMur3
--- PASS: TestMurMur3 (5.23s)
	hash_time_test.go:39: total time:152.084599ms
=== RUN   TestCRC32
--- PASS: TestCRC32 (6.07s)
	hash_time_test.go:49: total time:524.058613ms
=== RUN   TestFNV
--- PASS: TestFNV (5.64s)
	hash_time_test.go:59: total time:465.081755ms
PASS
ok  	command-line-arguments	23.846s
```
