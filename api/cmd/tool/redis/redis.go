package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 建立 Redis 連線
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 請根據您的 Redis 伺服器設定進行調整
		Password: "",               // 如果有設定密碼，請在此處填寫
		DB:       0,                // 要操作的資料庫編號
	})

	// 使用 SCAN 指令遍歷所有鍵
	iter := client.Scan(context.Background(), 0, "*", 0).Iterator()
	for iter.Next(context.Background()) {
		key := iter.Val()

		// 取得鍵對應的值
		value, err := client.Get(context.Background(), key).Result()
		if err != nil {
			fmt.Printf("無法取得鍵 %s 的值：%s\n", key, err)
			continue
		}

		// 列印鍵和值
		fmt.Printf("鍵：%s，值：%s\n", key, value)
	}
	if err := iter.Err(); err != nil {
		fmt.Printf("遍歷鍵時發生錯誤：%s\n", err)
	}

	// 清空所有資料
	errFlush := client.FlushAll(context.Background()).Err()
	if errFlush != nil {
		fmt.Printf("清空 Redis 資料時發生錯誤：%s\n", errFlush)
	} else {
		fmt.Println("已成功清空 Redis 資料")
	}

	// 關閉 Redis 連線
	err := client.Close()
	if err != nil {
		fmt.Printf("關閉 Redis 連線時發生錯誤：%s\n", err)
	}
}
