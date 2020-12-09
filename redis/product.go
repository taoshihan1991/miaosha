package redis

import (
	"fmt"
	"time"
)

func ProductInfo(id string) map[string]string {
	key := "product:" + id
	return HashGetAll(key)
}
func SetProduct(id string) {
	key := "product:" + id
	saleTime := fmt.Sprintf("%d", time.Now().UnixNano()/1e6+1000*60)
	HashSetV3(key, "title", "小米100高贵尊享版", "price", "1", "saletime", saleTime, "storge", 5)
}
func DecProductStorge(id string) int64 {
	key := "product:" + id
	inc := HashInc(key, "storge", -1)
	return inc
}
func PushRequestQueue(item string) {
	key := "product_request_queue"
	ListPush(key, item)
}
func PopRequestQueue() string {
	key := "product_request_queue"
	return ListPop(key)
}
func LenRequestQueue() int64 {
	key := "product_request_queue"
	return ListLen(key)
}
