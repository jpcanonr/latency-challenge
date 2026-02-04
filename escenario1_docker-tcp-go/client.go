package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    start := time.Now()
    conn.Write([]byte("estimulo"))

    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)
    elapsed := time.Since(start)

    fmt.Printf("Respuesta: %s\n", string(buffer[:n]))
    fmt.Printf("Latencia: %v\n", elapsed)
}
