# Escenario 4: Docker + HTTP + Python

Este documento describe la arquitectura y el flujo de comunicación del cuarto escenario de pruebas de latencia.

---

## 1. Componentes principales

- **Cliente Python (host):**
  - Script en Python que envía solicitudes HTTP.
  - Corre en la máquina host (fuera del contenedor).

- **Servidor Python (contenedor Docker):**
  - Script en Python que expone un endpoint HTTP usando `http.server`.
  - Corre dentro de un contenedor Docker.
  - Escucha en el puerto `8080`.

- **Conexión HTTP:**
  - Canal de comunicación entre el cliente y el servidor.
  - Flujo de datos:  
    - El cliente envía una solicitud `GET`.  
    - El servidor responde con `"respuesta"` y código HTTP 200.

---

## 2. Flujo de comunicación

1. El **cliente Python** realiza una solicitud HTTP `GET` a `http://127.0.0.1:8080`.
2. El **servidor Python** dentro del contenedor recibe la solicitud.
3. El servidor responde con `"respuesta"` y código HTTP 200.
4. El cliente mide el tiempo transcurrido (latencia).
5. Se repite el proceso para múltiples iteraciones.

---

## 3. Arquitectura visual

```plaintext
+----------------------+       HTTP GET       +-------------------------+
|                      | -------------------> |                         |
|   Cliente Python     |                      |   Docker Container      |
|   (Host Machine)     | <------------------- |   Servidor Python (HTTP)|
|                      |   HTTP 200 + body    |   Puerto 8080           |
+----------------------+                      +-------------------------+
```

---

## 4. Características del escenario

- **Motor:** Docker
- **Protocolo:** HTTP
- **Lenguaje:** Python
- **Puerto:** 8080
- **Métrica clave:** Latencia (ms)

---

## 5. Uso típico

- El cliente ejecuta un **benchmark.py** enviando 100 solicitudes TCP.
- Se calculan métricas de latencia: promedio, mínimo y máximo.
- Los resultados se registran en la tabla comparativa de escenarios.
