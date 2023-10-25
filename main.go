package main

import (
    "ecommerce/router"
    "fmt"
    "log"
    "os"
)

func main() {
    fmt.Println("Started Running")
    r := router.Router()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to 8080 if PORT is not set
    }
    log.Fatal(r.Run("0.0.0.0:" + port))
}
