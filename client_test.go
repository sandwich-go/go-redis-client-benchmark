package go_redis_client_benchmark

import (
	"context"
	"fmt"
	redisgocluster "github.com/chasex/redis-go-cluster"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	redispipe "github.com/joomcode/redispipe/redis"
	"github.com/joomcode/redispipe/rediscluster"
	"github.com/joomcode/redispipe/redisconn"
	"github.com/mediocregopher/radix/v4"
	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/rueidiscompat"
	"github.com/sandwich-go/go-redis-client-benchmark/config"
	redisConfig "github.com/sandwich-go/go-redis-client-benchmark/config/redis"
	"github.com/sandwich-go/redisson"
	"github.com/valyala/fastrand"
	"runtime"
	"strings"
	"testing"
	"time"
)

var emptyTarget = Target{}

type RedisMode string

const (
	RedisModeSingle   = "Single"
	RedisModeCluster  = "Cluster"
	RedisModeSentinel = "Sentinel"
)

type radixClient interface {
	Do(context.Context, radix.Action) error
	Close() error
}

type redisGoClient interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
	Close() error
}

type redisGoClientBuilder interface {
	Get() redisGoClient
	Close() error
}

type redisGoSingleClientBuilder struct {
	pool *redigo.Pool
}

func (r *redisGoSingleClientBuilder) Get() redisGoClient { return r.pool.Get() }
func (r *redisGoSingleClientBuilder) Close() error       { return r.pool.Close() }

type redisGoClusterClient struct {
	cluster *redisgocluster.Cluster
}

func (r *redisGoClusterClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return r.cluster.Do(commandName, args...)
}
func (r *redisGoClusterClient) Close() error { return nil }

type redisGoClusterClientBuilder struct {
	c *redisGoClusterClient
}

func (r *redisGoClusterClientBuilder) Get() redisGoClient { return r.c }
func (r *redisGoClusterClientBuilder) Close() error       { r.c.cluster.Close(); return nil }

func newRedissonWithFlushAll(cfg *config.Config, bench Benchmark, resp redisson.RESP) (redisson.Cmdable, error) {
	c, err := redisson.Connect(redisson.NewConf(
		redisson.WithAddrs(cfg.GetRedis().GetAddrs()...),
		redisson.WithCluster(cfg.GetRedis().GetCluster()),
		redisson.WithDevelopment(false),
		redisson.WithConnPoolSize(bench.PoolSize),
		redisson.WithEnableMonitor(false),
		redisson.WithResp(resp),
		redisson.WithMasterName(cfg.GetRedis().GetMasterName()),
	))
	if err != nil {
		return nil, err
	}
	err = c.FlushAll(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Benchmark struct {
	Key           string
	Val           string
	PoolSize      int
	TargetBuilder TargetBuilder
	RedisMode     RedisMode
	Mode          redisConfig.Mode
}

type TargetAction func(key, value string) error

type Target struct {
	Close func()
	Do    TargetAction
}

type TargetBuilder struct {
	Name  string
	Check func(Benchmark) bool
	Make  func(Benchmark) (Target, error)
}

func newRedissonRESP2TargetBuilder(cfg *config.Config, before func(redisson.Cmdable, Benchmark) error, make func(redisson.Cmdable) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/sandwich-go/redisson/RESP2",
		Make: func(benchmark Benchmark) (Target, error) {
			client, err := newRedissonWithFlushAll(cfg, benchmark, redisson.RESP2)
			if err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client),
			}, nil
		},
	}
}

func newRedissonRESP3CacheTargetBuilder(cfg *config.Config, before func(redisson.Cmdable, Benchmark) error, make func(redisson.CacheCmdable) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/sandwich-go/redisson/RESP3/cache",
		Make: func(benchmark Benchmark) (Target, error) {
			client, err := newRedissonWithFlushAll(cfg, benchmark, redisson.RESP3)
			if err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client.Cache(10 * time.Minute)),
			}, nil
		},
	}
}

func newRedissonRESP3TargetBuilder(cfg *config.Config, before func(redisson.Cmdable, Benchmark) error, make func(redisson.Cmdable) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/sandwich-go/redisson/RESP3",
		Make: func(benchmark Benchmark) (Target, error) {
			client, err := newRedissonWithFlushAll(cfg, benchmark, redisson.RESP3)
			if err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client),
			}, nil
		},
	}
}

func newRueidisWithFlushAll(cfg *config.Config) (rueidiscompat.Cmdable, rueidis.Client, error) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{
			InitAddress: cfg.GetRedis().GetAddrs(),
			ShuffleInit: true,
			Sentinel: rueidis.SentinelOption{
				MasterSet: cfg.GetRedis().GetMasterName(),
			},
		},
	)
	if err != nil {
		return nil, nil, err
	}
	if err = client.Do(context.Background(), client.B().Flushall().Build()).Error(); err != nil {
		return nil, nil, err
	}
	return rueidiscompat.NewAdapter(client), client, err
}

