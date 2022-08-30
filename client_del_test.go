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

// go test -bench=BenchmarkClientDel -benchmem -benchtime 2s -timeout 60m .
func BenchmarkClientDel(b *testing.B) {
	cfg, err := readConf()
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.Background()
	runBenchmark(b, compose(cfg, []TargetBuilder{
		newRedissonRESP2TargetBuilder(cfg, nil, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Del(ctx, key).Err()
			}
		}),
		newRedissonRESP3TargetBuilder(cfg, nil, func(client redisson.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Del(ctx, key).Err()
			}
		}),
		newRueidisTargetBuilder(cfg, nil, func(client rueidiscompat.Cmdable) TargetAction {
			return func(key string, value string) error {
				return client.Del(ctx, key).Err()
			}
		}),
		newGoRedisTargetBuilder(cfg, nil, func(client redis.UniversalClient) TargetAction {
			return func(key string, value string) error {
				return client.Del(ctx, key).Err()
			}
		}),
		newRadixTargetBuilder(cfg, nil, func(client radixClient) TargetAction {
			return func(key string, value string) error {
				return client.Do(ctx, radix.Cmd(nil, "DEL", key))
			}
		}),
		newRedispipeTargetBuilder(cfg, nil, func(client redispipe.Sender) TargetAction {
			return func(key string, value string) error {
				return redispipe.AsError(redispipe.Sync{client}.Do("DEL", key))
			}
		}),
	}))
}
