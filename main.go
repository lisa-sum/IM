package main

import (
	"flag"
	"github.com/joho/godotenv"
	"im/db"
	"im/router"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:4000", "http service address")

func main() {
	err := godotenv.Load("config/db.yaml", "config/minio.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db.RedisDBInit()
	db.MongoDBInit()
	router.Server() // 运行web服务
	log.Fatal(http.ListenAndServe(*addr, nil))
}
