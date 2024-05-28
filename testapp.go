package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

const portNum string = ":8080"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
        hostname, err := os.Hostname()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Fprintf(w, "Hostname: %s\n", hostname)
    })

    log.Println("Started on port", portNum)
    err := http.ListenAndServe(portNum, nil)
    if err != nil {
        log.Fatal(err)
    }
}