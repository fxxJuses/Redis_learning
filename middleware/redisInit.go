package middleware

import "github.com/go-redis/redis"

var Rdb *redis.Client


func RedisInit(){
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.110.162:6379",
		Password: "0102kyrie",
		DB:       0,
	})

}