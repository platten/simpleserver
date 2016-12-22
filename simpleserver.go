package main

import (
    "fmt"
    "net/http"
    "log"
)

func localIp(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "10.195.225.205")
}


func publicIp(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "108.171.135.164")
}

func main() {
    http.HandleFunc("/local-ipv4", localIp)
    http.HandleFunc("/public-ipv4", publicIp)
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}