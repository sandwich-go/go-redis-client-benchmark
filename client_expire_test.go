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

// go test -bench=BenchmarkClientExpire -benchmem -benchtime 2s -timeout 60m .
func BenchmarkClientExpire(b *testing.B) {
	cfg, err := readConf()
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.Background()
	runBenchmark(b, compose(cfg, []TargetBuilder{
		newRedissonRESP2TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 0).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Expire(ctx, key, 60*time.Second).Err()
			}
		}),
		newRedissonRESP3TargetBuilder(cfg, func(client redisson.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 0).Err()
		}, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Expire(ctx, key, 60*time.Second).Err()
			}
		}),
		newRueidisTargetBuilder(cfg, func(client rueidiscompat.Cmdable, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 0).Err()
		}, func(client rueidiscompat.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Expire(ctx, key, 60*time.Second).Err()
			}
		}),
		newGoRedisTargetBuilder(cfg, func(client redis.UniversalClient, benchmark Benchmark) error {
			return client.Set(ctx, benchmark.Key, benchmark.Val, 0).Err()
		}, func(client redis.UniversalClient) TargetAction {
			return func(key string, value string) error {
				return client.Expire(ctx, key, 60*time.Second).Err()
			}
		}),
		newRadixTargetBuilder(cfg, func(client radixClient, benchmark Benchmark) error {
			return client.Do(ctx, radix.Cmd(nil, "SET", benchmark.Key, benchmark.Val))
		}, func(client radixClient) TargetAction {
			return func(key string, value string) error {
				return client.Do(ctx, radix.Cmd(nil, "EXPIRE", key, "60"))
			}
		}),
		newRedispipeTargetBuilder(cfg, func(client redispipe.Sender, benchmark Benchmark) error {
			return redispipe.AsError(redispipe.Sync{client}.Do("SET", benchmark.Key, benchmark.Val))
		}, func(client redispipe.Sender) TargetAction {
			return func(key string, value string) error {
				return redispipe.AsError(redispipe.Sync{client}.Do("EXPIRE", key, 60))
			}
		}),
	}))
}
