// server.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "respuesta")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor HTTP escuchando en puerto 8080...")
    http.ListenAndServe(":8080", nil)
}
