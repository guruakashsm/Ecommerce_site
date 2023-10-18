package main

import (
	"ecommerce/router"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Started Running")
	r := router.Router()
	log.Fatal(r.Run("192.168.247.238:8080"))
	fmt.Println("Listening At PORT ... ")
}
