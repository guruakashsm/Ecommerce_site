	package main

import (
	"fmt"
	"log"
	"ecommerce/router"
)

func main() {
	fmt.Println("Started Running")
	r := router.Router()
	log.Fatal(r.Run(":8081"))
	fmt.Println("Listening At PORT ... ")
}