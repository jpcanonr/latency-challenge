package main

import (
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    fmt.Println("Servidor TCP escuchando en puerto 8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)
    // Responder inmediatamente
    conn.Write([]byte("respuesta"))
    fmt.Printf("Recibido: %s\n", string(buffer[:n]))
}
