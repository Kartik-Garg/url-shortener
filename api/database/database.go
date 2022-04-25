package database

import(
	"context"
	"github.com/go-redis/redis/v8"
	"os"

)

var Ctx = context.Background()

func createClient(dbNo int) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		//to create a new db client we need to pass 3 things to redis
		Addr:		os.Getenv("DB_ADDR"),
		Password:	os.Getenv("DB_PASS"),
		DB:			dbNo,
	})

	return rdb
}