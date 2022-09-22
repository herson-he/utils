package taoredis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func CreateClient(addr string, password string, db int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	Rdb = rdb
}

const KeyUserData = "userdata"
