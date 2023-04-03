package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// welcome to your channel go guru

// go mod init goguru
// go mod tidy

// Topic:- what is redis?

// how to connect with redis using golang
// how to do some basic operations with redis using golang

type Product struct {
	Id             int64
	Name           string `json:"name"`
	ProductType    string `json:"product_type"`
	ProductQuality string `json:"product_quality"`
}

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	products := Product{
		Id:             time.Now().Unix(),
		Name:           "product 2",
		ProductType:    "home",
		ProductQuality: "good",
	}

	obj, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}
	keyStr := strconv.Itoa(int(products.Id))
	err = client.Set(keyStr, obj, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(keyStr).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	var pro Product
	err = json.Unmarshal([]byte(val), &pro)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Del("2").Result()
	fmt.Println(resp, err)
}
