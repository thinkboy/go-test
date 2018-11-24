本片主要验证CityHash32、Murmur3、CRC32、FNV四种hash算法的散列均匀程度、计算速度。并不考虑计算CPU、内存消耗

hash_mod_test.go: 验证散列情况，test结果如下

```shell
$ go test -v hash_mod_test.go 
=== RUN   TestCityHash32
4 1249768
1 1249681
7 1253270
3 1249529
6 1249045
2 1248789
0 1249923
5 1249995
--- PASS: TestCityHash32 (2.23s)
=== RUN   TestMurmur3
1 1248237
4 1252014
7 1249278
5 1250732
0 1250782
6 1247427
2 1251659
3 1249871
--- PASS: TestMurmur3 (2.21s)
=== RUN   TestCRC32
5 1250000
3 1250000
0 1250000
6 1250000
4 1250000
2 1250000
1 1250000
7 1250000
--- PASS: TestCRC32 (3.36s)
=== RUN   TestFNV
1 1250000
6 1249996
7 1249996
4 1250000
5 1250000
2 1250004
3 1250004
0 1250000
--- PASS: TestFNV (2.59s)
PASS
ok  	command-line-arguments	10.395s
```

hash_time_test.go 验证计算速度，test结果如下

```shell
$ go test -v hash_time_test.go 
=== RUN   TestCityHash
--- PASS: TestCityHash (6.51s)
	hash_time_test.go:31: total time:205.769065ms
=== RUN   TestMurMur3
--- PASS: TestMurMur3 (5.68s)
	hash_time_test.go:41: total time:150.69422ms
=== RUN   TestCRC32
--- PASS: TestCRC32 (6.52s)
	hash_time_test.go:54: total time:1.376267139s
=== RUN   TestFNV
--- PASS: TestFNV (6.20s)
	hash_time_test.go:67: total time:731.954417ms
PASS
ok  	command-line-arguments	25.030s
```

###结论

4种算法都散列效果都还可以，速度方面MurMur3更胜一筹
