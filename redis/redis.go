package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/taoshihan1991/miaosha/setting"
	"log"
	"strconv"
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
func HashGet(key string, field string) string {
	str, err := rdb.HGet(ctx, key, field).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return str
}
func HashGetAll(key string) map[string]string {
	res, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return make(map[string]string)
	}
	return res
}
func HashExist(key string, field string) bool {
	bool, err := rdb.HExists(ctx, key, field).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}
func HashSetV4(key string, values ...interface{}) (int64, error) {
	num, err := rdb.HSet(ctx, key, values).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return num, err
}
func HashSetV3(key string, values ...interface{}) (bool, error) {
	res, err := rdb.HMSet(ctx, key, values).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return res, err
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
func ListPush(key string, value interface{}) {
	_, err := rdb.LPush(ctx, key, value).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func ListPop(key string) string {
	res, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return res
}
func ListLen(key string) int64 {
	res, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return res
}
func ListIndex(key string, index int64) string {
	res, err := rdb.LIndex(ctx, key, index).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return res
}
func DelKey(key string) {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
	}
}
func LimitFreqs(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	length := uint(ListLen(queueName))
	if length < count {
		ListPush(queueName, currTime)
		return true
	}
	//队列满了,取出最早访问的时间
	earlyTime, _ := strconv.ParseInt(ListIndex(queueName, int64(length)-1), 10, 64)
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		ListPop(queueName)
		ListPush(queueName, currTime)
	}
	return true
}
