package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/taoshihan1991/miaosha/setting"
	"log"
	"sync"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()
var mutex sync.Mutex

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
func HashInc(key string, field string, inc int64) int64 {
	res, err := rdb.HIncrBy(ctx, key, field, inc).Result()
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return res
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
	res := rdb.ZRange(ctx, key, start, stop)
	log.Println(res)

	list, err := res.Result()
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}
	return list
}
func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	bool, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}
func UnLock(key string) int64 {
	nums, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
