package redis

import (
	"fmt"
	"time"
)

func ProductInfo(id string) map[string]string {
	key := "product:" + id
	return HashGet(key)
}
func SetProduct(id string) {
	key := "product:" + id
	saleTime := fmt.Sprintf("%d", time.Now().UnixNano()/1e6+1000*20)
	HashSetV3(key, "title", "小米100高贵尊享版", "price", "1", "saletime", saleTime, "storge", "100")
}
