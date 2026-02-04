# 🚀 Reto de Latencia en Arquitecturas de Software

Este repositorio contiene 8 escenarios experimentales para medir la **latencia de comunicación** en sistemas diseñados para responder a estímulos en el menor tiempo posible.  
El objetivo es comparar diferentes combinaciones de **motor de ejecución (Docker/Kubernetes)**, **protocolo (TCP/HTTP)** y **lenguaje (Go/Python)**.

---

## 🎯 Objetivo del reto
Diseñar y construir un sistema que, ante un estímulo (ejemplo: un mensaje), responda con otro mensaje (ejemplo: `"respuesta"`) en el menor tiempo posible.  La meta es lograr latencias **menores a 1 ms** en condiciones óptimas.

---

## 🎯 Requisitos básicos

- El sistema debe escuchar permanentemente peticiones o estímulos.  
- Al recibir el estímulo, el sistema debe retornar una respuesta específica.  
- Debe medirse el tiempo transcurrido desde el momento en que se envía el estímulo hasta que se recibe la respuesta.


---

## 📂 Estructura del repositorio
Cada subcarpeta corresponde a un escenario:
```text
.
├── escenario1_docker-tcp-go
├── escenario2_docker-http-go
├── escenario3_docker-tcp-python
├── escenario4_docker-http-python
├── escenario5_k8s-tcp-go
├── escenario6_k8s-http-go
├── escenario7_k8s-tcp-python
└── escenario8_k8s-http-python
```

Dentro de cada carpeta encontrarás:
- `server.*` → Código del servidor que escucha estímulos.
- `client.*` → Código del cliente que envía estímulos y mide latencia.
- `Dockerfile` → Imagen para ejecutar el servidor en contenedor.
- `benchmark.py` → Código que envía 100o peticiones, y da a la salida el máximo, m´nimo y promedio de latencia.
- `README.md` → Instrucciones específicas de ejecución para ese escenario.

---

## 🧪 Cómo ejecutar un escenario
Dentro de cada subcarpeta se va a encontrar un archivo README.md con instrucciones específicas.
Ejemplo con **Docker + TCP + Go** (`docker-tcp-go`):

1. Construir la imagen:
   ```bash
   docker build -t go-tcp-server .