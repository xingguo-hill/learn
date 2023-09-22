package x

import (
	"context"
	_ "embed"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed ratelimit.lua
var testLuaScript string

func TestRateLimit(t *testing.T) {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.56.30:6379",
	})
	// var ctx = context.Background()
	key := fmt.Sprintf("%s:%s", "ip-limiter", "127.0.0.1")
	for i := 0; i < 105; i++ {
		limit, _ := redisClient.Eval(context.Background(), testLuaScript, []string{key},
			time.Second, 100, time.Now().UnixMilli()).Bool()
		if limit {
			t.Log("阈值生效")
			fmt.Println("阈值生效")
			return
		}
	}
}
