package go_redis_client_benchmark

import (
	"context"
	"github.com/go-redis/redis/v8"
	redispipe "github.com/joomcode/redispipe/redis"
	"github.com/mediocregopher/radix/v4"
	"github.com/rueian/rueidis/rueidiscompat"
	"github.com/sandwich-go/redisson"
	"testing"
	"time"
)

// go test -bench=BenchmarkClientTTL -benchmem -benchtime 2s -timeout 60m .
func BenchmarkClientTTL(b *testing.B) {
	cfg, err := readConf()
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.Background()
	runBenchmark(b, compose(cfg, []TargetBuilder{
		newRedissonRESP2TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 3600*time.Second).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.TTL(ctx, key).Err()
			}
		}),
		newRedissonRESP3TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 3600*time.Second).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.TTL(ctx, key).Err()
			}
		}),
		newRueidisTargetBuilder(cfg, func(client rueidiscompat.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 3600*time.Second).Err()
		}, func(client rueidiscompat.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.TTL(ctx, key).Err()
			}
		}),
		newRueidisCacheTargetBuilder(cfg, func(client rueidiscompat.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 3600*time.Second).Err()
		}, func(client rueidiscompat.CacheCompat) TargetAction {
			return func(key string, value string) error {
				return client.TTL(ctx, key).Err()
			}
		}),
		newGoRedisTargetBuilder(cfg, func(client redis.UniversalClient, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 3600*time.Second).Err()
		}, func(client redis.UniversalClient) TargetAction {
			return func(key string, value string) error {
				return client.TTL(ctx, key).Err()
			}
		}),
		newRadixTargetBuilder(cfg, func(client radixClient, benchmark Benchmark) error {
			return client.Do(ctx, radix.Cmd(nil, "SETEX", benchmark.Key, "3600", benchmark.Val))
		}, func(client radixClient) TargetAction {
			return func(key string, value string) error {
				return client.Do(ctx, radix.Cmd(nil, "TTL", key))
			}
		}),
		newRedispipeTargetBuilder(cfg, func(client redispipe.Sender, benchmark Benchmark) error {
			return redispipe.AsError(redispipe.Sync{client}.Do("SETEX", benchmark.Key, 3600, benchmark.Val))
		}, func(client redispipe.Sender) TargetAction {
			return func(key string, value string) error {
				return redispipe.AsError(redispipe.Sync{client}.Do("TTL", key))
			}
		}),
	}))
}
