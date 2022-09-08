## Benchmark
### Environment
- [go-redis/redis](https://github.com/go-redis/redis) v8.11.5
- [joomcode/redispipe](https://github.com/joomcode/redispipe) v0.9.4
- [mediocregopher/radix](https://github.com/mediocregopher/radix) v4.1.1
- [rueian/rueidis](https://github.com/rueian/rueidis) v0.0.74
- [sandwich-go/redisson](https://github.com/sandwich-go/redisson) v1.1.14

### Redis Server
AWS ElasticCache for Redis，Cache type node `cache.m6g.large`.
- Single
- Cluster, 1 Master, 1 slave

Max connection pool size 100 and 1000.  
Length of Key 16.  
Length of Value 64,256 and 1024.  
Run benchmark using serial and parallel mode.

### Benchmarking Result
#### Single Redis
```markdown
+---------------------------------------------------+-----------+---------+------+-----------+
| Single Serial Get                                 | iteration | ns/op   | B/op | allocs/op |
+===================================================+===========+=========+======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4953      | 487019  | 276  |         6 |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4978      | 488346  | 276  |         6 |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4966      | 485089  | 484  |         6 |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4916      | 484548  | 484  |         6 |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4929      | 487335  | 1348 |         6 |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4933      | 488121  | 1348 |         6 |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4958      | 485398  | 320  |         4 |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4935      | 485190  | 320  |         4 |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4966      | 484629  | 512  |         4 |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4939      | 488952  | 512  |         4 |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4954      | 489221  | 1280 |         4 |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4926      | 491789  | 1280 |         4 |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4987      | 483422  | 256  |         4 |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4968      | 487068  | 256  |         4 |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4986      | 486305  | 448  |         4 |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4924      | 487030  | 448  |         4 |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4922      | 487140  | 1216 |         4 |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4966      | 487747  | 1216 |         4 |
| go-redis/redis/v8:Val(64):Pool(100)               | 4953      | 488424  | 276  |         6 |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4971      | 495383  | 276  |         6 |
| go-redis/redis/v8:Val(256):Pool(100)              | 4981      | 486368  | 484  |         6 |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4936      | 484809  | 484  |         6 |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4887      | 487234  | 1348 |         6 |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4928      | 496201  | 1348 |         6 |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4537      | 506352  | 876  |         4 |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4512      | 533221  | 8989 |         3 |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4537      | 510415  | 869  |         4 |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4526      | 532493  | 8787 |         3 |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4369      | 516728  | 903  |         4 |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4312      | 537812  | 9222 |         3 |
| joomcode/redispipe:Val(64):Pool(100)              | 1546      | 1557321 | 168  |         5 |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1565308 | 168  |         5 |
| joomcode/redispipe:Val(256):Pool(100)             | 1540      | 1573693 | 376  |         5 |
| joomcode/redispipe:Val(256):Pool(1000)            | 1526      | 1555364 | 376  |         5 |
| joomcode/redispipe:Val(1024):Pool(100)            | 1534      | 1561348 | 1240 |         5 |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1532      | 1561439 | 1240 |         5 |
+---------------------------------------------------+-----------+---------+------+-----------+
```

```markdown
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial Set                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+===========================================================================================================+===========+===============+===========+==============+
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4977      | 483504 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4923      | 485434 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4930      | 487251 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4968      | 487332 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4926      | 488972 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4897      | 486064 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4939      | 483098 ns/op  | 276 B/op  | 5 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4999      | 481544 ns/op  | 276 B/op  | 5 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4928      | 484242 ns/op  | 280 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4860      | 482453 ns/op  | 280 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4968      | 484771 ns/op  | 280 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4927      | 496207 ns/op  | 280 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4808      | 484878 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4986      | 486763 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4946      | 485818 ns/op  | 216 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4922      | 484952 ns/op  | 216 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4945      | 486969 ns/op  | 216 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4947      | 488603 ns/op  | 216 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4957      | 482288 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4899      | 487379 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4897      | 488155 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4876      | 485042 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4926      | 489551 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4962      | 489919 ns/op  | 264 B/op  | 7 allocs/op  |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4506      | 519044 ns/op  | 903 B/op  | 4 allocs/op  |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4232      | 547546 ns/op  | 9624 B/op | 39 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4315      | 522241 ns/op  | 946 B/op  | 4 allocs/op  |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4465      | 532153 ns/op  | 8998 B/op | 36 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4530      | 509190 ns/op  | 915 B/op  | 4 allocs/op  |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4512      | 534769 ns/op  | 9090 B/op | 36 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1548      | 1558676 ns/op | 114 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1540      | 1564027 ns/op | 114 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1540      | 1564637 ns/op | 114 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1527      | 1559428 ns/op | 114 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1537      | 1556612 ns/op | 115 B/op  | 6 allocs/op  |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1536      | 1567739 ns/op | 115 B/op  | 6 allocs/op  |
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial Del                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+===========================================================================================================+===========+===============+===========+==============+
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4960      | 491199 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4969      | 487374 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4960      | 484226 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4956      | 486280 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4965      | 488280 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4981      | 481978 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4953      | 482947 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4989      | 485699 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4952      | 482216 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4965      | 483253 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4914      | 493253 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4960      | 487147 ns/op  | 288 B/op  | 5 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 5024      | 482986 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4995      | 487758 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4846      | 482795 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 5001      | 482372 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4974      | 482227 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 5006      | 483856 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4936      | 488249 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4957      | 485687 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4912      | 486787 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4861      | 485094 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 5000      | 485747 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4932      | 484418 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4606      | 507947 ns/op  | 872 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4581      | 531004 ns/op  | 8862 B/op | 36 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4508      | 505590 ns/op  | 877 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4557      | 531726 ns/op  | 8738 B/op | 35 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4518      | 505961 ns/op  | 873 B/op  | 4 allocs/op  |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4596      | 527602 ns/op  | 8664 B/op | 35 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1568      | 1558597 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1537      | 1562103 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1537      | 1555240 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1540      | 1557134 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1539      | 1553183 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1532      | 1558317 ns/op | 64 B/op   | 3 allocs/op  |
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+--------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial Exists                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+==============================================================================================================+===========+===============+===========+==============+
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4958      | 484697 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4974      | 481462 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4950      | 483400 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4930      | 490863 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4958      | 487538 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4948      | 484440 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4966      | 484385 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4999      | 484252 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4946      | 485223 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4930      | 484162 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4978      | 481993 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4932      | 482680 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4954      | 491144 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4918      | 484869 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4968      | 486363 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4947      | 481133 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4986      | 482829 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4974      | 486667 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4920      | 484542 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4915      | 490139 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4956      | 486650 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4944      | 478302 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 5011      | 488393 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4963      | 479900 ns/op  | 196 B/op  | 6 allocs/op  |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4549      | 508087 ns/op  | 878 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4411      | 534180 ns/op  | 9196 B/op | 37 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4520      | 512916 ns/op  | 874 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4533      | 525759 ns/op  | 8777 B/op | 36 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4558      | 511653 ns/op  | 864 B/op  | 4 allocs/op  |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4494      | 527723 ns/op  | 8855 B/op | 36 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1530      | 1559326 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1542      | 1556280 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1536      | 1564361 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1540      | 1555467 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1542      | 1559493 ns/op | 64 B/op   | 3 allocs/op  |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1543      | 1559163 ns/op | 64 B/op   | 3 allocs/op  |
+--------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+--------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial Expire                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+==============================================================================================================+===========+===============+===========+==============+
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4849      | 483997 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4945      | 484310 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 5010      | 486427 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4995      | 488633 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4952      | 493152 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4996      | 487953 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4923      | 484355 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4980      | 484325 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4998      | 487745 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4972      | 482925 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4975      | 490133 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4963      | 484281 ns/op  | 240 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 5010      | 487674 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4954      | 487549 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4953      | 487863 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4956      | 486432 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4944      | 483093 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4971      | 484827 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4948      | 486768 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4911      | 484882 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4910      | 489564 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4963      | 485558 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4948      | 481531 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 5002      | 489293 ns/op  | 212 B/op  | 5 allocs/op  |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4508      | 502751 ns/op  | 905 B/op  | 4 allocs/op  |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4417      | 545005 ns/op  | 9205 B/op | 37 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4478      | 511238 ns/op  | 900 B/op  | 4 allocs/op  |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4557      | 546055 ns/op  | 8749 B/op | 35 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4486      | 512832 ns/op  | 896 B/op  | 4 allocs/op  |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4533      | 541219 ns/op  | 8794 B/op | 36 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1534      | 1564919 ns/op | 80 B/op   | 3 allocs/op  |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1534      | 1559231 ns/op | 80 B/op   | 3 allocs/op  |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1540      | 1558428 ns/op | 80 B/op   | 3 allocs/op  |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1464      | 1553472 ns/op | 80 B/op   | 3 allocs/op  |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1540      | 1556370 ns/op | 80 B/op   | 3 allocs/op  |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1537      | 1569699 ns/op | 80 B/op   | 3 allocs/op  |
+--------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial TTL                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+===========================================================================================================+===========+===============+===========+==============+
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4976      | 484425 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4957      | 485723 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4929      | 487783 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4981      | 481543 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 5007      | 483076 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4998      | 482210 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4934      | 486218 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4952      | 481857 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 5004      | 485067 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4988      | 481826 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4984      | 482331 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4914      | 481700 ns/op  | 256 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4958      | 484226 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4936      | 480418 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4988      | 490442 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4960      | 485770 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 5013      | 481530 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4986      | 485811 ns/op  | 184 B/op  | 3 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4972      | 485517 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4956      | 487745 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4946      | 482132 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4958      | 487410 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4989      | 489336 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4934      | 480211 ns/op  | 196 B/op  | 5 allocs/op  |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4586      | 518961 ns/op  | 863 B/op  | 4 allocs/op  |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4474      | 538017 ns/op  | 9070 B/op | 36 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4412      | 508648 ns/op  | 893 B/op  | 4 allocs/op  |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4404      | 525297 ns/op  | 9037 B/op | 37 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4431      | 509667 ns/op  | 889 B/op  | 4 allocs/op  |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4510      | 538578 ns/op  | 8825 B/op | 36 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1532      | 1562151 ns/op | 72 B/op   | 4 allocs/op  |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1543      | 1553243 ns/op | 72 B/op   | 4 allocs/op  |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1543      | 1562111 ns/op | 72 B/op   | 4 allocs/op  |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1540      | 1558208 ns/op | 72 B/op   | 4 allocs/op  |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1540      | 1560193 ns/op | 72 B/op   | 4 allocs/op  |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1538      | 1557668 ns/op | 72 B/op   | 4 allocs/op  |
+-----------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial HSet                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+============================================================================================================+===========+===============+===========+==============+
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4922      | 485689 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4970      | 487203 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4984      | 486250 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4920      | 484587 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4876      | 487357 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4922      | 489322 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4917      | 484549 ns/op  | 336 B/op  | 7 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4952      | 488471 ns/op  | 336 B/op  | 7 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4902      | 487053 ns/op  | 339 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4906      | 494815 ns/op  | 339 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4840      | 486261 ns/op  | 340 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4946      | 489426 ns/op  | 340 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4968      | 485816 ns/op  | 280 B/op  | 7 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4994      | 486246 ns/op  | 280 B/op  | 7 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4834      | 488234 ns/op  | 283 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4992      | 485545 ns/op  | 283 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4850      | 485980 ns/op  | 284 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4910      | 489439 ns/op  | 284 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4927      | 485680 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4957      | 491002 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4930      | 488967 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4918      | 486592 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4905      | 488907 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4941      | 486194 ns/op  | 276 B/op  | 8 allocs/op  |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4525      | 523867 ns/op  | 919 B/op  | 4 allocs/op  |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4422      | 530238 ns/op  | 9223 B/op | 37 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4452      | 517385 ns/op  | 932 B/op  | 4 allocs/op  |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4507      | 531133 ns/op  | 8929 B/op | 36 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4333      | 519020 ns/op  | 971 B/op  | 4 allocs/op  |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4188      | 537108 ns/op  | 9800 B/op | 39 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1533      | 1561166 ns/op | 128 B/op  | 5 allocs/op  |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1540      | 1554166 ns/op | 128 B/op  | 5 allocs/op  |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1540      | 1556652 ns/op | 128 B/op  | 5 allocs/op  |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1533      | 1558041 ns/op | 128 B/op  | 5 allocs/op  |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1533      | 1560643 ns/op | 129 B/op  | 5 allocs/op  |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1538      | 1561489 ns/op | 129 B/op  | 5 allocs/op  |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial HGet                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+============================================================================================================+===========+===============+===========+==============+
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 4852      | 501210 ns/op  | 308 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4935      | 484346 ns/op  | 308 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4962      | 487274 ns/op  | 516 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4989      | 489618 ns/op  | 516 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4910      | 488948 ns/op  | 1380 B/op | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4941      | 488380 ns/op  | 1380 B/op | 7 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4905      | 492828 ns/op  | 320 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4620      | 485737 ns/op  | 320 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 4894      | 486326 ns/op  | 512 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4976      | 485972 ns/op  | 512 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4916      | 487300 ns/op  | 1280 B/op | 4 allocs/op  |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4900      | 490358 ns/op  | 1280 B/op | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4921      | 483546 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4980      | 485752 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4962      | 487444 ns/op  | 448 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4887      | 490200 ns/op  | 448 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4870      | 487478 ns/op  | 1216 B/op | 4 allocs/op  |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4951      | 486974 ns/op  | 1216 B/op | 4 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4966      | 487846 ns/op  | 308 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4981      | 489512 ns/op  | 308 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4888      | 483954 ns/op  | 516 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4893      | 490433 ns/op  | 516 B/op  | 7 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4934      | 491366 ns/op  | 1380 B/op | 7 allocs/op  |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4909      | 488366 ns/op  | 1380 B/op | 7 allocs/op  |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4548      | 511410 ns/op  | 891 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4538      | 540596 ns/op  | 8954 B/op | 36 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4513      | 512147 ns/op  | 892 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4458      | 545670 ns/op  | 8939 B/op | 36 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4315      | 508803 ns/op  | 928 B/op  | 4 allocs/op  |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4521      | 533149 ns/op  | 8815 B/op | 35 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1558      | 1561483 ns/op | 200 B/op  | 6 allocs/op  |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1536      | 1561215 ns/op | 200 B/op  | 6 allocs/op  |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1536      | 1559134 ns/op | 408 B/op  | 6 allocs/op  |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1531      | 1557950 ns/op | 408 B/op  | 6 allocs/op  |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1531      | 1559150 ns/op | 1272 B/op | 6 allocs/op  |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1519      | 1563919 ns/op | 1272 B/op | 6 allocs/op  |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
| Single Serial HDel                                                                                         | iteration | ns/op         | B/op      | allocs/op    |
+============================================================================================================+===========+===============+===========+==============+
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100)      | 5004      | 485622 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000)     | 4939      | 479118 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100)     | 4953      | 486586 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000)    | 4951      | 482200 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100)    | 4984      | 482487 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000)   | 4825      | 485219 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100)      | 4954      | 485462 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000)     | 4982      | 486066 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100)     | 5006      | 481440 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000)    | 4962      | 484292 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100)    | 4963      | 484289 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000)   | 4932      | 481204 ns/op  | 256 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100)    | 4933      | 486686 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000)   | 4972      | 480096 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100)   | 4975      | 486593 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000)  | 4970      | 489757 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100)  | 4945      | 483823 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000) | 4995      | 482580 ns/op  | 200 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100)               | 4964      | 484106 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000)              | 4959      | 482416 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100)              | 4956      | 482754 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000)             | 4917      | 500200 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100)             | 4975      | 484676 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000)            | 4896      | 481217 ns/op  | 228 B/op  | 7 allocs/op  |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100)         | 4549      | 503538 ns/op  | 896 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000)        | 4596      | 546049 ns/op  | 8853 B/op | 35 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100)        | 4508      | 513116 ns/op  | 897 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000)       | 4544      | 536907 ns/op  | 8779 B/op | 35 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100)       | 4437      | 506036 ns/op  | 908 B/op  | 4 allocs/op  |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000)      | 4518      | 525352 ns/op  | 8829 B/op | 36 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100)              | 1538      | 1558144 ns/op | 96 B/op   | 4 allocs/op  |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000)             | 1538      | 1555319 ns/op | 96 B/op   | 4 allocs/op  |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100)             | 1538      | 1556155 ns/op | 96 B/op   | 4 allocs/op  |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000)            | 1534      | 1567664 ns/op | 96 B/op   | 4 allocs/op  |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100)            | 1537      | 1560043 ns/op | 96 B/op   | 4 allocs/op  |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000)           | 1540      | 1567453 ns/op | 96 B/op   | 4 allocs/op  |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+--------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
| Single Parallel Get                                                                                              | iteration | ns/op      | B/op      | allocs/op   |
+==================================================================================================================+===========+============+===========+=============+
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 362365    | 6136 ns/op | 279 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 504202    | 4731 ns/op | 286 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 362181    | 6334 ns/op | 487 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 481341    | 4946 ns/op | 495 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 332634    | 6822 ns/op | 1351 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 451609    | 5299 ns/op | 1360 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1208716   | 1923 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1263682   | 1870 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2013 ns/op | 512 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 986697    | 2180 ns/op | 512 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 728786    | 2816 ns/op | 1281 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 815188    | 2772 ns/op | 1281 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1253146   | 1847 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1300909   | 1874 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1000000   | 2034 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1000000   | 2091 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 792254    | 2686 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 775755    | 2779 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 369186    | 6098 ns/op | 279 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 506796    | 4750 ns/op | 286 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 357454    | 6266 ns/op | 487 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 486217    | 4919 ns/op | 495 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 331382    | 6779 ns/op | 1351 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 452067    | 5307 ns/op | 1360 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 596540    | 4284 ns/op | 26 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 589083    | 4902 ns/op | 54 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 576108    | 4384 ns/op | 27 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 597157    | 4993 ns/op | 54 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 573411    | 4539 ns/op | 27 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 559611    | 5062 ns/op | 56 B/op   | 1 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1109589   | 2137 ns/op | 168 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1000000   | 2129 ns/op | 169 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1000000   | 2170 ns/op | 377 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 999955    | 2188 ns/op | 377 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 958350    | 2442 ns/op | 1241 B/op | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 900044    | 2481 ns/op | 1241 B/op | 5 allocs/op |
+------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel Set                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+==================================================================================================================+===========+============+==========+=============+
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 363010    | 6189 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 462764    | 5175 ns/op | 275 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 355052    | 6228 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 442138    | 5378 ns/op | 276 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 353577    | 6293 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 428907    | 5608 ns/op | 276 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1000000   | 2057 ns/op | 276 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1000000   | 2015 ns/op | 276 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1205239   | 1978 ns/op | 279 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1191349   | 1981 ns/op | 279 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1128141   | 2095 ns/op | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1140604   | 2090 ns/op | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1235583   | 1916 ns/op | 212 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1245774   | 1871 ns/op | 212 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1254902   | 1901 ns/op | 215 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1000000   | 2013 ns/op | 215 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1159359   | 2095 ns/op | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1160755   | 2048 ns/op | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 364204    | 6139 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 465283    | 5168 ns/op | 275 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 356728    | 6214 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 446660    | 5381 ns/op | 276 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 352990    | 6262 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 430051    | 5570 ns/op | 276 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 575492    | 4293 ns/op | 43 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 576790    | 5200 ns/op | 73 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 565288    | 4328 ns/op | 43 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 563811    | 5171 ns/op | 72 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 585178    | 4411 ns/op | 43 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 541306    | 5242 ns/op | 75 B/op  | 1 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2072 ns/op | 115 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1235098   | 1940 ns/op | 115 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1355856   | 1978 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1429597   | 1983 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1208988   | 2020 ns/op | 134 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1246899   | 1923 ns/op | 132 B/op | 6 allocs/op |
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel Del                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+==================================================================================================================+===========+============+==========+=============+
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 356623    | 6106 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 527402    | 4529 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 366289    | 6140 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 529042    | 4551 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 362114    | 6083 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 529304    | 4518 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1342072   | 1757 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1225780   | 2011 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2058 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1000000   | 2072 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 984410    | 2055 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1218895   | 2027 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1228944   | 1898 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1215169   | 1898 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1238926   | 1915 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1270002   | 1924 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1226511   | 1925 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1211758   | 1856 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 359362    | 6063 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 532927    | 4550 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 369366    | 6031 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 527113    | 4504 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 359216    | 6040 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 530523    | 4497 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 596355    | 4252 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 645172    | 4854 ns/op | 53 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 612879    | 4248 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 628474    | 4831 ns/op | 53 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 508616    | 4265 ns/op | 28 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 624930    | 4816 ns/op | 53 B/op  | 1 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2165 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 982747    | 2095 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1000000   | 2153 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1000000   | 2122 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1000000   | 2110 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1000000   | 2151 ns/op | 64 B/op  | 3 allocs/op |
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+---------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel Exists                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+=====================================================================================================================+===========+============+==========+=============+
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 355356    | 6091 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 521980    | 4566 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 365894    | 6054 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 524054    | 4552 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 374433    | 6045 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 527689    | 4526 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1184643   | 1981 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1209036   | 1998 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1206621   | 2020 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1219750   | 2028 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1234136   | 1972 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1236686   | 1958 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1212618   | 1974 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1237340   | 1962 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1233102   | 1970 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1256658   | 2007 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1235353   | 2025 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1245318   | 1955 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 362736    | 6061 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 525018    | 4546 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 370290    | 6048 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 531529    | 4532 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 360061    | 6053 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 525180    | 4567 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 590653    | 4283 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 672496    | 4972 ns/op | 54 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 609844    | 4275 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 605505    | 4803 ns/op | 53 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 601099    | 4266 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 632251    | 4842 ns/op | 52 B/op  | 1 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2185 ns/op | 64 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1128975   | 2085 ns/op | 64 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1000000   | 2146 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1133168   | 2213 ns/op | 64 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1000000   | 2187 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1113512   | 2188 ns/op | 64 B/op  | 3 allocs/op |
+---------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+---------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel Expire                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+=====================================================================================================================+===========+============+==========+=============+
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 332355    | 6199 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 439668    | 5446 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 353678    | 6150 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 440203    | 5493 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 357690    | 6174 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 440194    | 5441 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 928232    | 2208 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1060755   | 2162 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2210 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 967376    | 2238 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 946177    | 2289 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 984056    | 2365 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1000000   | 2167 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2354 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1000000   | 2405 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1125740   | 2221 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1000000   | 2432 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 991659    | 2414 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 364699    | 6139 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 441435    | 5436 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 357466    | 6231 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 415408    | 5548 ns/op | 225 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 363788    | 6138 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 439591    | 5446 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 566857    | 4225 ns/op | 43 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 599179    | 5271 ns/op | 75 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 630436    | 4215 ns/op | 41 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 545533    | 5173 ns/op | 74 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 575806    | 4233 ns/op | 42 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 574726    | 5222 ns/op | 74 B/op  | 1 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 834244    | 2574 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 767469    | 2607 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 863950    | 2601 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 898284    | 2591 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 848350    | 2592 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 865720    | 2549 ns/op | 81 B/op  | 3 allocs/op |
+---------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel TTL                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+==================================================================================================================+===========+============+==========+=============+
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 354982    | 6075 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 503134    | 4750 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 365344    | 6065 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 503127    | 4740 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 366339    | 6101 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 502210    | 4760 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1341606   | 1876 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1258068   | 1915 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1200594   | 1875 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1247716   | 1878 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1283361   | 1883 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1204140   | 1947 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1305181   | 1891 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1310246   | 1892 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1259728   | 1913 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1290460   | 1851 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1283109   | 1792 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1298436   | 1889 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 366732    | 6067 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 505086    | 4748 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 367401    | 6028 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 494394    | 4890 ns/op | 207 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 361278    | 6096 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 506158    | 4755 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 585817    | 4270 ns/op | 27 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 545426    | 4779 ns/op | 55 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 610243    | 4251 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 561709    | 4851 ns/op | 54 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 632330    | 4226 ns/op | 26 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 584240    | 4805 ns/op | 55 B/op  | 1 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2192 ns/op | 72 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 954081    | 2201 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1000000   | 2170 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1119945   | 2214 ns/op | 72 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1051593   | 2263 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1000000   | 2181 ns/op | 73 B/op  | 4 allocs/op |
+------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel HSet                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+===================================================================================================================+===========+============+==========+=============+
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 353491    | 6200 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 463742    | 5195 ns/op | 287 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 351422    | 6251 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 446103    | 5350 ns/op | 288 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 350900    | 6418 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 399879    | 5643 ns/op | 289 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 910920    | 2214 ns/op | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 971934    | 2147 ns/op | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1246284   | 2083 ns/op | 339 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1239913   | 2046 ns/op | 339 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1116276   | 2146 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 958636    | 2166 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1218153   | 1953 ns/op | 280 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1214400   | 2089 ns/op | 280 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1230264   | 1962 ns/op | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1247301   | 1909 ns/op | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1139823   | 2136 ns/op | 284 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1122312   | 2141 ns/op | 284 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 351799    | 6188 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 462050    | 5170 ns/op | 287 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 356505    | 6564 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 433468    | 5468 ns/op | 288 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 349581    | 6268 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 427269    | 5660 ns/op | 288 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 570483    | 4336 ns/op | 59 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 613366    | 5156 ns/op | 88 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 564765    | 4397 ns/op | 59 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 578872    | 5339 ns/op | 90 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 531266    | 4459 ns/op | 60 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 555976    | 5264 ns/op | 90 B/op  | 1 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2131 ns/op | 129 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1111388   | 2208 ns/op | 129 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1380260   | 1688 ns/op | 131 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1210473   | 1790 ns/op | 130 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1215700   | 1963 ns/op | 149 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1000000   | 2037 ns/op | 151 B/op | 5 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
| Single Parallel HGet                                                                                              | iteration | ns/op      | B/op      | allocs/op   |
+===================================================================================================================+===========+============+===========+=============+
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 354025    | 6300 ns/op | 311 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 463798    | 5210 ns/op | 320 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 351859    | 6380 ns/op | 519 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 472824    | 5059 ns/op | 527 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 326384    | 6987 ns/op | 1383 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 434271    | 5425 ns/op | 1392 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1110710   | 2235 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 928776    | 2240 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 940675    | 2402 ns/op | 512 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1000000   | 2304 ns/op | 512 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 785270    | 2829 ns/op | 1281 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 720984    | 2796 ns/op | 1281 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 965226    | 2091 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2143 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 909498    | 2309 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1081938   | 2244 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 861998    | 2790 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 853593    | 2790 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 365218    | 6171 ns/op | 311 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 486632    | 5064 ns/op | 319 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 349214    | 6398 ns/op | 519 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 469039    | 5075 ns/op | 527 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 331870    | 6819 ns/op | 1383 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 440664    | 5458 ns/op | 1392 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 583803    | 4315 ns/op | 43 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 632470    | 4945 ns/op | 71 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 661624    | 4388 ns/op | 41 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 641383    | 5014 ns/op | 70 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 597493    | 4558 ns/op | 42 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 582472    | 5167 ns/op | 72 B/op   | 1 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 924786    | 2202 ns/op | 201 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1000000   | 2181 ns/op | 201 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1000000   | 2149 ns/op | 409 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1000000   | 2172 ns/op | 409 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 950053    | 2523 ns/op | 1273 B/op | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 828849    | 2565 ns/op | 1273 B/op | 6 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Single Parallel HDel                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+===================================================================================================================+===========+============+==========+=============+
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 358890    | 6126 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 520810    | 4607 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 357205    | 6100 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 507561    | 4626 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 362708    | 6104 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 519915    | 4616 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1000000   | 2063 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1211716   | 2042 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2081 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1200013   | 2010 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1000000   | 2085 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1000000   | 2073 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1000000   | 2009 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2081 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1000000   | 2007 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1238990   | 2002 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1000000   | 2118 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1000000   | 2046 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(100):Parallel(128)-4               | 361191    | 6105 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 507292    | 4753 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(100):Parallel(128)-4              | 361250    | 6112 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 520888    | 4604 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 364586    | 6070 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 513022    | 4623 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(100):Parallel(128)-4         | 605346    | 4328 ns/op | 42 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 610455    | 4834 ns/op | 69 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(100):Parallel(128)-4        | 575258    | 4321 ns/op | 43 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 638984    | 4868 ns/op | 69 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 620847    | 4379 ns/op | 42 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 594288    | 4893 ns/op | 70 B/op  | 1 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1182980   | 2212 ns/op | 96 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 994881    | 2047 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1210677   | 2046 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1153348   | 1976 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1000000   | 2075 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Single:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1199133   | 1984 ns/op | 96 B/op  | 4 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

#### Cluster Redis
```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
| Cluster Serial Get                                                                                         | iteration | ns/op         | B/op      | allocs/op   |
+============================================================================================================+===========+===============+===========+=============+
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4890      | 487417 ns/op  | 276 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4915      | 488523 ns/op  | 276 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4964      | 491417 ns/op  | 484 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4939      | 484393 ns/op  | 484 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4914      | 488911 ns/op  | 1348 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4900      | 488106 ns/op  | 1348 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4938      | 488312 ns/op  | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4958      | 487591 ns/op  | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4960      | 485184 ns/op  | 512 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4983      | 487830 ns/op  | 512 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4928      | 488460 ns/op  | 1280 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4878      | 489054 ns/op  | 1280 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 5013      | 486590 ns/op  | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4929      | 489007 ns/op  | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4918      | 485077 ns/op  | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4926      | 485526 ns/op  | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4899      | 490463 ns/op  | 1216 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4939      | 491525 ns/op  | 1216 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4861      | 489484 ns/op  | 276 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4958      | 484996 ns/op  | 276 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4929      | 488630 ns/op  | 484 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4863      | 486737 ns/op  | 484 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4876      | 497180 ns/op  | 1348 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4872      | 494822 ns/op  | 1348 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4729      | 517134 ns/op  | 118 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4700      | 521479 ns/op  | 175 B/op  | 3 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4725      | 506786 ns/op  | 118 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4663      | 529271 ns/op  | 176 B/op  | 3 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4563      | 512028 ns/op  | 118 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4707      | 523469 ns/op  | 175 B/op  | 3 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1538      | 1560667 ns/op | 169 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1536      | 1560187 ns/op | 169 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1533      | 1563273 ns/op | 377 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1540      | 1559477 ns/op | 377 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1533      | 1576437 ns/op | 1241 B/op | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1534      | 1562195 ns/op | 1242 B/op | 5 allocs/op |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial Set                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4906      | 493421 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4958      | 489022 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4881      | 487076 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4879      | 494568 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4892      | 493964 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4861      | 493905 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4915      | 489125 ns/op  | 342 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4898      | 489537 ns/op  | 276 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4933      | 486003 ns/op  | 345 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4933      | 491336 ns/op  | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4860      | 494323 ns/op  | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4892      | 493207 ns/op  | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4927      | 486205 ns/op  | 212 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4935      | 490395 ns/op  | 277 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4960      | 491011 ns/op  | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4945      | 486436 ns/op  | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4914      | 488440 ns/op  | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4784      | 489360 ns/op  | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4870      | 496824 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4898      | 493131 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4876      | 500501 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4875      | 487160 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4855      | 497155 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4845      | 496502 ns/op  | 264 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4707      | 520846 ns/op  | 136 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4606      | 545805 ns/op  | 212 B/op | 3 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4518      | 512600 ns/op  | 140 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 3937      | 533491 ns/op  | 264 B/op | 3 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4635      | 509365 ns/op  | 159 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4528      | 548624 ns/op  | 448 B/op | 3 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1564      | 1559278 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1532      | 1564312 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1527      | 1564032 ns/op | 119 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1536      | 1562740 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1530      | 1567822 ns/op | 117 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1530      | 1568534 ns/op | 117 B/op | 6 allocs/op |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial Del                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4954      | 482505 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4909      | 490404 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4972      | 489540 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4965      | 484667 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4951      | 483778 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4947      | 486319 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4915      | 485578 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4992      | 483628 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4927      | 486778 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4989      | 483615 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4966      | 486536 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4923      | 481059 ns/op  | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4885      | 486639 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4980      | 483367 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 5012      | 488939 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4950      | 482863 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4968      | 482569 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4848      | 484740 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4888      | 494681 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4899      | 483609 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4956      | 488602 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4934      | 487820 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4910      | 490078 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4882      | 486974 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4572      | 516742 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4581      | 533192 ns/op  | 176 B/op | 3 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4382      | 524153 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4647      | 515357 ns/op  | 174 B/op | 3 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4366      | 525362 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4699      | 517601 ns/op  | 176 B/op | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1549      | 1557796 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1540      | 1557442 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1537      | 1559602 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1540      | 1554963 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1539      | 1571336 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1542      | 1552898 ns/op | 66 B/op  | 3 allocs/op |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+---------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial Exists                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+===============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4948      | 486207 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4945      | 483354 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4980      | 484337 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4996      | 492530 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4989      | 484051 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4928      | 489495 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4976      | 481683 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4958      | 483620 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4990      | 484575 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4978      | 484997 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4920      | 486551 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4952      | 487337 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4978      | 481569 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4996      | 486897 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4992      | 483937 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4963      | 488358 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4953      | 483852 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4945      | 484625 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4926      | 491615 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4908      | 486800 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4912      | 488241 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4918      | 509994 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4884      | 485003 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4978      | 486328 ns/op  | 196 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4714      | 514884 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4471      | 528009 ns/op  | 176 B/op | 3 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4380      | 507985 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4460      | 542699 ns/op  | 178 B/op | 3 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4792      | 508285 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4303      | 517149 ns/op  | 174 B/op | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1545      | 1559345 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1537      | 1556615 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1544      | 1583717 ns/op | 66 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1536      | 1557796 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1543      | 1559059 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1533      | 1559116 ns/op | 65 B/op  | 3 allocs/op |
+---------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+---------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
| Cluster Serial Expire                                                                                         | iteration | ns/op         | B/op      | allocs/op   |
+===============================================================================================================+===========+===============+===========+=============+
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4939      | 490287 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4898      | 490532 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4868      | 488661 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4870      | 487149 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4879      | 492483 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4934      | 489043 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4960      | 487726 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4890      | 488136 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4909      | 498026 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4893      | 512902 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4914      | 492731 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4892      | 491587 ns/op  | 240 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4940      | 489265 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4921      | 484519 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4894      | 493868 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4920      | 489021 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4936      | 488031 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4934      | 484993 ns/op  | 184 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4863      | 494336 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4887      | 490703 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4845      | 501066 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4850      | 492617 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4904      | 493192 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4854      | 487685 ns/op  | 212 B/op  | 5 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4663      | 529580 ns/op  | 1296 B/op | 4 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4604      | 542000 ns/op  | 194 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4531      | 515191 ns/op  | 134 B/op  | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4590      | 536272 ns/op  | 193 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4308      | 519950 ns/op  | 134 B/op  | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4333      | 531320 ns/op  | 193 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1533      | 1564005 ns/op | 81 B/op   | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1536      | 1572816 ns/op | 81 B/op   | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1531      | 1565808 ns/op | 81 B/op   | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1528      | 1567128 ns/op | 81 B/op   | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1530      | 1569485 ns/op | 82 B/op   | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1532      | 1563870 ns/op | 82 B/op   | 3 allocs/op |
+---------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
```

```markdown
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial TTL                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4986      | 485173 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4983      | 487435 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4971      | 486494 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4951      | 485958 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4977      | 486930 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4965      | 485725 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4988      | 482261 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4963      | 482398 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4976      | 484914 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4550      | 481445 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4994      | 485167 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4988      | 480536 ns/op  | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4981      | 480642 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4958      | 487868 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4965      | 480767 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 5011      | 484296 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4977      | 483220 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4971      | 489272 ns/op  | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4956      | 486490 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4953      | 489336 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4921      | 487203 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4896      | 490285 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4909      | 489089 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4899      | 484813 ns/op  | 196 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4392      | 516994 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4162      | 556218 ns/op  | 180 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4772      | 503889 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4322      | 525965 ns/op  | 176 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4254      | 506998 ns/op  | 118 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4707      | 531403 ns/op  | 176 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1569      | 1550010 ns/op | 74 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1497      | 1574281 ns/op | 74 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1533      | 1562076 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1540      | 1555787 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1537      | 1556139 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1538      | 1561719 ns/op | 73 B/op  | 4 allocs/op |
+------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial HSet                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+=============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4882      | 490112 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4894      | 492661 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4893      | 489208 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4897      | 492833 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4894      | 501095 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4898      | 493391 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4939      | 492085 ns/op  | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4915      | 494991 ns/op  | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4918      | 486724 ns/op  | 339 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4898      | 489289 ns/op  | 339 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4868      | 492934 ns/op  | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4898      | 489974 ns/op  | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4922      | 495294 ns/op  | 345 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4759      | 490248 ns/op  | 348 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4770      | 487892 ns/op  | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4927      | 489624 ns/op  | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4884      | 494230 ns/op  | 284 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4854      | 494940 ns/op  | 350 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4828      | 492591 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4860      | 492739 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4928      | 497156 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4886      | 490350 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4856      | 506016 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4818      | 497876 ns/op  | 276 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4711      | 516181 ns/op  | 152 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4293      | 525695 ns/op  | 227 B/op | 3 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4612      | 521138 ns/op  | 156 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4635      | 525211 ns/op  | 270 B/op | 3 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4178      | 531872 ns/op  | 177 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4273      | 543087 ns/op  | 475 B/op | 3 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1528      | 1562925 ns/op | 130 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1525      | 1563617 ns/op | 130 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1536      | 1558767 ns/op | 133 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1534      | 1570781 ns/op | 130 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1533      | 1578202 ns/op | 131 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1533      | 1570661 ns/op | 131 B/op | 5 allocs/op |
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
| Cluster Serial HGet                                                                                         | iteration | ns/op         | B/op      | allocs/op   |
+=============================================================================================================+===========+===============+===========+=============+
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4922      | 486146 ns/op  | 308 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4922      | 488823 ns/op  | 308 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4920      | 487590 ns/op  | 516 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4942      | 490708 ns/op  | 516 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4910      | 490163 ns/op  | 1380 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4906      | 488618 ns/op  | 1380 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4980      | 487716 ns/op  | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4942      | 491867 ns/op  | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4894      | 489296 ns/op  | 512 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4902      | 490110 ns/op  | 512 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4876      | 484093 ns/op  | 1280 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4924      | 488573 ns/op  | 1280 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4981      | 483550 ns/op  | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4946      | 489059 ns/op  | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4957      | 483505 ns/op  | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 4920      | 490375 ns/op  | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 4945      | 486025 ns/op  | 1216 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 4915      | 491819 ns/op  | 1216 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4938      | 486744 ns/op  | 308 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4887      | 488740 ns/op  | 308 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4840      | 491426 ns/op  | 516 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4894      | 498615 ns/op  | 516 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4921      | 493442 ns/op  | 1380 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4929      | 489069 ns/op  | 1380 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4664      | 527804 ns/op  | 134 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4432      | 541879 ns/op  | 193 B/op  | 3 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4320      | 512200 ns/op  | 134 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4291      | 520926 ns/op  | 192 B/op  | 3 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4744      | 507930 ns/op  | 1264 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4654      | 552060 ns/op  | 196 B/op  | 3 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1555      | 1558220 ns/op | 201 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1540      | 1563677 ns/op | 201 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1533      | 1563890 ns/op | 409 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1532      | 1556235 ns/op | 409 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1534      | 1567017 ns/op | 1274 B/op | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1531      | 1561188 ns/op | 1273 B/op | 6 allocs/op |
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+-----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
| Cluster Serial HDel                                                                                         | iteration | ns/op         | B/op     | allocs/op   |
+=============================================================================================================+===========+===============+==========+=============+
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100)      | 4929      | 487506 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000)     | 4665      | 486410 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100)     | 4964      | 486650 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000)    | 4972      | 483611 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100)    | 4966      | 489641 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000)   | 4958      | 485558 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100)      | 4800      | 491455 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000)     | 4951      | 479973 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100)     | 4956      | 483460 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000)    | 4996      | 485347 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100)    | 4981      | 485060 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000)   | 4986      | 487348 ns/op  | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100)    | 4888      | 483396 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000)   | 4968      | 483600 ns/op  | 265 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100)   | 4976      | 488035 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000)  | 5008      | 481217 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100)  | 5020      | 485092 ns/op  | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000) | 5013      | 483332 ns/op  | 264 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100)               | 4900      | 488946 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000)              | 4950      | 513867 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100)              | 4910      | 487995 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000)             | 4875      | 484749 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100)             | 4922      | 489232 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000)            | 4942      | 489995 ns/op  | 228 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100)         | 4756      | 519971 ns/op  | 134 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000)        | 4586      | 513765 ns/op  | 190 B/op | 3 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100)        | 4449      | 528388 ns/op  | 134 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000)       | 4317      | 515906 ns/op  | 191 B/op | 3 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100)       | 4656      | 522122 ns/op  | 134 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000)      | 4429      | 514812 ns/op  | 191 B/op | 3 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100)              | 1539      | 1555304 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000)             | 1539      | 1557930 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100)             | 1539      | 1572486 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000)            | 1405      | 1554021 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100)            | 1538      | 1555258 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000)           | 1536      | 1557521 ns/op | 98 B/op  | 4 allocs/op |
+-------------------------------------------------------------------------------------------------------------+-----------+---------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
| Cluster Parallel Get                                                                                              | iteration | ns/op      | B/op      | allocs/op   |
+===================================================================================================================+===========+============+===========+=============+
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 361689    | 6246 ns/op | 279 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 494625    | 4819 ns/op | 286 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 353413    | 6439 ns/op | 487 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 478305    | 5035 ns/op | 494 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 324940    | 6992 ns/op | 1351 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 441291    | 5472 ns/op | 1360 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1036126   | 2275 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1067506   | 2250 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1008175   | 2420 ns/op | 513 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 883516    | 2394 ns/op | 512 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 766168    | 2906 ns/op | 1282 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 729106    | 2878 ns/op | 1282 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 946216    | 2266 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2216 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 924811    | 2292 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 969163    | 2306 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 856582    | 2802 ns/op | 1218 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 814442    | 2819 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 351850    | 6251 ns/op | 279 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 489259    | 4821 ns/op | 286 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 356703    | 6385 ns/op | 487 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 478236    | 5012 ns/op | 494 B/op  | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 333362    | 6972 ns/op | 1351 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 443264    | 5386 ns/op | 1360 B/op | 6 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 477573    | 4598 ns/op | 113 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 386779    | 5431 ns/op | 114 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 459818    | 4737 ns/op | 113 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 383200    | 5656 ns/op | 114 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 451070    | 4911 ns/op | 114 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 356745    | 5745 ns/op | 114 B/op  | 2 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1091751   | 2147 ns/op | 170 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1116862   | 2133 ns/op | 170 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1088572   | 2298 ns/op | 379 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1073509   | 2189 ns/op | 379 B/op  | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 800530    | 2548 ns/op | 1246 B/op | 5 allocs/op |
| BenchmarkClientGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 934084    | 2595 ns/op | 1246 B/op | 5 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel Set                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+===================================================================================================================+===========+============+==========+=============+
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 349974    | 6494 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 404748    | 5936 ns/op | 277 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 337305    | 6662 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 324763    | 6220 ns/op | 280 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 323359    | 6979 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 299281    | 6823 ns/op | 281 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 946690    | 2327 ns/op | 276 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 927264    | 2324 ns/op | 276 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 814446    | 2518 ns/op | 279 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 820002    | 2476 ns/op | 279 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 621984    | 3377 ns/op | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 662259    | 3376 ns/op | 280 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 939330    | 2276 ns/op | 212 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 888968    | 2264 ns/op | 212 B/op | 5 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 874808    | 2477 ns/op | 215 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 837128    | 2462 ns/op | 215 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 656972    | 3348 ns/op | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 639235    | 3363 ns/op | 216 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 336496    | 6471 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 339688    | 5931 ns/op | 279 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 342784    | 6586 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 384579    | 6300 ns/op | 277 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 318220    | 7096 ns/op | 267 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 300900    | 6820 ns/op | 281 B/op | 7 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 487761    | 4476 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 387138    | 5689 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 547970    | 4467 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 370246    | 5938 ns/op | 131 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 495886    | 4527 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 345834    | 6434 ns/op | 133 B/op | 2 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 981516    | 2269 ns/op | 115 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 933568    | 2328 ns/op | 115 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1018083   | 2366 ns/op | 117 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 928414    | 2350 ns/op | 116 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 654692    | 3278 ns/op | 139 B/op | 6 allocs/op |
| BenchmarkClientSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 678476    | 3260 ns/op | 134 B/op | 6 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel Del                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+===================================================================================================================+===========+============+==========+=============+
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 358837    | 6260 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 497517    | 4832 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 350084    | 6197 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 493366    | 4828 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 363262    | 6290 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 501684    | 4814 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 993585    | 2125 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1137520   | 2114 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 985495    | 2119 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1000000   | 2146 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 987890    | 2216 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1000000   | 2274 ns/op | 288 B/op | 5 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 945025    | 2213 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2179 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1123978   | 2176 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1111022   | 2169 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1104888   | 2157 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1000000   | 2150 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 364538    | 6146 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 511624    | 4701 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 370584    | 6229 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 514002    | 4622 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 368612    | 6152 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 514106    | 4768 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 471505    | 4619 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 394359    | 5503 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 467673    | 4604 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 389193    | 5547 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 484116    | 4625 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 399212    | 5503 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2101 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 941641    | 2151 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1140456   | 2041 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1181277   | 2073 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1180410   | 2071 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1000000   | 2036 ns/op | 65 B/op  | 3 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+----------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel Exists                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+======================================================================================================================+===========+============+==========+=============+
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 356186    | 6209 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 501278    | 4735 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 360022    | 6176 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 508154    | 4783 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 359449    | 6184 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 507560    | 4717 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 931364    | 2190 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 939031    | 2162 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 972061    | 2169 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 919816    | 2222 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1000000   | 2201 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 994989    | 2221 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1106865   | 2175 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2169 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 935419    | 2221 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1069153   | 2119 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1000000   | 2153 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1101631   | 2160 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 360195    | 6205 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 511290    | 4669 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 360906    | 6171 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 516358    | 4667 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 359160    | 6164 ns/op | 199 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 516640    | 4687 ns/op | 206 B/op | 6 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 477600    | 4610 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 465466    | 5489 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 478651    | 4640 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 385738    | 5520 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 474693    | 4699 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 381096    | 5538 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1180543   | 2046 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 986827    | 2100 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1204221   | 2064 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1000000   | 2058 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1177262   | 1998 ns/op | 65 B/op  | 3 allocs/op |
| BenchmarkClientExists/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1210532   | 2071 ns/op | 65 B/op  | 3 allocs/op |
+----------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+----------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel Expire                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+======================================================================================================================+===========+============+==========+=============+
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 336084    | 6540 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 391800    | 6132 ns/op | 225 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 334539    | 6553 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 391501    | 6132 ns/op | 225 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 333081    | 6559 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 391326    | 6168 ns/op | 225 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 792888    | 2812 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 668226    | 2996 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 714886    | 2977 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 691064    | 2906 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 792148    | 3052 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 732633    | 2942 ns/op | 240 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 752437    | 3002 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 719923    | 2945 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 739682    | 2749 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 809517    | 2800 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 689179    | 3083 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 858776    | 2785 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 342097    | 6526 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 329268    | 6150 ns/op | 227 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 345883    | 6516 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 331440    | 6168 ns/op | 227 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 335524    | 6554 ns/op | 215 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 328159    | 6153 ns/op | 227 B/op | 5 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 495571    | 4423 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 365419    | 5965 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 503380    | 4353 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 378637    | 5840 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 502240    | 4421 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 366000    | 5830 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 836614    | 2915 ns/op | 81 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 821937    | 2908 ns/op | 82 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 718962    | 2864 ns/op | 82 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 818583    | 2888 ns/op | 82 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 737113    | 2878 ns/op | 82 B/op  | 3 allocs/op |
| BenchmarkClientExpire/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 754645    | 2848 ns/op | 82 B/op  | 3 allocs/op |
+----------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel TTL                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+===================================================================================================================+===========+============+==========+=============+
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 359133    | 6306 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 498634    | 4800 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 358323    | 6203 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 492385    | 5047 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 356343    | 6209 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 501309    | 4797 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1279167   | 1906 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1265374   | 1836 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2022 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1200952   | 1857 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1259340   | 1892 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1246321   | 1950 ns/op | 256 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1335216   | 1821 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1281430   | 1934 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1308994   | 1771 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1357956   | 1896 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1292203   | 1777 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1306335   | 1805 ns/op | 184 B/op | 3 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 360853    | 6166 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 499804    | 4811 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 360001    | 6177 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 497072    | 4797 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 366268    | 6249 ns/op | 199 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 500016    | 4806 ns/op | 206 B/op | 5 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 488238    | 4596 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 388599    | 5569 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 459261    | 4610 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 466982    | 5558 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 474943    | 4618 ns/op | 113 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 387424    | 5465 ns/op | 114 B/op | 2 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1000000   | 2138 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1121539   | 2161 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 999075    | 2116 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 987087    | 2186 ns/op | 74 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 997465    | 2101 ns/op | 73 B/op  | 4 allocs/op |
| BenchmarkClientTTL/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1000000   | 2104 ns/op | 73 B/op  | 4 allocs/op |
+-------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel HSet                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+====================================================================================================================+===========+============+==========+=============+
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 344802    | 6604 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 392251    | 6089 ns/op | 289 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 335762    | 6712 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 313497    | 6391 ns/op | 292 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 312363    | 7201 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 289394    | 7092 ns/op | 293 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 953574    | 2484 ns/op | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 837054    | 2635 ns/op | 336 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 761954    | 2690 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 813997    | 2686 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 571730    | 3632 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 619077    | 3599 ns/op | 340 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 867099    | 2418 ns/op | 280 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 981193    | 2418 ns/op | 280 B/op | 7 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 792901    | 2668 ns/op | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 834783    | 2667 ns/op | 283 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 596102    | 3606 ns/op | 284 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 602301    | 3586 ns/op | 284 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 344406    | 6586 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 393264    | 6158 ns/op | 289 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 359312    | 6700 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 315304    | 6400 ns/op | 292 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 305941    | 7163 ns/op | 279 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 290716    | 7080 ns/op | 293 B/op | 8 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 495270    | 4440 ns/op | 145 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 364306    | 5865 ns/op | 146 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 488095    | 4542 ns/op | 145 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 550398    | 6096 ns/op | 146 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 467059    | 4760 ns/op | 146 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 327129    | 6858 ns/op | 150 B/op | 2 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1019394   | 2344 ns/op | 129 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1004547   | 2343 ns/op | 129 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 857570    | 2537 ns/op | 132 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 864350    | 2527 ns/op | 129 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 637536    | 3487 ns/op | 152 B/op | 5 allocs/op |
| BenchmarkClientHSet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 628860    | 3477 ns/op | 153 B/op | 5 allocs/op |
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```

```markdown
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
| Cluster Parallel HGet                                                                                              | iteration | ns/op      | B/op      | allocs/op   |
+====================================================================================================================+===========+============+===========+=============+
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 369856    | 6348 ns/op | 311 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 476847    | 5070 ns/op | 318 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 366747    | 6478 ns/op | 519 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 463682    | 5152 ns/op | 527 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 331086    | 7036 ns/op | 1383 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 429591    | 5580 ns/op | 1392 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 959172    | 2170 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1124628   | 2200 ns/op | 320 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 914352    | 2503 ns/op | 513 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 913231    | 2486 ns/op | 513 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 713877    | 2979 ns/op | 1282 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 755667    | 3003 ns/op | 1282 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 927482    | 2310 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1000000   | 2375 ns/op | 256 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 941809    | 2533 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 875191    | 2390 ns/op | 448 B/op  | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 672908    | 3015 ns/op | 1218 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 706266    | 2874 ns/op | 1217 B/op | 4 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 351468    | 6528 ns/op | 311 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 475314    | 5027 ns/op | 318 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 363056    | 6469 ns/op | 519 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 464019    | 5153 ns/op | 527 B/op  | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 326744    | 7086 ns/op | 1383 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 431764    | 5551 ns/op | 1392 B/op | 7 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 481654    | 4661 ns/op | 129 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 375290    | 5464 ns/op | 130 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 464568    | 4747 ns/op | 130 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 382231    | 5677 ns/op | 130 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 450594    | 4950 ns/op | 130 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 372163    | 5862 ns/op | 130 B/op  | 2 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 887011    | 2281 ns/op | 202 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1134375   | 2208 ns/op | 202 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 993236    | 2236 ns/op | 411 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1102100   | 2224 ns/op | 411 B/op  | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 828285    | 2649 ns/op | 1279 B/op | 6 allocs/op |
| BenchmarkClientHGet/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 918459    | 2625 ns/op | 1278 B/op | 6 allocs/op |
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+-----------+-------------+
```

```markdown
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
| Cluster Parallel HDel                                                                                              | iteration | ns/op      | B/op     | allocs/op   |
+====================================================================================================================+===========+============+==========+=============+
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 352831    | 6228 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 500095    | 4741 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 359395    | 6259 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 480048    | 4798 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 361002    | 6223 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP2:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 503013    | 4750 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4      | 1405093   | 1662 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4     | 1337671   | 1827 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4     | 1000000   | 2035 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4    | 1208085   | 1976 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4    | 1000000   | 2048 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/sandwich-go/redisson/RESP3:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4   | 1209093   | 1975 ns/op | 256 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4    | 1202887   | 1981 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4   | 1213291   | 2040 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4   | 1186718   | 1925 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4  | 1280984   | 1961 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4  | 1210052   | 2007 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/rueian/rueidis/rueidiscompat:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4 | 1192233   | 2016 ns/op | 200 B/op | 4 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4               | 364222    | 6288 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4              | 499416    | 4691 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4              | 357858    | 6211 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4             | 509692    | 4681 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4             | 351211    | 6187 ns/op | 231 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/go-redis/redis/v8:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4            | 507789    | 4724 ns/op | 238 B/op | 7 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4         | 469335    | 4646 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4        | 388164    | 5522 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4        | 474021    | 4652 ns/op | 129 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4       | 386308    | 5786 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4       | 481960    | 4648 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/mediocregopher/radix/v4:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4      | 391644    | 5565 ns/op | 130 B/op | 2 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(100):Parallel(128)-4              | 1201076   | 2015 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(64):Pool(1000):Parallel(128)-4             | 1000000   | 2030 ns/op | 98 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(100):Parallel(128)-4             | 1177996   | 2075 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(256):Pool(1000):Parallel(128)-4            | 1176146   | 2010 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(100):Parallel(128)-4            | 1000000   | 2132 ns/op | 97 B/op  | 4 allocs/op |
| BenchmarkClientHDel/github.com/joomcode/redispipe:Cluster:Key(16):Value(1024):Pool(1000):Parallel(128)-4           | 1201182   | 1997 ns/op | 97 B/op  | 4 allocs/op |
+--------------------------------------------------------------------------------------------------------------------+-----------+------------+----------+-------------+
```
