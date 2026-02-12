// benchmark_persistente.go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    host := "127.0.0.1:8080"
    iteraciones := 100
    latencias := make([]float64, 0, iteraciones)

    // Abrir conexión persistente
    conn, err := net.Dial("tcp", host)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // Activar TCP_NODELAY para deshabilitar Nagle’s Algorithm
    if tcpConn, ok := conn.(*net.TCPConn); ok {
        tcpConn.SetNoDelay(true)
    }

    buffer := make([]byte, 1024)

    for i := 0; i < iteraciones; i++ {
        start := time.Now()
        conn.Write([]byte("estimulo"))
        n, _ := conn.Read(buffer)
        elapsed := time.Since(start).Seconds() * 1000 // ms

        latencias = append(latencias, elapsed)
        fmt.Printf("Iteración %d: %.3f ms - Respuesta: %s\n", i+1, elapsed, string(buffer[:n]))
    }

    // calcular métricas
    var suma, min, max float64
    min, max = latencias[0], latencias[0]
    for _, v := range latencias {
        suma += v
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    promedio := suma / float64(iteraciones)

    fmt.Println("\n📊 Resultados del benchmark optimizado")
    fmt.Printf("Solicitudes: %d\n", iteraciones)
    fmt.Printf("Latencia promedio: %.3f ms\n", promedio)
    fmt.Printf("Latencia mínima: %.3f ms\n", min)
    fmt.Printf("Latencia máxima: %.3f ms\n", max)
}
