package utils

import (
	"im/db"
	"log"
)

func HMSet(key string, data map[string]any) (string, error) {
	result, err := db.Redis.HMSet(key, data).Result()
	if err != nil {
		log.Fatalln("保存消息失败: " + err.Error())
		return "", err
	}
	return result, nil
}
