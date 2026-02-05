## Requisitos

- Sistema Operativo: **macOS (Apple Silicon / Chip M)**

---

## Paso 1: Preparar el entorno

1. Instalar **Docker Desktop** en tu máquina y verificar la instalación:
   ```bash
   docker --version
   ```

2. Instalar Golang e inicializar (Crear archivo go.mod):
   ```bash
   go version
   go mod init server
   ```

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
   ```

---

## Paso 5: Ejecutar pruebas de latencia

1. Corre el cliente Go en tu máquina host.
   ```bash
   go run client.go
   ```

Se debe obtener una salida como esta:
   ```bash
   Respuesta: respuesta
   Latencia: 3.9675ms
   ```

---

## Paso 6: Ejecutar pruebas de latencia con benchmark

1. El archivo benchmark.go permite enviar 1000 peticiones y recibir en la salida

   ```bash
   Iteración 1: 1.029 ms - Respuesta: respuesta
   Iteración 2: 0.366 ms - Respuesta: respuesta
   ...
   Iteración 99: 0.683 ms - Respuesta: respuesta
   Iteración 100: 0.623 ms - Respuesta: respuesta
   
   📊 Resultados del benchmark
   Solicitudes: 100
   Latencia promedio: 0.843 ms
   Latencia mínima: 0.366 ms
   Latencia máxima: 3.171 ms
   ```

2. La forma de ejecutar el benchmark es:
   ```bash
   go run benchmark.go
   ```
