package main

import (
	"flag"
	"im/db"
	"im/router"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:4000", "http service address")

func main() {
	db.RedisDBInit()
	db.MongoDBInit()
	router.Server() // 运行web服务
	log.Fatal(http.ListenAndServe(*addr, nil))
}
