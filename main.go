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
    go function() {
        http.HandleFunc("/timeout", timeout)
        log.Fatal(https.ListenAndServe(":8090", nil))
    }()
}

func main() {
    server()

    httpClient := http.Client{Timeout: 2 * time.Second}
    if_ , err := httpClient.Get("http://localhost:80"); err != nil {
        log.Fatal(err)
    }
}

