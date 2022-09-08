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
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial Get                                 | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4953      | 487019   | 276   |       6   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4978      | 488346   | 276   |       6   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4966      | 485089   | 484   |       6   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4916      | 484548   | 484   |       6   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4929      | 487335   | 1348  |       6   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4933      | 488121   | 1348  |       6   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4958      | 485398   | 320   |       4   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4935      | 485190   | 320   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4966      | 484629   | 512   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4939      | 488952   | 512   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4954      | 489221   | 1280  |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4926      | 491789   | 1280  |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4987      | 483422   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4968      | 487068   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4986      | 486305   | 448   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4924      | 487030   | 448   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4922      | 487140   | 1216  |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4966      | 487747   | 1216  |       4   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4953      | 488424   | 276   |       6   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4971      | 495383   | 276   |       6   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4981      | 486368   | 484   |       6   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4936      | 484809   | 484   |       6   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4887      | 487234   | 1348  |       6   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4928      | 496201   | 1348  |       6   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4537      | 506352   | 876   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4512      | 533221   | 8989  |       36  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4537      | 510415   | 869   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4526      | 532493   | 8787  |       35  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4369      | 516728   | 903   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4312      | 537812   | 9222  |       37  |
| joomcode/redispipe:Val(64):Pool(100)              | 1546      | 1557321  | 168   |       5   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1565308  | 168   |       5   |
| joomcode/redispipe:Val(256):Pool(100)             | 1540      | 1573693  | 376   |       5   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1526      | 1555364  | 376   |       5   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1534      | 1561348  | 1240  |       5   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1532      | 1561439  | 1240  |       5   |
+---------------------------------------------------+-----------+----------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial Set                                 | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4977      | 483504   | 264   |       7   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4923      | 485434   | 264   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4930      | 487251   | 264   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4968      | 487332   | 264   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4926      | 488972   | 264   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4897      | 486064   | 264   |       7   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4939      | 483098   | 276   |       5   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4999      | 481544   | 276   |       5   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4928      | 484242   | 280   |       6   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4860      | 482453   | 280   |       6   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4968      | 484771   | 280   |       6   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4927      | 496207   | 280   |       6   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4808      | 484878   | 212   |       5   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4986      | 486763   | 212   |       5   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4946      | 485818   | 216   |       6   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4922      | 484952   | 216   |       6   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4945      | 486969   | 216   |       6   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4947      | 488603   | 216   |       6   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4957      | 482288   | 264   |       7   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4899      | 487379   | 264   |       7   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4897      | 488155   | 264   |       7   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4876      | 485042   | 264   |       7   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4926      | 489551   | 264   |       7   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4962      | 489919   | 264   |       7   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4506      | 519044   | 903   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4232      | 547546   | 9624  |       39  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4315      | 522241   | 946   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4465      | 532153   | 8998  |       36  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4530      | 509190   | 915   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4512      | 534769   | 9090  |       36  |
| joomcode/redispipe:Val(64):Pool(100)              | 1548      | 1558676  | 114   |       6   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1564027  | 114   |       6   |
| joomcode/redispipe:Val(256):Pool(100)             | 1540      | 1564637  | 114   |       6   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1527      | 1559428  | 114   |       6   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1537      | 1556612  | 115   |       6   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1536      | 1567739  | 115   |       6   |
+---------------------------------------------------+-----------+----------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial Del                                 | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4960      | 491199   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4969      | 487374   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4960      | 484226   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4956      | 486280   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4965      | 488280   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4981      | 481978   | 228   |       7   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4953      | 482947   | 288   |       5   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4989      | 485699   | 288   |       5   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4952      | 482216   | 288   |       5   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4965      | 483253   | 288   |       5   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4914      | 493253   | 288   |       5   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4960      | 487147   | 288   |       5   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 5024      | 482986   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4995      | 487758   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4846      | 482795   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 5001      | 482372   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4974      | 482227   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 5006      | 483856   | 200   |       4   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4936      | 488249   | 196   |       6   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4957      | 485687   | 196   |       6   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4912      | 486787   | 196   |       6   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4861      | 485094   | 196   |       6   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 5000      | 485747   | 196   |       6   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4932      | 484418   | 196   |       6   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4606      | 507947   | 872   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4581      | 531004   | 8862  |       36  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4508      | 505590   | 877   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4557      | 531726   | 8738  |       35  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4518      | 505961   | 873   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4596      | 527602   | 8664  |       35  |
| joomcode/redispipe:Val(64):Pool(100)              | 1568      | 1558597  | 64    |       3   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1537      | 1562103  | 64    |       3   |
| joomcode/redispipe:Val(256):Pool(100)             | 1537      | 1555240  | 64    |       3   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1557134  | 64    |       3   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1539      | 1553183  | 64    |       3   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1532      | 1558317  | 64    |       3   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial Exists                              | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4958      | 484697   | 196   |       6   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4974      | 481462   | 196   |       6   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4950      | 483400   | 196   |       6   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4930      | 490863   | 196   |       6   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4958      | 487538   | 196   |       6   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4948      | 484440   | 196   |       6   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4966      | 484385   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4999      | 484252   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4946      | 485223   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4930      | 484162   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4978      | 481993   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4932      | 482680   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4954      | 491144   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4918      | 484869   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4968      | 486363   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4947      | 481133   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4986      | 482829   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4974      | 486667   | 200   |       4   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4920      | 484542   | 196   |       6   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4915      | 490139   | 196   |       6   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4956      | 486650   | 196   |       6   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4944      | 478302   | 196   |       6   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 5011      | 488393   | 196   |       6   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4963      | 479900   | 196   |       6   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4549      | 508087   | 878   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4411      | 534180   | 9196  |       37  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4520      | 512916   | 874   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4533      | 525759   | 8777  |       36  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4558      | 511653   | 864   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4494      | 527723   | 8855  |       36  |
| joomcode/redispipe:Val(64):Pool(100)              | 1530      | 1559326  | 64    |       3   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1542      | 1556280  | 64    |       3   |
| joomcode/redispipe:Val(256):Pool(100)             | 1536      | 1564361  | 64    |       3   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1555467  | 64    |       3   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1542      | 1559493  | 64    |       3   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1543      | 1559163  | 64    |       3   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial Expire                              | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4849      | 483997   | 212   |       5   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4945      | 484310   | 212   |       5   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 5010      | 486427   | 212   |       5   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4995      | 488633   | 212   |       5   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4952      | 493152   | 212   |       5   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4996      | 487953   | 212   |       5   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4923      | 484355   | 240   |       3   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4980      | 484325   | 240   |       3   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4998      | 487745   | 240   |       3   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4972      | 482925   | 240   |       3   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4975      | 490133   | 240   |       3   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4963      | 484281   | 240   |       3   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 5010      | 487674   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4954      | 487549   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4953      | 487863   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4956      | 486432   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4944      | 483093   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4971      | 484827   | 184   |       3   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4948      | 486768   | 212   |       5   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4911      | 484882   | 212   |       5   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4910      | 489564   | 212   |       5   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4963      | 485558   | 212   |       5   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4948      | 481531   | 212   |       5   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 5002      | 489293   | 212   |       5   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4508      | 502751   | 905   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4417      | 545005   | 9205  |       37  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4478      | 511238   | 900   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4557      | 546055   | 8749  |       35  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4486      | 512832   | 896   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4533      | 541219   | 8794  |       36  |
| joomcode/redispipe:Val(64):Pool(100)              | 1534      | 1564919  | 80    |       3   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1534      | 1559231  | 80    |       3   |
| joomcode/redispipe:Val(256):Pool(100)             | 1540      | 1558428  | 80    |       3   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1464      | 1553472  | 80    |       3   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1540      | 1556370  | 80    |       3   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1537      | 1569699  | 80    |       3   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial TTL                                 | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4976      | 484425   | 196   |       5   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4957      | 485723   | 196   |       5   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4929      | 487783   | 196   |       5   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4981      | 481543   | 196   |       5   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 5007      | 483076   | 196   |       5   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4998      | 482210   | 196   |       5   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4934      | 486218   | 256   |       3   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4952      | 481857   | 256   |       3   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 5004      | 485067   | 256   |       3   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4988      | 481826   | 256   |       3   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4984      | 482331   | 256   |       3   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4914      | 481700   | 256   |       3   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4958      | 484226   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4936      | 480418   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4988      | 490442   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4960      | 485770   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 5013      | 481530   | 184   |       3   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4986      | 485811   | 184   |       3   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4972      | 485517   | 196   |       5   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4956      | 487745   | 196   |       5   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4946      | 482132   | 196   |       5   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4958      | 487410   | 196   |       5   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4989      | 489336   | 196   |       5   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4934      | 480211   | 196   |       5   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4586      | 518961   | 863   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4474      | 538017   | 9070  |       36  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4412      | 508648   | 893   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4404      | 525297   | 9037  |       37  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4431      | 509667   | 889   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4510      | 538578   | 8825  |       36  |
| joomcode/redispipe:Val(64):Pool(100)              | 1532      | 1562151  | 72    |       4   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1543      | 1553243  | 72    |       4   |
| joomcode/redispipe:Val(256):Pool(100)             | 1543      | 1562111  | 72    |       4   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1558208  | 72    |       4   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1540      | 1560193  | 72    |       4   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1538      | 1557668  | 72    |       4   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial HSet                                | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4922      | 485689   | 276   |       8   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4970      | 487203   | 276   |       8   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4984      | 486250   | 276   |       8   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4920      | 484587   | 276   |       8   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4876      | 487357   | 276   |       8   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4922      | 489322   | 276   |       8   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4917      | 484549   | 336   |       7   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4952      | 488471   | 336   |       7   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4902      | 487053   | 339   |       8   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4906      | 494815   | 339   |       8   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4840      | 486261   | 340   |       8   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4946      | 489426   | 340   |       8   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4968      | 485816   | 280   |       7   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4994      | 486246   | 280   |       7   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4834      | 488234   | 283   |       8   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4992      | 485545   | 283   |       8   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4850      | 485980   | 284   |       8   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4910      | 489439   | 284   |       8   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4927      | 485680   | 276   |       8   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4957      | 491002   | 276   |       8   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4930      | 488967   | 276   |       8   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4918      | 486592   | 276   |       8   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4905      | 488907   | 276   |       8   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4941      | 486194   | 276   |       8   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4525      | 523867   | 919   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4422      | 530238   | 9223  |       37  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4452      | 517385   | 932   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4507      | 531133   | 8929  |       36  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4333      | 519020   | 971   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4188      | 537108   | 9800  |       39  |
| joomcode/redispipe:Val(64):Pool(100)              | 1533      | 1561166  | 128   |       5   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1554166  | 128   |       5   |
| joomcode/redispipe:Val(256):Pool(100)             | 1540      | 1556652  | 128   |       5   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1533      | 1558041  | 128   |       5   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1533      | 1560643  | 129   |       5   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1538      | 1561489  | 129   |       5   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial HGet                                | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4852      | 501210   | 308   |       7   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4935      | 484346   | 308   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4962      | 487274   | 516   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4989      | 489618   | 516   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4910      | 488948   | 1380  |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4941      | 488380   | 1380  |       7   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4905      | 492828   | 320   |       4   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4620      | 485737   | 320   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4894      | 486326   | 512   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4976      | 485972   | 512   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4916      | 487300   | 1280  |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4900      | 490358   | 1280  |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4921      | 483546   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4980      | 485752   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4962      | 487444   | 448   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4887      | 490200   | 448   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4870      | 487478   | 1216  |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4951      | 486974   | 1216  |       4   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4966      | 487846   | 308   |       7   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4981      | 489512   | 308   |       7   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4888      | 483954   | 516   |       7   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4893      | 490433   | 516   |       7   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4934      | 491366   | 1380  |       7   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4909      | 488366   | 1380  |       7   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4548      | 511410   | 891   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4538      | 540596   | 8954  |       36  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4513      | 512147   | 892   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4458      | 545670   | 8939  |       36  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4315      | 508803   | 928   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4521      | 533149   | 8815  |       35  |
| joomcode/redispipe:Val(64):Pool(100)              | 1558      | 1561483  | 200   |       6   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1536      | 1561215  | 200   |       6   |
| joomcode/redispipe:Val(256):Pool(100)             | 1536      | 1559134  | 408   |       6   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1531      | 1557950  | 408   |       6   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1531      | 1559150  | 1272  |       6   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1519      | 1563919  | 1272  |       6   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Single Serial HDel                                | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 5004      | 485622   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4939      | 479118   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4953      | 486586   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4951      | 482200   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4984      | 482487   | 228   |       7   |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4825      | 485219   | 228   |       7   |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4954      | 485462   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4982      | 486066   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 5006      | 481440   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4962      | 484292   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4963      | 484289   | 256   |       4   |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4932      | 481204   | 256   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4933      | 486686   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4972      | 480096   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4975      | 486593   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4970      | 489757   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4945      | 483823   | 200   |       4   |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4995      | 482580   | 200   |       4   |
| go-redis/redis/v8:Val(64):Pool(100)               | 4964      | 484106   | 228   |       7   |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4959      | 482416   | 228   |       7   |
| go-redis/redis/v8:Val(256):Pool(100)              | 4956      | 482754   | 228   |       7   |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4917      | 500200   | 228   |       7   |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4975      | 484676   | 228   |       7   |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4896      | 481217   | 228   |       7   |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4549      | 503538   | 896   |       4   |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4596      | 546049   | 8853  |       35  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4508      | 513116   | 897   |       4   |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4544      | 536907   | 8779  |       35  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4437      | 506036   | 908   |       4   |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4518      | 525352   | 8829  |       36  |
| joomcode/redispipe:Val(64):Pool(100)              | 1538      | 1558144  | 96    |       4   |
| joomcode/redispipe:Val(64):Pool(1000)             | 1538      | 1555319  | 96    |       4   |
| joomcode/redispipe:Val(256):Pool(100)             | 1538      | 1556155  | 96    |       4   |
| joomcode/redispipe:Val(256):Pool(1000)            | 1534      | 1567664  | 96    |       4   |
| joomcode/redispipe:Val(1024):Pool(100)            | 1537      | 1560043  | 96    |       4   |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1540      | 1567453  | 96    |       4   |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) Get                          | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 362365    | 6136  | 279   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 504202    | 4731  | 286   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 362181    | 6334  | 487   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 481341    | 4946  | 495   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 332634    | 6822  | 1351  |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 451609    | 5299  | 1360  |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1208716   | 1923  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1263682   | 1870  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2013  | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 986697    | 2180  | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 728786    | 2816  | 1281  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 815188    | 2772  | 1281  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1253146   | 1847  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1300909   | 1874  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1000000   | 2034  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1000000   | 2091  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 792254    | 2686  | 1217  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 775755    | 2779  | 1217  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 369186    | 6098  | 279   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 506796    | 4750  | 286   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 357454    | 6266  | 487   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 486217    | 4919  | 495   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 331382    | 6779  | 1351  |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 452067    | 5307  | 1360  |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 596540    | 4284  | 26    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 589083    | 4902  | 54    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 576108    | 4384  | 27    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 597157    | 4993  | 54    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 573411    | 4539  | 27    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 559611    | 5062  | 56    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1109589   | 2137  | 168   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1000000   | 2129  | 169   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 1000000   | 2170  | 377   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 999955    | 2188  | 377   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 958350    | 2442  | 1241  |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 900044    | 2481  | 1241  |        5  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) Set                          | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 363010    | 6189  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 462764    | 5175  | 275   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 355052    | 6228  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 442138    | 5378  | 276   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 353577    | 6293  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 428907    | 5608  | 276   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1000000   | 2057  | 276   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1000000   | 2015  | 276   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1205239   | 1978  | 279   |        6  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1191349   | 1981  | 279   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1128141   | 2095  | 280   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1140604   | 2090  | 280   |        6  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1235583   | 1916  | 212   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1245774   | 1871  | 212   |        5  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1254902   | 1901  | 215   |        6  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1000000   | 2013  | 215   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1159359   | 2095  | 216   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1160755   | 2048  | 216   |        6  |
| go-redis/redis/v8:Val(64):Pool(100)               | 364204    | 6139  | 267   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 465283    | 5168  | 275   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 356728    | 6214  | 267   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 446660    | 5381  | 276   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 352990    | 6262  | 267   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 430051    | 5570  | 276   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 575492    | 4293  | 43    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 576790    | 5200  | 73    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 565288    | 4328  | 43    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 563811    | 5171  | 72    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 585178    | 4411  | 43    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 541306    | 5242  | 75    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2072  | 115   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1235098   | 1940  | 115   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 1355856   | 1978  | 116   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1429597   | 1983  | 116   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1208988   | 2020  | 134   |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1246899   | 1923  | 132   |        6  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) Del                          | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 356623    | 6106  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 527402    | 4529  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 366289    | 6140  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 529042    | 4551  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 362114    | 6083  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 529304    | 4518  | 238   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1342072   | 1757  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1225780   | 2011  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2058  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1000000   | 2072  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 984410    | 2055  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1218895   | 2027  | 288   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1228944   | 1898  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1215169   | 1898  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1238926   | 1915  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1270002   | 1924  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1226511   | 1925  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1211758   | 1856  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 359362    | 6063  | 199   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 532927    | 4550  | 206   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 369366    | 6031  | 199   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 527113    | 4504  | 206   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 359216    | 6040  | 199   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 530523    | 4497  | 206   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 596355    | 4252  | 26    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 645172    | 4854  | 53    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 612879    | 4248  | 26    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 628474    | 4831  | 53    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 508616    | 4265  | 28    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 624930    | 4816  | 53    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2165  | 65    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 982747    | 2095  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1000000   | 2153  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1000000   | 2122  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1000000   | 2110  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1000000   | 2151  | 64    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) Exists                       | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 355356    | 6091  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 521980    | 4566  | 206   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 365894    | 6054  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 524054    | 4552  | 206   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 374433    | 6045  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 527689    | 4526  | 206   |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1184643   | 1981  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1209036   | 1998  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1206621   | 2020  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1219750   | 2028  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1234136   | 1972  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1236686   | 1958  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1212618   | 1974  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1237340   | 1962  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1233102   | 1970  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1256658   | 2007  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1235353   | 2025  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1245318   | 1955  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 362736    | 6061  | 199   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 525018    | 4546  | 206   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 370290    | 6048  | 199   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 531529    | 4532  | 206   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 360061    | 6053  | 199   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 525180    | 4567  | 206   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 590653    | 4283  | 26    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 672496    | 4972  | 54    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 609844    | 4275  | 26    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 605505    | 4803  | 53    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 601099    | 4266  | 26    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 632251    | 4842  | 52    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2185  | 64    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1128975   | 2085  | 64    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1000000   | 2146  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1133168   | 2213  | 64    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1000000   | 2187  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1113512   | 2188  | 64    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) Expire                       | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 332355    | 6199  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 439668    | 5446  | 224   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 353678    | 6150  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 440203    | 5493  | 224   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 357690    | 6174  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 440194    | 5441  | 224   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 928232    | 2208  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1060755   | 2162  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2210  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 967376    | 2238  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 946177    | 2289  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 984056    | 2365  | 240   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1000000   | 2167  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2354  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1000000   | 2405  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1125740   | 2221  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1000000   | 2432  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 991659    | 2414  | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 364699    | 6139  | 215   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 441435    | 5436  | 224   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 357466    | 6231  | 215   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 415408    | 5548  | 225   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 363788    | 6138  | 215   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 439591    | 5446  | 224   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 566857    | 4225  | 43    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 599179    | 5271  | 75    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 630436    | 4215  | 41    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 545533    | 5173  | 74    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 575806    | 4233  | 42    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 574726    | 5222  | 74    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 834244    | 2574  | 81    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 767469    | 2607  | 81    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 863950    | 2601  | 81    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 898284    | 2591  | 81    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 848350    | 2592  | 81    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 865720    | 2549  | 81    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) TTL                          | iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 354982    | 6075  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 503134    | 4750  | 206   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 365344    | 6065  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 503127    | 4740  | 206   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 366339    | 6101  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 502210    | 4760  | 206   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1341606   | 1876  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1258068   | 1915  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1200594   | 1875  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1247716   | 1878  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1283361   | 1883  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1204140   | 1947  | 256   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1305181   | 1891  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1310246   | 1892  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1259728   | 1913  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1290460   | 1851  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1283109   | 1792  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1298436   | 1889  | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 366732    | 6067  | 199   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 505086    | 4748  | 206   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 367401    | 6028  | 199   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 494394    | 4890  | 207   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 361278    | 6096  | 199   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 506158    | 4755  | 206   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 585817    | 4270  | 27    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 545426    | 4779  | 55    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 610243    | 4251  | 26    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 561709    | 4851  | 54    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 632330    | 4226  | 26    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 584240    | 4805  | 55    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2192  | 72    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 954081    | 2201  | 73    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 1000000   | 2170  | 73    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1119945   | 2214  | 72    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1051593   | 2263  | 73    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1000000   | 2181  | 73    |        4  |
+---------------------------------------------------+-----------+-------+-------+-----------+   
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) HSet                       	| iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 353491    | 6200  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 463742    | 5195  | 287   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 351422    | 6251  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 446103    | 5350  | 288   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 350900    | 6418  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 399879    | 5643  | 289   |        8  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 910920    | 2214  | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 971934    | 2147  | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1246284   | 2083  | 339   |        8  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1239913   | 2046  | 339   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1116276   | 2146  | 340   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 958636    | 2166  | 340   |        8  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1218153   | 1953  | 280   |        7  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1214400   | 2089  | 280   |        7  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1230264   | 1962  | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1247301   | 1909  | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1139823   | 2136  | 284   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1122312   | 2141  | 284   |        8  |
| go-redis/redis/v8:Val(64):Pool(100)               | 351799    | 6188  | 279   |        8  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 462050    | 5170  | 287   |        8  |
| go-redis/redis/v8:Val(256):Pool(100)              | 356505    | 6564  | 279   |        8  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 433468    | 5468  | 288   |        8  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 349581    | 6268  | 279   |        8  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 427269    | 5660  | 288   |        8  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 570483    | 4336  | 59    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 613366    | 5156  | 88    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 564765    | 4397  | 59    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 578872    | 5339  | 90    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 531266    | 4459  | 60    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 555976    | 5264  | 90    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2131  | 129   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1111388   | 2208  | 129   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 1380260   | 1688  | 131   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1210473   | 1790  | 130   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1215700   | 1963  | 149   |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1000000   | 2037  | 151   |        5  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) HGet                       	| iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 354025    | 6300  | 311   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 463798    | 5210  | 320   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 351859    | 6380  | 519   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 472824    | 5059  | 527   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 326384    | 6987  | 1383  |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 434271    | 5425  | 1392  |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1110710   | 2235  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 928776    | 2240  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 940675    | 2402  | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1000000   | 2304  | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 785270    | 2829  | 1281  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 720984    | 2796  | 1281  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 965226    | 2091  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2143  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 909498    | 2309  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1081938   | 2244  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 861998    | 2790  | 1217  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 853593    | 2790  | 1217  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 365218    | 6171  | 311   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 486632    | 5064  | 319   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 349214    | 6398  | 519   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 469039    | 5075  | 527   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 331870    | 6819  | 1383  |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 440664    | 5458  | 1392  |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 583803    | 4315  | 43    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 632470    | 4945  | 71    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 661624    | 4388  | 41    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 641383    | 5014  | 70    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 597493    | 4558  | 42    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 582472    | 5167  | 72    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 924786    | 2202  | 201   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1000000   | 2181  | 201   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 1000000   | 2149  | 409   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1000000   | 2172  | 409   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 950053    | 2523  | 1273  |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 828849    | 2565  | 1273  |        6  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+
| Single Parallel(128) HDel                       	| iteration | ns/op | B/op  | allocs/op |
+===================================================+===========+=======+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 358890    | 6126  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 520810    | 4607  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 357205    | 6100  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 507561    | 4626  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 362708    | 6104  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 519915    | 4616  | 238   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1000000   | 2063  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1211716   | 2042  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2081  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1200013   | 2010  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1000000   | 2085  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1000000   | 2073  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1000000   | 2009  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2081  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1000000   | 2007  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1238990   | 2002  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1000000   | 2118  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1000000   | 2046  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 361191    | 6105  | 231   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 507292    | 4753  | 238   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 361250    | 6112  | 231   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 520888    | 4604  | 238   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 364586    | 6070  | 231   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 513022    | 4623  | 238   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 605346    | 4328  | 42    |        1  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 610455    | 4834  | 69    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 575258    | 4321  | 43    |        1  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 638984    | 4868  | 69    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 620847    | 4379  | 42    |        1  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 594288    | 4893  | 70    |        1  |
| joomcode/redispipe:Val(64):Pool(100)              | 1182980   | 2212  | 96    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 994881    | 2047  | 97    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 1210677   | 2046  | 97    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1153348   | 1976  | 97    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1000000   | 2075  | 97    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1199133   | 1984  | 96    |        4  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

