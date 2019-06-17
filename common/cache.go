package common

import (
	"log"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	redisClient *redis.Client
}

var redisclient *redis.Client

func InitCache(redis_dsn string) *redis.Client {

	redisclient = redis.NewClient(&redis.Options{
		//Addr: os.Getenv("REDIS_DSN"), // TODO set for config
		Addr:     redis_dsn, // TODO set for config
		Password: "",
		DB:       0,
	})

	return redisclient
}

func SetItem(key string, v interface{}) error {

	err := redisclient.Set(key, v, 0).Err()

	if err != nil {
		log.Println(err)
		return err
	}
	return err

}

func GetItem(key string) (string, error) {

	val, err := redisclient.Get(key).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return val, err

}

func GetCache() *redis.Client {
	return redisclient
}
