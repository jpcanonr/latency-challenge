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

## Paso 2: Crear el servidor TCP en Go

1. Crear un archivo server.go
Este servidor escuchará permanentemente y responderá de inmediato al recibir un estímulo.

---

## Paso 3: Crear el cliente TCP en Go

1. Crear un archivo client.go
Este cliente enviará el estímulo y medirá la latencia.

---

## Paso 4: Dockerizar el servidor

1. Crear un archivo Dockerfile para el servidor
2. Construye y corre el contenedor:
   ```bash
   docker build -t go-tcp-server .
   docker run -p 8080:8080 go-tcp-server

---

## Paso 5: Ejecutar pruebas de latencia

1. Corre el cliente Go en tu máquina host.
   ```bash
   go run client.go
