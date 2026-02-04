// benchmark.go
package main

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
)

func main() {
    url := "http://127.0.0.1:8080"
    iteraciones := 100
    latencias := make([]float64, 0, iteraciones)

    for i := 0; i < iteraciones; i++ {
        start := time.Now()
        resp, err := http.Get(url)
        if err != nil {
            panic(err)
        }
        _, _ = ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        elapsed := time.Since(start).Seconds() * 1000 // ms
        latencias = append(latencias, elapsed)
        fmt.Printf("Iteración %d: %.3f ms\n", i+1, elapsed)
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

    fmt.Println("\n📊 Resultados del benchmark")
    fmt.Printf("Solicitudes: %d\n", iteraciones)
    fmt.Printf("Latencia promedio: %.3f ms\n", promedio)
    fmt.Printf("Latencia mínima: %.3f ms\n", min)
    fmt.Printf("Latencia máxima: %.3f ms\n", max)
}
