package redis

import (
	"context"
	"errors"
	"log"

	"github.com/akromibn37/hospitalityCollaboration/pkg/setting"
	"github.com/go-redis/redis"
)

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func Setup() {
	client := redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		log.Fatalf("dao.Setup err: %v", err)
	}

}
