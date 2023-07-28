package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	// client := redis.NewClient(&redis.Options{
	// 	Addr:      "charlesredis1.redis.cache.windows.net:6380",
	// 	Password:  "<redis-access-key>”,
	// 	DB:        0,
	// 	TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	// })

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"charlesredis1.redis.cache.windows.net:13000",
			"charlesredis1.redis.cache.windows.net:13001",
		},
		Password:  "<redis-access-key>”,
	})

	ctx := context.Background()

	// err := client.Set(ctx, "foo", "bar", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get(ctx, "foo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("foo", val)

	val, err := client.Get(ctx, "shard0key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("shard0key: ", val)

	val1, err := client.Get(ctx, "shard1key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("shard1key: ", val1)
}