#### Cluster Redis
```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial Get                       		    | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4890      | 487417   | 276   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4915      | 488523   | 276   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4964      | 491417   | 484   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4939      | 484393   | 484   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4914      | 488911   | 1348  |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4900      | 488106   | 1348  |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4938      | 488312   | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4958      | 487591   | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4960      | 485184   | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4983      | 487830   | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4928      | 488460   | 1280  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4878      | 489054   | 1280  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 5013      | 486590   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4929      | 489007   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4918      | 485077   | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4926      | 485526   | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4899      | 490463   | 1216  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4939      | 491525   | 1216  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4861      | 489484   | 276   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4958      | 484996   | 276   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4929      | 488630   | 484   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4863      | 486737   | 484   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4876      | 497180   | 1348  |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4872      | 494822   | 1348  |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4729      | 517134   | 118   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4700      | 521479   | 175   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4725      | 506786   | 118   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4663      | 529271   | 176   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4563      | 512028   | 118   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4707      | 523469   | 175   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1538      | 1560667  | 169   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1536      | 1560187  | 169   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 1533      | 1563273  | 377   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1559477  | 377   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1533      | 1576437  | 1241  |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1534      | 1562195  | 1242  |        5  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial Set                       		    | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4906      | 493421   | 264   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4958      | 489022   | 264   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4881      | 487076   | 264   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4879      | 494568   | 264   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4892      | 493964   | 264   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4861      | 493905   | 264   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4915      | 489125   | 342   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4898      | 489537   | 276   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4933      | 486003   | 345   |        6  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4933      | 491336   | 280   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4860      | 494323   | 280   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4892      | 493207   | 280   |        6  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4927      | 486205   | 212   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4935      | 490395   | 277   |        5  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4960      | 491011   | 216   |        6  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4945      | 486436   | 216   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4914      | 488440   | 216   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4784      | 489360   | 216   |        6  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4870      | 496824   | 264   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4898      | 493131   | 264   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4876      | 500501   | 264   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4875      | 487160   | 264   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4855      | 497155   | 264   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4845      | 496502   | 264   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4707      | 520846   | 136   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4606      | 545805   | 212   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4518      | 512600   | 140   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 3937      | 533491   | 264   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4635      | 509365   | 159   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4528      | 548624   | 448   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1564      | 1559278  | 116   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1532      | 1564312  | 116   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 1527      | 1564032  | 119   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1536      | 1562740  | 116   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1530      | 1567822  | 117   |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1530      | 1568534  | 117   |        6  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial Del                       		    | iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4954      | 482505   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4909      | 490404   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4972      | 489540   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4965      | 484667   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4951      | 483778   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4947      | 486319   | 228   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4915      | 485578   | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4992      | 483628   | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4927      | 486778   | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4989      | 483615   | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4966      | 486536   | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4923      | 481059   | 288   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4885      | 486639   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4980      | 483367   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 5012      | 488939   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4950      | 482863   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4968      | 482569   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4848      | 484740   | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4888      | 494681   | 196   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4899      | 483609   | 196   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4956      | 488602   | 196   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4934      | 487820   | 196   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4910      | 490078   | 196   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4882      | 486974   | 196   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4572      | 516742   | 118   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4581      | 533192   | 176   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4382      | 524153   | 118   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4647      | 515357   | 174   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4366      | 525362   | 118   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4699      | 517601   | 176   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1549      | 1557796  | 66    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1557442  | 66    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1537      | 1559602  | 66    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1554963  | 66    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1539      | 1571336  | 66    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1542      | 1552898  | 66    |        3  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial Exists                       		| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4948      | 486207   | 196   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4945      | 483354   | 196   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4980      | 484337   | 196   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4996      | 492530   | 196   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4989      | 484051   | 196   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4928      | 489495   | 196   |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4976      | 481683   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4958      | 483620   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4990      | 484575   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4978      | 484997   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4920      | 486551   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4952      | 487337   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4978      | 481569   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4996      | 486897   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4992      | 483937   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4963      | 488358   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4953      | 483852   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4945      | 484625   | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4926      | 491615   | 196   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4908      | 486800   | 196   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4912      | 488241   | 196   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4918      | 509994   | 196   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4884      | 485003   | 196   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4978      | 486328   | 196   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4714      | 514884   | 118   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4471      | 528009   | 176   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4380      | 507985   | 118   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4460      | 542699   | 178   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4792      | 508285   | 118   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4303      | 517149   | 174   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1545      | 1559345  | 65    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1537      | 1556615  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1544      | 1583717  | 66    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1536      | 1557796  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1543      | 1559059  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1533      | 1559116  | 65    |        3  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial Expire                       		| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4939      | 490287   | 212   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4898      | 490532   | 212   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4868      | 488661   | 212   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4870      | 487149   | 212   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4879      | 492483   | 212   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4934      | 489043   | 212   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4960      | 487726   | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4890      | 488136   | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4909      | 498026   | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4893      | 512902   | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4914      | 492731   | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4892      | 491587   | 240   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4940      | 489265   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4921      | 484519   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4894      | 493868   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4920      | 489021   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4936      | 488031   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4934      | 484993   | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4863      | 494336   | 212   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4887      | 490703   | 212   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4845      | 501066   | 212   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4850      | 492617   | 212   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4904      | 493192   | 212   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4854      | 487685   | 212   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4663      | 529580   | 1296  |        4  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4604      | 542000   | 194   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4531      | 515191   | 134   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4590      | 536272   | 193   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4308      | 519950   | 134   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4333      | 531320   | 193   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1533      | 1564005  | 81    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1536      | 1572816  | 81    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1531      | 1565808  | 81    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1528      | 1567128  | 81    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1530      | 1569485  | 82    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1532      | 1563870  | 82    |        3  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial TTL                       			| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4986      | 485173   | 196   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4983      | 487435   | 196   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4971      | 486494   | 196   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4951      | 485958   | 196   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4977      | 486930   | 196   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4965      | 485725   | 196   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4988      | 482261   | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4963      | 482398   | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4976      | 484914   | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4550      | 481445   | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4994      | 485167   | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4988      | 480536   | 256   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4981      | 480642   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4958      | 487868   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4965      | 480767   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 5011      | 484296   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4977      | 483220   | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4971      | 489272   | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4956      | 486490   | 196   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4953      | 489336   | 196   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4921      | 487203   | 196   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4896      | 490285   | 196   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4909      | 489089   | 196   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4899      | 484813   | 196   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4392      | 516994   | 118   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4162      | 556218   | 180   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4772      | 503889   | 118   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4322      | 525965   | 176   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4254      | 506998   | 118   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4707      | 531403   | 176   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1569      | 1550010  | 74    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1497      | 1574281  | 74    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 1533      | 1562076  | 73    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1540      | 1555787  | 73    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1537      | 1556139  | 73    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1538      | 1561719  | 73    |        4  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial HSet                       		| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4882      | 490112   | 276   |        8  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4894      | 492661   | 276   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4893      | 489208   | 276   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4897      | 492833   | 276   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4894      | 501095   | 276   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4898      | 493391   | 276   |        8  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4939      | 492085   | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4915      | 494991   | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4918      | 486724   | 339   |        8  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4898      | 489289   | 339   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4868      | 492934   | 340   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4898      | 489974   | 340   |        8  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4922      | 495294   | 345   |        7  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4759      | 490248   | 348   |        7  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4770      | 487892   | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4927      | 489624   | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4884      | 494230   | 284   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4854      | 494940   | 350   |        8  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4828      | 492591   | 276   |        8  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4860      | 492739   | 276   |        8  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4928      | 497156   | 276   |        8  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4886      | 490350   | 276   |        8  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4856      | 506016   | 276   |        8  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4818      | 497876   | 276   |        8  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4711      | 516181   | 152   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4293      | 525695   | 227   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4612      | 521138   | 156   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4635      | 525211   | 270   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4178      | 531872   | 177   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4273      | 543087   | 475   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1528      | 1562925  | 130   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1525      | 1563617  | 130   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 1536      | 1558767  | 133   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1534      | 1570781  | 130   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1533      | 1578202  | 131   |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1533      | 1570661  | 131   |        5  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial HGet                       		| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4922      | 486146   | 308   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4922      | 488823   | 308   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4920      | 487590   | 516   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4942      | 490708   | 516   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4910      | 490163   | 1380  |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4906      | 488618   | 1380  |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4980      | 487716   | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4942      | 491867   | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4894      | 489296   | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4902      | 490110   | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4876      | 484093   | 1280  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4924      | 488573   | 1280  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4981      | 483550   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4946      | 489059   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4957      | 483505   | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 4920      | 490375   | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 4945      | 486025   | 1216  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 4915      | 491819   | 1216  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4938      | 486744   | 308   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4887      | 488740   | 308   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4840      | 491426   | 516   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4894      | 498615   | 516   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4921      | 493442   | 1380  |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4929      | 489069   | 1380  |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4664      | 527804   | 134   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4432      | 541879   | 193   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4320      | 512200   | 134   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4291      | 520926   | 192   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4744      | 507930   | 1264  |        4  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4654      | 552060   | 196   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1555      | 1558220  | 201   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1540      | 1563677  | 201   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 1533      | 1563890  | 409   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1532      | 1556235  | 409   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1534      | 1567017  | 1274  |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1531      | 1561188  | 1273  |        6  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+----------+-------+-----------+
| Cluster Serial HDel                       		| iteration | ns/op    | B/op  | allocs/op |
+===================================================+===========+==========+=======+===========+
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 4929      | 487506   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 4665      | 486410   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 4964      | 486650   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 4972      | 483611   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 4966      | 489641   | 228   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 4958      | 485558   | 228   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 4800      | 491455   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 4951      | 479973   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 4956      | 483460   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 4996      | 485347   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 4981      | 485060   | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 4986      | 487348   | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 4888      | 483396   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 4968      | 483600   | 265   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 4976      | 488035   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 5008      | 481217   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 5020      | 485092   | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 5013      | 483332   | 264   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 4900      | 488946   | 228   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 4950      | 513867   | 228   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 4910      | 487995   | 228   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 4875      | 484749   | 228   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 4922      | 489232   | 228   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 4942      | 489995   | 228   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 4756      | 519971   | 134   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 4586      | 513765   | 190   |        3  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 4449      | 528388   | 134   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 4317      | 515906   | 191   |        3  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 4656      | 522122   | 134   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 4429      | 514812   | 191   |        3  |
| joomcode/redispipe:Val(64):Pool(100)              | 1539      | 1555304  | 98    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1539      | 1557930  | 98    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 1539      | 1572486  | 98    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1405      | 1554021  | 98    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1538      | 1555258  | 98    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1536      | 1557521  | 98    |        4  |
+---------------------------------------------------+-----------+----------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) Get                         | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 361689    | 6246  | 279   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 494625    | 4819  | 286   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 353413    | 6439  | 487   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 478305    | 5035  | 494   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 324940    | 6992  | 1351  |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 441291    | 5472  | 1360  |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1036126   | 2275  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1067506   | 2250  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1008175   | 2420  | 513   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 883516    | 2394  | 512   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 766168    | 2906  | 1282  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 729106    | 2878  | 1282  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 946216    | 2266  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2216  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 924811    | 2292  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 969163    | 2306  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 856582    | 2802  | 1218  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 814442    | 2819  | 1217  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 351850    | 6251  | 279   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 489259    | 4821  | 286   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 356703    | 6385  | 487   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 478236    | 5012  | 494   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 333362    | 6972  | 1351  |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 443264    | 5386  | 1360  |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 477573    | 4598  | 113   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 386779    | 5431  | 114   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 459818    | 4737  | 113   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 383200    | 5656  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 451070    | 4911  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 356745    | 5745  | 114   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1091751   | 2147  | 170   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1116862   | 2133  | 170   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 1088572   | 2298  | 379   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1073509   | 2189  | 379   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 800530    | 2548  | 1246  |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 934084    | 2595  | 1246  |        5  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) Set                         | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 349974    | 6494  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 404748    | 5936  | 277   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 337305    | 6662  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 324763    | 6220  | 280   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 323359    | 6979  | 267   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 299281    | 6823  | 281   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 946690    | 2327  | 276   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 927264    | 2324  | 276   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 814446    | 2518  | 279   |        6  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 820002    | 2476  | 279   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 621984    | 3377  | 280   |        6  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 662259    | 3376  | 280   |        6  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 939330    | 2276  | 212   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 888968    | 2264  | 212   |        5  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 874808    | 2477  | 215   |        6  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 837128    | 2462  | 215   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 656972    | 3348  | 216   |        6  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 639235    | 3363  | 216   |        6  |
| go-redis/redis/v8:Val(64):Pool(100)               | 336496    | 6471  | 267   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 339688    | 5931  | 279   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 342784    | 6586  | 267   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 384579    | 6300  | 277   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 318220    | 7096  | 267   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 300900    | 6820  | 281   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 487761    | 4476  | 129   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 387138    | 5689  | 130   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 547970    | 4467  | 129   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 370246    | 5938  | 131   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 495886    | 4527  | 129   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 345834    | 6434  | 133   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 981516    | 2269  | 115   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 933568    | 2328  | 115   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 1018083   | 2366  | 117   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 928414    | 2350  | 116   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 654692    | 3278  | 139   |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 678476    | 3260  | 134   |        6  |
+---------------------------------------------------+-----------+-------+-------+-----------+  
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) Del                         | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 358837    | 6260  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 497517    | 4832  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 350084    | 6197  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 493366    | 4828  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 363262    | 6290  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 501684    | 4814  | 238   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 993585    | 2125  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1137520   | 2114  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 985495    | 2119  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1000000   | 2146  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 987890    | 2216  | 288   |        5  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1000000   | 2274  | 288   |        5  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 945025    | 2213  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2179  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1123978   | 2176  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1111022   | 2169  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1104888   | 2157  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1000000   | 2150  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 364538    | 6146  | 199   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 511624    | 4701  | 206   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 370584    | 6229  | 199   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 514002    | 4622  | 206   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 368612    | 6152  | 199   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 514106    | 4768  | 206   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 471505    | 4619  | 114   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 394359    | 5503  | 114   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 467673    | 4604  | 113   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 389193    | 5547  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 484116    | 4625  | 113   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 399212    | 5503  | 114   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2101  | 65    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 941641    | 2151  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1140456   | 2041  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1181277   | 2073  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1180410   | 2071  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1000000   | 2036  | 65    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) Exists                      | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 356186    | 6209  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 501278    | 4735  | 206   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 360022    | 6176  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 508154    | 4783  | 206   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 359449    | 6184  | 199   |        6  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 507560    | 4717  | 206   |        6  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 931364    | 2190  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 939031    | 2162  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 972061    | 2169  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 919816    | 2222  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1000000   | 2201  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 994989    | 2221  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1106865   | 2175  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2169  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 935419    | 2221  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1069153   | 2119  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1000000   | 2153  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1101631   | 2160  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 360195    | 6205  | 199   |        6  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 511290    | 4669  | 206   |        6  |
| go-redis/redis/v8:Val(256):Pool(100)              | 360906    | 6171  | 199   |        6  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 516358    | 4667  | 206   |        6  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 359160    | 6164  | 199   |        6  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 516640    | 4687  | 206   |        6  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 477600    | 4610  | 114   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 465466    | 5489  | 113   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 478651    | 4640  | 113   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 385738    | 5520  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 474693    | 4699  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 381096    | 5538  | 114   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1180543   | 2046  | 65    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 986827    | 2100  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 1204221   | 2064  | 65    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1000000   | 2058  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1177262   | 1998  | 65    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1210532   | 2071  | 65    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) Expire                      | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 336084    | 6540  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 391800    | 6132  | 225   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 334539    | 6553  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 391501    | 6132  | 225   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 333081    | 6559  | 215   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 391326    | 6168  | 225   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 792888    | 2812  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 668226    | 2996  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 714886    | 2977  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 691064    | 2906  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 792148    | 3052  | 240   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 732633    | 2942  | 240   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 752437    | 3002  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 719923    | 2945  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 739682    | 2749  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 809517    | 2800  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 689179    | 3083  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 858776    | 2785  | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 342097    | 6526  | 215   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 329268    | 6150  | 227   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 345883    | 6516  | 215   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 331440    | 6168  | 227   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 335524    | 6554  | 215   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 328159    | 6153  | 227   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 495571    | 4423  | 129   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 365419    | 5965  | 130   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 503380    | 4353  | 129   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 378637    | 5840  | 130   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 502240    | 4421  | 129   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 366000    | 5830  | 130   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 836614    | 2915  | 81    |        3  |
| joomcode/redispipe:Val(64):Pool(1000)             | 821937    | 2908  | 82    |        3  |
| joomcode/redispipe:Val(256):Pool(100)             | 718962    | 2864  | 82    |        3  |
| joomcode/redispipe:Val(256):Pool(1000)            | 818583    | 2888  | 82    |        3  |
| joomcode/redispipe:Val(1024):Pool(100)            | 737113    | 2878  | 82    |        3  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 754645    | 2848  | 82    |        3  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) TTL                         | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 359133    | 6306  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 498634    | 4800  | 206   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 358323    | 6203  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 492385    | 5047  | 206   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 356343    | 6209  | 199   |        5  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 501309    | 4797  | 206   |        5  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1279167   | 1906  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1265374   | 1836  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2022  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1200952   | 1857  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1259340   | 1892  | 256   |        3  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1246321   | 1950  | 256   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1335216   | 1821  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1281430   | 1934  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1308994   | 1771  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1357956   | 1896  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1292203   | 1777  | 184   |        3  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1306335   | 1805  | 184   |        3  |
| go-redis/redis/v8:Val(64):Pool(100)               | 360853    | 6166  | 199   |        5  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 499804    | 4811  | 206   |        5  |
| go-redis/redis/v8:Val(256):Pool(100)              | 360001    | 6177  | 199   |        5  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 497072    | 4797  | 206   |        5  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 366268    | 6249  | 199   |        5  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 500016    | 4806  | 206   |        5  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 488238    | 4596  | 114   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 388599    | 5569  | 114   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 459261    | 4610  | 113   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 466982    | 5558  | 114   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 474943    | 4618  | 113   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 387424    | 5465  | 114   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1000000   | 2138  | 73    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1121539   | 2161  | 73    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 999075    | 2116  | 73    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 987087    | 2186  | 74    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 997465    | 2101  | 73    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1000000   | 2104  | 73    |        4  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) HSet                        | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 344802    | 6604  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 392251    | 6089  | 289   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 335762    | 6712  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 313497    | 6391  | 292   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 312363    | 7201  | 279   |        8  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 289394    | 7092  | 293   |        8  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 953574    | 2484  | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 837054    | 2635  | 336   |        7  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 761954    | 2690  | 340   |        8  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 813997    | 2686  | 340   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 571730    | 3632  | 340   |        8  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 619077    | 3599  | 340   |        8  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 867099    | 2418  | 280   |        7  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 981193    | 2418  | 280   |        7  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 792901    | 2668  | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 834783    | 2667  | 283   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 596102    | 3606  | 284   |        8  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 602301    | 3586  | 284   |        8  |
| go-redis/redis/v8:Val(64):Pool(100)               | 344406    | 6586  | 279   |        8  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 393264    | 6158  | 289   |        8  |
| go-redis/redis/v8:Val(256):Pool(100)              | 359312    | 6700  | 279   |        8  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 315304    | 6400  | 292   |        8  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 305941    | 7163  | 279   |        8  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 290716    | 7080  | 293   |        8  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 495270    | 4440  | 145   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 364306    | 5865  | 146   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 488095    | 4542  | 145   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 550398    | 6096  | 146   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 467059    | 4760  | 146   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 327129    | 6858  | 150   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1019394   | 2344  | 129   |        5  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1004547   | 2343  | 129   |        5  |
| joomcode/redispipe:Val(256):Pool(100)             | 857570    | 2537  | 132   |        5  |
| joomcode/redispipe:Val(256):Pool(1000)            | 864350    | 2527  | 129   |        5  |
| joomcode/redispipe:Val(1024):Pool(100)            | 637536    | 3487  | 152   |        5  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 628860    | 3477  | 153   |        5  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) HGet                        | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 369856    | 6348  | 311   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 476847    | 5070  | 318   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 366747    | 6478  | 519   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 463682    | 5152  | 527   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 331086    | 7036  | 1383  |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 429591    | 5580  | 1392  |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 959172    | 2170  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1124628   | 2200  | 320   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 914352    | 2503  | 513   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 913231    | 2486  | 513   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 713877    | 2979  | 1282  |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 755667    | 3003  | 1282  |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 927482    | 2310  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1000000   | 2375  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 941809    | 2533  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 875191    | 2390  | 448   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 672908    | 3015  | 1218  |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 706266    | 2874  | 1217  |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 351468    | 6528  | 311   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 475314    | 5027  | 318   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 363056    | 6469  | 519   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 464019    | 5153  | 527   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 326744    | 7086  | 1383  |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 431764    | 5551  | 1392  |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 481654    | 4661  | 129   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 375290    | 5464  | 130   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 464568    | 4747  | 130   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 382231    | 5677  | 130   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 450594    | 4950  | 130   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 372163    | 5862  | 130   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 887011    | 2281  | 202   |        6  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1134375   | 2208  | 202   |        6  |
| joomcode/redispipe:Val(256):Pool(100)             | 993236    | 2236  | 411   |        6  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1102100   | 2224  | 411   |        6  |
| joomcode/redispipe:Val(1024):Pool(100)            | 828285    | 2649  | 1279  |        6  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 918459    | 2625  | 1278  |        6  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```

