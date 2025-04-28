package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	GetData(key string) (string, error)
	SetData(key string, val string) error
	HSetData(key, string, field string, val string) error
	SetRedisConn()
}

type RedisClient struct {
	Client *redis.Client
}

var (
	ctx = context.Background()
)

func NewRedisClient() *RedisClient {

	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:            "localhost:6379", // 접근 url 및 port
			Password:        "",               // password ""값은 없다는 뜻
			DB:              0,                // 기본 DB 사용
			ConnMaxIdleTime: 30 * time.Minute,
			ConnMaxLifetime: 30 * time.Minute,
			MaxIdleConns:    1000,
			PoolSize:        25,
		}),
	}
}

func (c *RedisClient) GetData(key string) (string, error) {

	result, err := c.Client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return result, nil

}

func (c *RedisClient) SetData(key string, val interface{}) error {

	err := c.Client.Set(ctx, key, val, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (c *RedisClient) HSetData(key string, field string, val string) error {

	v := c.Client.HSet(ctx, key, field, val).Val()
	fmt.Println("#########################3")
	fmt.Println(v)
	fmt.Println("#########################3")
	/*
		if err != nil {
			return err
		}*/

	return nil
}

func (c *RedisClient) HGetData(key string, field string) (string, error) {

	str, err := c.Client.HGet(ctx, key, field).Result()

	if err != nil {
		return "", err
	}

	return str, nil
}

/*
func GetAllData(key string) (string, error) {

	c := RedisClient.NewRedisClient

	result, err := c.client.Get(key).Result()

	if err != nil {
		return "", err
	}

	return result, nil

}
*/
