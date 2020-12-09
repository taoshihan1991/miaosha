package redis

import (
	"encoding/json"
	"log"
	"time"
)

type Order struct {
	Time      string `json:"time"`
	User      string `json:"user"`
	ProductId string `json:"productid"`
}

func InsertOrder(user string, productId string) {
	order := Order{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		User:      user,
		ProductId: productId,
	}
	msg, err := json.Marshal(order)
	if err != nil {
		log.Println(err.Error())
		return
	}
	SortedSetAdd("orders", string(msg), float64(time.Now().Unix()))
	SetOrderDetail(productId, user, order.Time)
}
func GetOrders(start int64, stop int64) []*Order {
	list := SortedSetList("orders", start, stop)
	rows := make([]*Order, 0)
	for _, r := range list {
		order := &Order{}
		json.Unmarshal([]byte(r), order)
		rows = append(rows, order)
	}
	return rows
}
func SetOrderDetail(pid string, user string, time string) {
	key := "order_detail:" + user
	HashSetV3(key, "productid", pid, "user", user, "ordertime", time)
}
func OrderInfo(user string) map[string]string {
	key := "order_detail:" + user
	return HashGetAll(key)
}
func OrderExist(user string) bool {
	key := "order_detail:" + user
	return HashExist(key, "ordertime")
}
func OrderDel() {
	key := "orders"
	itemKey := "order_detail:"
	list := SortedSetList(key, 0, -1)
	for _, r := range list {
		order := &Order{}
		json.Unmarshal([]byte(r), order)
		DelKey(itemKey + order.User)
	}
	DelKey(key)
}
