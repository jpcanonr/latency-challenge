# Escenario 2: Docker + HTTP + Go

Este documento describe la arquitectura y el flujo de comunicación del segundo escenario de pruebas de latencia.

---

## 1. Componentes principales

- **Cliente Go (host):**
  - Programa escrito en Go que envía solicitudes HTTP.
  - Corre en la máquina host (fuera del contenedor).

- **Servidor Go (contenedor Docker):**
  - Programa escrito en Go que expone un endpoint HTTP.
  - Corre dentro de un contenedor Docker.
  - Escucha en el puerto `8080`.

- **Conexión HTTP:**
  - Canal de comunicación entre el cliente y el servidor.
  - Flujo de datos:  
    - El cliente envía una solicitud `GET`.  
    - El servidor responde con `"respuesta"`.

---

## 2. Flujo de comunicación

1. El **cliente Go** realiza una solicitud HTTP `GET` a `http://127.0.0.1:8080`.
2. El **servidor Go** dentro del contenedor recibe la solicitud.
3. El servidor responde con `"respuesta"` y código HTTP 200.
4. El cliente mide el tiempo transcurrido (latencia).
5. Se repite el proceso para múltiples iteraciones.

---

## 3. Arquitectura visual

```plaintext
+-------------------+       HTTP GET       +-------------------------+
|                   | -------------------> |                         |
|   Cliente Go      |                      |   Docker Container      |
|   (Host Machine)  | <------------------- |   Servidor Go (HTTP)    |
|                   |   HTTP 200 + body    |   Puerto 8080           |
+-------------------+                      +-------------------------+
```

---

## 4. Características del escenario

- **Motor:** Docker
- **Protocolo:** HTTP
- **Lenguaje:** Go
- **Puerto:** 8080
- **Métrica clave:** Latencia (ms)

---

## 5. Uso típico

- El cliente ejecuta un **benchmark.go** enviando 100 solicitudes TCP.
- Se calculan métricas de latencia: promedio, mínimo y máximo.
- Los resultados se registran en la tabla comparativa de escenarios.
