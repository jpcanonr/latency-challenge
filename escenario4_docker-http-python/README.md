## Requisitos

- Sistema Operativo: **macOS (Apple Silicon / Chip M)**

---

## Paso 1: Preparar el entorno

1. Instalar **Docker Desktop** en tu máquina y verificar la instalación:
   ```bash
   docker --version
2. Instalar python:
   ```bash
   python3 --version
3. Nota: Necesitas instalar requests si no lo tienes:
   ```bash
   pip3 install requests

---

## Paso 2: Crear el servidor HTTP en Python

1. Crear un archivo server.py
Este servidor escuchará permanentemente y responderá de inmediato al recibir un estímulo.

---

## Paso 3: Crear el cliente HTTP en Python

1. Crear un archivo client.py
Este cliente enviará el estímulo y medirá la latencia.

---

## Paso 4: Dockerizar el servidor

1. Crear un archivo Dockerfile para el servidor
2. Construye y corre el contenedor:
   ```bash
   docker build -t python-http-server .
   docker run -p 8080:8080 python-http-server

---

## Paso 5: Ejecutar pruebas de latencia

1. Corre el cliente Python en tu máquina host.
   ```bash
   python3 client.py

---

## Paso 6: Optimizar para latencia mínima
Usa --network host para evitar overhead de NAT:
   ```bash
   docker run --network host python-http-server

---

## paso 7: Ejecutar pruebas de latencia
1. Corre el cliente Go en tu máquina host.
   ```bash
   python3 client.py

Nota:
Con docker run -p 8080:8080 ... → el contenedor expone el puerto 8080 al host, y el cliente se conecta a 127.0.0.1:8080.

Con docker run --network host ... → el contenedor usa la red del host directamente.

En Linux: el servidor dentro del contenedor escucha en el host real en 0.0.0.0:8080. El cliente puede seguir usando 127.0.0.1:8080.

En macOS y Windows: --network host no funciona igual. Docker Desktop usa una VM interna, y el contenedor no comparte la red del host. Por eso el cliente no encuentra nada y lanza ConnectionRefusedError.