func newRueidisCacheTargetBuilder(cfg *config.Config, before func(rueidiscompat.Cmdable, Benchmark) error, make func(rueidiscompat.CacheCompat) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/rueian/rueidis/rueidiscompat/cache",
		Make: func(benchmark Benchmark) (Target, error) {
			adapter, client, err := newRueidisWithFlushAll(cfg)
			if err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(adapter, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: client.Close,
				Do:    make(adapter.Cache(10 * time.Minute)),
			}, nil
		},
	}
}

func newRueidisTargetBuilder(cfg *config.Config, before func(rueidiscompat.Cmdable, Benchmark) error, make func(rueidiscompat.Cmdable) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/rueian/rueidis/rueidiscompat",
		Make: func(benchmark Benchmark) (Target, error) {
			adapter, client, err := newRueidisWithFlushAll(cfg)
			if err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(adapter, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: client.Close,
				Do:    make(adapter),
			}, nil
		},
	}
}

func newGoRedisTargetBuilder(cfg *config.Config, before func(redis.UniversalClient, Benchmark) error, make func(redis.UniversalClient) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/go-redis/redis/v8",
		Make: func(benchmark Benchmark) (Target, error) {
			var client redis.UniversalClient
			switch benchmark.RedisMode {
			case RedisModeSingle:
				client = redis.NewClient(&redis.Options{Addr: cfg.GetRedis().GetAddrs()[0], PoolSize: benchmark.PoolSize})
			case RedisModeSentinel:
				client = redis.NewUniversalClient(&redis.UniversalOptions{Addrs: cfg.GetRedis().GetAddrs(), MasterName: cfg.GetRedis().GetMasterName(), PoolSize: benchmark.PoolSize})
			case RedisModeCluster:
				client = redis.NewClusterClient(&redis.ClusterOptions{Addrs: cfg.GetRedis().GetAddrs(), PoolSize: benchmark.PoolSize})
			}
			if err := client.FlushAll(context.Background()).Err(); err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err := before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client),
			}, nil
		},
	}
}

func newRedigoTargetBuilder(cfg *config.Config, before func(redisGoClientBuilder, Benchmark) error, make func(redisGoClientBuilder) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/garyburd/redigo",
		Make: func(benchmark Benchmark) (Target, error) {
			var client redisGoClientBuilder
			switch benchmark.RedisMode {
			case RedisModeSingle:
				client = &redisGoSingleClientBuilder{pool: &redigo.Pool{
					Dial: func() (redigo.Conn, error) {
						return redigo.Dial("tcp", cfg.GetRedis().GetAddrs()[0])
					},
					MaxIdle:   benchmark.PoolSize,
					MaxActive: benchmark.PoolSize,
					Wait:      true,
				}}
			case RedisModeCluster:
				cluster, err1 := redisgocluster.NewCluster(
					&redisgocluster.Options{
						StartNodes:  cfg.GetRedis().GetAddrs(),
						ConnTimeout: time.Minute,
						KeepAlive:   128,
						AliveTime:   time.Minute,
					})
				if err1 != nil {
					return emptyTarget, err1
				}
				client = &redisGoClusterClientBuilder{c: &redisGoClusterClient{cluster: cluster}}
			}
			conn := client.Get()
			if _, err := conn.Do("FLUSHALL"); err != nil {
				_ = conn.Close()
				return emptyTarget, err
			}
			_ = conn.Close()
			if before != nil {
				if err := before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client),
			}, nil
		},
		Check: func(benchmark Benchmark) bool {
			return benchmark.RedisMode != RedisModeSentinel
		},
	}
}

func newRadixTargetBuilder(cfg *config.Config, before func(radixClient, Benchmark) error, make func(radixClient) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/mediocregopher/radix/v4",
		Make: func(benchmark Benchmark) (Target, error) {
			var err error
			var client radixClient
			var ctx = context.Background()
			switch benchmark.RedisMode {
			case RedisModeSingle:
				client, err = (radix.PoolConfig{Size: benchmark.PoolSize}).New(ctx, "tcp", cfg.GetRedis().GetAddrs()[0])
			case RedisModeSentinel:
				client, err = radix.SentinelConfig{PoolConfig: radix.PoolConfig{Size: benchmark.PoolSize}}.New(ctx, cfg.GetRedis().GetMasterName(), cfg.GetRedis().GetAddrs())
			case RedisModeCluster:
				client, err = radix.ClusterConfig{PoolConfig: radix.PoolConfig{Size: benchmark.PoolSize}}.New(ctx, cfg.GetRedis().GetAddrs())
			}
			if err != nil {
				return emptyTarget, err
			}
			if err = client.Do(ctx, radix.Cmd(nil, "FLUSHALL")); err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: func() {
					_ = client.Close()
				},
				Do: make(client),
			}, nil
		},
	}
}

