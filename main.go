package main

import (
    "fmt"
    "net"
)

func main() {
    ports := 0
    var i uint
    for i = 0; i < 1024 ; i++{
        addr := fmt.Sprintf("scanme.nmap.org:%d",i)
        conn, err := net.Dial("tcp", addr)
        if err == nil {
            fmt.Println(i, "Port  ouvert")
            conn.Close()
            ports = ports + 1
        } else {
            fmt.Println(i, "Port fermé")
        }
    }
    fmt.Println("Scan efectué")
    fmt.Println(ports," ports sont ouverts.")
    
}

