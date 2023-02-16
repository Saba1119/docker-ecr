package main

import (
    "log"
    "net/http"
    "time"
)

func timeout(w http.ResponseWriter, req *http.Request) {
    time.Sleep(5 * time.Second)
}

func server() {
    go func() {
        http.HandleFunc("/timeout", timeout)
        log.Fatal(http.ListenAndServe(":8090", nil))
    }()
}

func main() {
    server()

    httpClient := http.Client{Timeout: 2 * time.Second}
    if_ , err := httpClient.Get("http://localhost:8090/timeout"); err != nil {
        log.Fatal(err)
    }
}

