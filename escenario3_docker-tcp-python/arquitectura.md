# Escenario 3: Docker + TCP + Python

Este documento describe la arquitectura y el flujo de comunicación del tercer escenario de pruebas de latencia.

---

## 1. Componentes principales

- **Cliente Python (host):**
  - Script en Python que envía solicitudes TCP.
  - Corre en la máquina host (fuera del contenedor).

- **Servidor Python (contenedor Docker):**
  - Script en Python que escucha en el puerto `8080` usando `socket`.
  - Corre dentro de un contenedor Docker.

- **Conexión TCP:**
  - Canal de comunicación entre el cliente y el servidor.
  - Flujo de datos:  
    - El cliente envía un **“estimulo”**.  
    - El servidor responde con **“respuesta”**.

---

## 2. Flujo de comunicación

1. El **cliente Python** abre una conexión TCP hacia `127.0.0.1:8080`.
2. Envía el mensaje `"estimulo"`.
3. El **servidor Python** dentro del contenedor recibe el mensaje.
4. El servidor responde con `"respuesta"`.
5. El cliente mide el tiempo transcurrido (latencia).
6. Se repite el proceso para múltiples iteraciones.

---

## 3. Arquitectura visual

```plaintext
+-------------------+        TCP        +-------------------------+
|                   |    "estimulo"     |                         |
|  Cliente Python   | ----------------> |  Docker Container       |
|  (Host Machine)   |                   |  Servidor Python (TCP)  |
|                   | <---------------- |  Puerto 8080            |
|                   |   "respuesta"     |                         |
+-------------------+                   +-------------------------+
```

---

## 4. Características del escenario

- **Motor:** Docker
- **Protocolo:** TCP
- **Lenguaje:** Python
- **Puerto:** 8080
- **Métrica clave:** Latencia (ms)

---

## 5. Uso típico

- El cliente ejecuta un **benchmark.py** enviando 100 solicitudes TCP.
- Se calculan métricas de latencia: promedio, mínimo y máximo.
- Los resultados se registran en la tabla comparativa de escenarios.