func newRedispipeTargetBuilder(cfg *config.Config, before func(redispipe.Sender, Benchmark) error, make func(redispipe.Sender) TargetAction) TargetBuilder {
	return TargetBuilder{
		Name: "github.com/joomcode/redispipe",
		Make: func(benchmark Benchmark) (Target, error) {
			var err error
			var client redispipe.Sender
			var ctx = context.Background()
			switch benchmark.RedisMode {
			case RedisModeSingle:
				client, err = redisconn.Connect(ctx, cfg.GetRedis().GetAddrs()[0], redisconn.Opts{Logger: redisconn.NoopLogger{}})
			case RedisModeCluster:
				client, err = rediscluster.NewCluster(ctx, cfg.GetRedis().GetAddrs(), rediscluster.Opts{Logger: rediscluster.NoopLogger{}})
			}
			if err != nil {
				return emptyTarget, err
			}
			if err = redispipe.AsError(redispipe.Sync{client}.Do("FLUSHALL")); err != nil {
				return emptyTarget, err
			}
			if before != nil {
				if err = before(client, benchmark); err != nil {
					return emptyTarget, err
				}
			}
			return Target{
				Close: client.Close,
				Do:    make(client),
			}, nil
		},
		Check: func(benchmark Benchmark) bool {
			return benchmark.RedisMode != RedisModeSentinel
		},
	}
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(fastrand.Uint32n(26) + 'a'))
	}
	return sb.String()
}

func readConf() (*config.Config, error) {
	return config.ReadConfigFiles("./setting.yaml")
}

func getRedisMode(c *config.Config) RedisMode {
	if len(c.GetRedis().GetMasterName()) > 0 {
		return RedisModeSentinel
	}
	if c.GetRedis().GetCluster() {
		return RedisModeCluster
	}
	return RedisModeSingle
}

func getRunMode(c *config.Config) redisConfig.Mode {
	switch strings.ToLower(c.GetRedis().GetMode()) {
	case redisConfig.ModeParallel:
		return redisConfig.ModeParallel
	default:
		return redisConfig.ModeSerial
	}
}

func compose(c *config.Config, builders []TargetBuilder) []Benchmark {
	redisMode := getRedisMode(c)
	benchmarks := make([]Benchmark, 0, len(c.GetRedis().GetKeySizes())*len(c.GetRedis().GetValueSizes())*len(builders))
	for _, k := range c.GetRedis().GetKeySizes() {
		key := gen(k)
		for _, v := range c.GetRedis().GetValueSizes() {
			val := gen(v)
			for _, poolSize := range c.GetRedis().GetPoolSizes() {
				for _, builder := range builders {
					benchmarks = append(benchmarks, Benchmark{Key: key, Val: val, PoolSize: poolSize, TargetBuilder: builder, RedisMode: redisMode, Mode: getRunMode(c)})
				}
			}
		}
	}
	return benchmarks
}

func runBenchmark(b *testing.B, benchmarks []Benchmark) {
	for _, bench := range benchmarks {
		bc := bench
		if bc.TargetBuilder.Check != nil {
			if !bc.TargetBuilder.Check(bc) {
				continue
			}
		}
		var parallel = 0
		var modeString string
		switch bc.Mode {
		case redisConfig.ModeSerial:
			modeString = "Serial"
		case redisConfig.ModeParallel:
			parallel = runtime.GOMAXPROCS(0) * 8
			modeString = fmt.Sprintf("Parallel(%d)", parallel)
		}
		b.Run(fmt.Sprintf("%s:%s:Key(%d):Value(%d):Pool(%d):%s", bc.TargetBuilder.Name, bc.RedisMode, len(bc.Key), len(bc.Val), bc.PoolSize, modeString), func(b *testing.B) {
			target, err := bc.TargetBuilder.Make(bc)
			if err != nil {
				b.Fatalf("%s setup fail: %v", bc.TargetBuilder.Name, err)
			}
			if parallel > 0 {
				b.SetParallelism(parallel)
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						if err = target.Do(bench.Key, bench.Val); err != nil {
							b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
						}
					}
				})
			} else {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if err = target.Do(bench.Key, bench.Val); err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}
				}
			}

			b.StopTimer()
			target.Close()
		})
	}
}
