package redis

import (
	"github.com/go-redis/redis/v7"
)

var RDBFollower *redis.Client
var RDBFollowing *redis.Client
var RDBFriend *redis.Client

func InitRedis() {
	RDBFollower = redis.NewClient(&redis.Options{
		DB: 0,
	})
	RDBFollower = redis.NewClient(&redis.Options{
		DB: 1,
	})
	RDBFriend = redis.NewClient(&redis.Options{
		DB: 2,
	})

}
