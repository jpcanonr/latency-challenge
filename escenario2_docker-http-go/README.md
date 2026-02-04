## Requisitos

- Sistema Operativo: **macOS (Apple Silicon / Chip M)**

---

## Paso 1: Preparar el entorno

1. Instalar **Docker Desktop** en tu máquina y verificar la instalación:
   ```bash
   docker --version
2. Instalar Golang e inicializar (Crear archivo go.mod):
   ```bash
   go version
   go mod init server

---

## Paso 2: Crear el servidor HTTP en Go

1. Crear un archivo server.go
Este servidor escuchará permanentemente y responderá de inmediato al recibir un estímulo.

---

## Paso 3: Crear el cliente HTTP en Go

1. Crear un archivo client.go
Este cliente enviará el estímulo y medirá la latencia.

---

## Paso 4: Dockerizar el servidor

1. Crear un archivo Dockerfile para el servidor
2. Construye y corre el contenedor:
   ```bash
   docker build -t go-http-server .
   docker run -p 8080:8080 go-http-server

---

## Paso 5: Ejecutar pruebas de latencia

1. Corre el cliente Go en tu máquina host.
   ```bash
   go run client.go

---

## Paso 6: Optimizar para latencia mínima
Usa --network host para evitar overhead de NAT:
   ```bash
   docker run --network host go-http-server

---

## paso 7: Ejecutar pruebas de latencia
1. Corre el cliente Go en tu máquina host.
   ```bash
   go run client2.go

Nota: La diferencia entre el cliente 1 y el cliente 2 es que el primero debe conectarse a http://localhost:8080. Si se utiliza la opción "--network host", el cliente debe conectarse a http://127.0.0.1:8080 (no localhost, porque a veces localhost resuelve a IPv6 ::1 y falla).