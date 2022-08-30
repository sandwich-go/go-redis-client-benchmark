module github.com/sandwich-go/go-redis-client-benchmark

go 1.16

require (
	github.com/chasex/redis-go-cluster v1.0.1-0.20161207023922-222d81891f1d
	github.com/garyburd/redigo v1.6.3 // indirect
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gomodule/redigo v1.8.9
	github.com/joomcode/errorx v1.1.0 // indirect
	github.com/joomcode/redispipe v0.9.4
	github.com/mediocregopher/radix.v2 v0.0.0-20181115013041-b67df6e626f9 // indirect
	github.com/mediocregopher/radix/v4 v4.1.1
	github.com/rueian/rueidis v0.0.74
	github.com/sandwich-go/redisson v1.1.13
	github.com/sandwich-go/rueidis v0.1.7
	github.com/sandwich-go/xconf v0.3.22
	github.com/valyala/fastrand v1.1.0
)

//replace github.com/sandwich-go/redisson => ../redisson
