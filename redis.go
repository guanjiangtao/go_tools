package go_utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

type GoRedis struct {
	Host     string        `json:"host"`     // 账号
	Password string        `json:"password"` // 密码
	DB       int           `json:"DB"`       // 数据库
	Rdb      *redis.Client `json:"rdb"`      // redis客户端
}

// NewClient redis客户端
func (db *GoRedis) NewClient() {
	db.Rdb = redis.NewClient(&redis.Options{
		Addr:     db.Host,
		Password: db.Password,
		DB:       db.DB,
	})
}

// HSet 新增数据
func (db *GoRedis) HSet(key string, field string, value string) {
	//redis的键的set方法
	err := db.Rdb.HSet(ctx, key, field, value).Err()
	if err != nil {
		panic(err)
	}
}

// HExists 判断Key是否存在
func (db *GoRedis) HExists(key string, value string) bool {
	val, err := db.Rdb.HExists(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// HDel 删除key
func (db *GoRedis) HDel(key string, field string) bool {
	err := db.Rdb.HDel(ctx, key, field).Err()
	if err != nil {
		return false
	}
	return true
}

// HGetAll 获取所有
func (db *GoRedis) HGetAll(key string) map[string]string {
	result, err := db.Rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}
	}
	return result
}

// HClear 一键删除所有的Key
func (db *GoRedis) HClear(key string) bool {
	err := db.Rdb.Del(ctx, key).Err()
	if err != nil {
		return false
	}
	return true
}

// HGet 获取对应的Key
func (db *GoRedis) HGet(key string, field string) string {
	result, err := db.Rdb.HGet(ctx, key, field).Result()
	if err != nil {
		return ""
	}
	return result
}

// Expire 设置超时时间
func (db *GoRedis) Expire(key string, expiration time.Duration) bool {
	result, err := db.Rdb.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false
	}
	return result
}

// HKeys 获取所有的key
func (db *GoRedis) HKeys(key string) []string {
	result, err := db.Rdb.HKeys(ctx, key).Result()
	if err != nil {
		return []string{}
	}
	return result
}
