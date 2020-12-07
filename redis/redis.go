package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/taoshihan1991/miaosha/setting"
	"log"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

func NewRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Ip + ":" + setting.Redis.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func GetStr(key string) string {
	str, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return str
}
func SetStr(key string, value interface{}, expire time.Duration) string {
	str, err := rdb.Set(ctx, key, value, expire).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return str
}
func HashGet(key string) map[string]string {
	res, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return make(map[string]string)
	}
	return res
}
func HashSetV4(key string, values ...interface{}) {
	rdb.HSet(ctx, key, values)
}
func HashSetV3(key string, values ...interface{}) {
	_, err := rdb.HMSet(ctx, key, values).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func DelKey(key string) {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func HashInc(key string, field string, inc int64) {
	_, err := rdb.HIncrBy(ctx, key, field, inc).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func SortedSetAdd(key string, member interface{}, score float64) {
	z := &redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := rdb.ZAdd(ctx, key, z).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func SortedSetList(key string, start int64, stop int64) []string {
	list, err := rdb.ZRange(ctx, key, start, stop).Result()
	log.Println(list, err)
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}
	return list
}
