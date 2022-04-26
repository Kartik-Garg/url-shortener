package database

import(
	"context"
	"github.com/go-redis/redis/v8"
	"os"

)
/*context are used to add "context"/ properties to the project
like a timeout, deadline, etc. for eg we can have a timeout conetxt, where if an API does not respond
for a long time, it will end execution 
@conext.Background - used to return the empty context where values can actually be added
*/ 
var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		//to create a new db client we need to pass 3 things to redis
		Addr:		os.Getenv("DB_ADDR"),
		Password:	os.Getenv("DB_PASS"),
		DB:			dbNo,
	})

	return rdb
}