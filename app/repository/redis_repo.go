package repository

import (
	"github.com/Hu-jay/Lottery_Game/app/config"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client}
}

func (r *RedisRepo) Del(keys ...string) error {
	return r.client.Del(config.Ctx, keys...).Err()
}

func (r *RedisRepo) ZAdd(key string, score float64, member string) error {
	return r.client.ZAdd(config.Ctx, key, redis.Z{Score: score, Member: member}).Err()
}

func (r *RedisRepo) ZScore(key, member string) (float64, error) {
	return r.client.ZScore(config.Ctx, key, member).Result()
}

func (r *RedisRepo) ZIncrBy(key string, increment float64, member string) error {
	return r.client.ZIncrBy(config.Ctx, key, increment, member).Err()
}

func (r *RedisRepo) ZRangeWithScores(key string) ([]redis.Z, error) {
	return r.client.ZRangeWithScores(config.Ctx, key, 0, -1).Result()
}