```markdown
+---------------------------------------------------+-----------+-------+-------+-----------+ 
| Cluster Parallel(128) HDel                        | iteration | ns/op | B/op  | allocs/op | 
+===================================================+===========+=======+=======+===========+ 
| sandwich-go/redisson/RESP2:Val(64):Pool(100)      | 352831    | 6228  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(64):Pool(1000)     | 500095    | 4741  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(100)     | 359395    | 6259  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(256):Pool(1000)    | 480048    | 4798  | 238   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(100)    | 361002    | 6223  | 231   |        7  |
| sandwich-go/redisson/RESP2:Val(1024):Pool(1000)   | 503013    | 4750  | 238   |        7  |
| sandwich-go/redisson/RESP3:Val(64):Pool(100)      | 1405093   | 1662  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(64):Pool(1000)     | 1337671   | 1827  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(100)     | 1000000   | 2035  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(256):Pool(1000)    | 1208085   | 1976  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(100)    | 1000000   | 2048  | 256   |        4  |
| sandwich-go/redisson/RESP3:Val(1024):Pool(1000)   | 1209093   | 1975  | 256   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(100)    | 1202887   | 1981  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(64):Pool(1000)   | 1213291   | 2040  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(100)   | 1186718   | 1925  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(256):Pool(1000)  | 1280984   | 1961  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(100)  | 1210052   | 2007  | 200   |        4  |
| rueian/rueidis/rueidiscompat:Val(1024):Pool(1000) | 1192233   | 2016  | 200   |        4  |
| go-redis/redis/v8:Val(64):Pool(100)               | 364222    | 6288  | 231   |        7  |
| go-redis/redis/v8:Val(64):Pool(1000)              | 499416    | 4691  | 238   |        7  |
| go-redis/redis/v8:Val(256):Pool(100)              | 357858    | 6211  | 231   |        7  |
| go-redis/redis/v8:Val(256):Pool(1000)             | 509692    | 4681  | 238   |        7  |
| go-redis/redis/v8:Val(1024):Pool(100)             | 351211    | 6187  | 231   |        7  |
| go-redis/redis/v8:Val(1024):Pool(1000)            | 507789    | 4724  | 238   |        7  |
| mediocregopher/radix/v4:Val(64):Pool(100)         | 469335    | 4646  | 130   |        2  |
| mediocregopher/radix/v4:Val(64):Pool(1000)        | 388164    | 5522  | 130   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(100)        | 474021    | 4652  | 129   |        2  |
| mediocregopher/radix/v4:Val(256):Pool(1000)       | 386308    | 5786  | 130   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(100)       | 481960    | 4648  | 130   |        2  |
| mediocregopher/radix/v4:Val(1024):Pool(1000)      | 391644    | 5565  | 130   |        2  |
| joomcode/redispipe:Val(64):Pool(100)              | 1201076   | 2015  | 97    |        4  |
| joomcode/redispipe:Val(64):Pool(1000)             | 1000000   | 2030  | 98    |        4  |
| joomcode/redispipe:Val(256):Pool(100)             | 1177996   | 2075  | 97    |        4  |
| joomcode/redispipe:Val(256):Pool(1000)            | 1176146   | 2010  | 97    |        4  |
| joomcode/redispipe:Val(1024):Pool(100)            | 1000000   | 2132  | 97    |        4  |
| joomcode/redispipe:Val(1024):Pool(1000)           | 1201182   | 1997  | 97    |        4  |
+---------------------------------------------------+-----------+-------+-------+-----------+ 
```
