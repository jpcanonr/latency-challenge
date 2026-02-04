// client.go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {
    start := time.Now()
    resp, err := http.Get("http://127.0.0.1:8080")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    elapsed := time.Since(start)

    fmt.Printf("Respuesta: %s\n", string(body))
    fmt.Printf("Latencia: %v\n", elapsed)
}
