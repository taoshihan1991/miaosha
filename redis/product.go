package redis

func ProductInfo(id string) map[string]string {
	key := "product:" + id
	return HashGet(key)
}
