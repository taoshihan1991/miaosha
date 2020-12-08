package redis

import "time"

func GetUserInfo(session string) string {
	key := "session_" + session
	return GetStr(key)
}
func SetUserInfo(session string, str string) string {
	key := "session_" + session
	return SetStr(key, str, 30*time.Minute)
}
