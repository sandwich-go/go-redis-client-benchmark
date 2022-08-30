package go_redis_client_benchmark

import (
	"context"
	"github.com/go-redis/redis/v8"
	redispipe "github.com/joomcode/redispipe/redis"
	"github.com/mediocregopher/radix/v4"
	"github.com/rueian/rueidis/rueidiscompat"
	"github.com/sandwich-go/redisson"
	"testing"
)

// go test -bench=BenchmarkClientHGet -benchmem -benchtime 2s -timeout 60m .
func BenchmarkClientHGet(b *testing.B) {
	cfg, err := readConf()
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.Background()
	runBenchmark(b, compose(cfg, []TargetBuilder{
		newRedissonRESP2TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newRedissonRESP3TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newRedissonRESP3CacheTargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client redisson.CacheCmdable) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newRueidisTargetBuilder(cfg, func(client rueidiscompat.Cmdable, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client rueidiscompat.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newRueidisCacheTargetBuilder(cfg, func(client rueidiscompat.Cmdable, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client rueidiscompat.CacheCompat) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newGoRedisTargetBuilder(cfg, func(client redis.UniversalClient, benchmark Benchmark) error {
			return client.HSet(ctx, benchmark.Key, benchmark.Key, benchmark.Val).Err()
		}, func(client redis.UniversalClient) TargetAction {
			return func(key string, value string) error {
				return client.HGet(ctx, key, key).Err()
			}
		}),
		newRadixTargetBuilder(cfg, func(client radixClient, benchmark Benchmark) error {
			return client.Do(ctx, radix.Cmd(nil, "HSET", benchmark.Key, benchmark.Key, benchmark.Val))
		}, func(client radixClient) TargetAction {
			return func(key string, value string) error {
				return client.Do(ctx, radix.Cmd(nil, "HGET", key, key))
			}
		}),
		newRedispipeTargetBuilder(cfg, func(client redispipe.Sender, benchmark Benchmark) error {
			return redispipe.AsError(redispipe.Sync{client}.Do("HSET", benchmark.Key, benchmark.Key, benchmark.Val))
		}, func(client redispipe.Sender) TargetAction {
			return func(key string, value string) error {
				return redispipe.AsError(redispipe.Sync{client}.Do("HGET", key, key))
			}
		}),
	}))
}
