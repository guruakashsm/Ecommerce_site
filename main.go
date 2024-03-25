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

    
    certFile := "./localhost.crt"
    keyFile := "./decrypted_key.pem" 

    
    if _, err := os.Stat(certFile); os.IsNotExist(err) {
        log.Fatal("SSL certificate file not found")
    }
    
    if _, err := os.Stat(keyFile); os.IsNotExist(err) {
        log.Fatal("SSL key file not found")
    }

    // Get the port from the environment variable or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Run the server with TLS/SSL
    if err := r.RunTLS("0.0.0.0:"+port, certFile, keyFile); err != nil {
        log.Fatalf("Error starting HTTPS server: %v", err)
    }
}
