package redis

import (
	"encoding/json"
	"log"
	"time"
)

type Order struct {
	Time      time.Time `json:"time"`
	User      string    `json:"user"`
	ProductId string    `json:"productid"`
}

func InsertOrder(user string, productId string) {
	order := Order{
		Time:      time.Now(),
		User:      user,
		ProductId: productId,
	}
	msg, err := json.Marshal(order)
	if err != nil {
		log.Println(err.Error())
		return
	}
	SortedSetAdd("orders", string(msg), float64(time.Now().Unix()))
}
func GetOrders() []string {
	list := SortedSetList("orders", 0, 10)
	return list
}
