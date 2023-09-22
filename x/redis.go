package x

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

func TestRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.30:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	var ctx = context.Background()
	key := "test"
	err := rdb.Set(ctx, key, 1, 0).Err()
	if err != nil {
		panic(err)
	}
	id, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("Incr:%s:%d", key, id)
	time.Sleep(10 * time.Second)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	id = cast.ToInt64(val)
	logrus.Infof("get:%s:%d", key, id)
}
