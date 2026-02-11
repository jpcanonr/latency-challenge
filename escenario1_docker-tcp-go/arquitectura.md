# Escenario 1: Docker + TCP + Go

Este documento describe la arquitectura y el flujo de comunicación del primer escenario de pruebas de latencia.

---

## 1. Componentes principales

- **Cliente Go (host):**
  - Programa escrito en Go que envía solicitudes TCP.
  - Corre en la máquina host (fuera del contenedor).

- **Servidor Go (contenedor Docker):**
  - Programa escrito en Go que escucha en el puerto `8080`.
  - Corre dentro de un contenedor Docker.

- **Conexión TCP:**
  - Canal de comunicación entre el cliente y el servidor.
  - Flujo de datos:  
    - El cliente envía un **“estimulo”**.  
    - El servidor responde con **“respuesta”**.

---

## 2. Flujo de comunicación

1. El **cliente Go** abre una conexión TCP hacia `localhost:8080`.
2. Envía el mensaje `"estimulo"`.
3. El **servidor Go** dentro del contenedor recibe el mensaje.
4. El servidor procesa y devuelve `"respuesta"`.
5. El cliente mide el tiempo transcurrido (latencia).
6. Se repite el proceso para múltiples iteraciones.

---

## 3. Arquitectura visual

```plaintext
+-------------------+        TCP        +-------------------------+
|                   | ----------------> |                         |
|   Cliente Go      |                   |   Docker Container      |
|   (Host Machine)  | <---------------- |   Servidor Go (TCP)     |
|                   |   "respuesta"     |   Puerto 8080           |
+-------------------+                   +-------------------------+
         "estimulo"
```

---

## 4. Características del escenario

- **Motor:** Docker
- **Protocolo:** TCP
- **Lenguaje:** Go
- **Puerto:** 8080
- **Métrica clave:** Latencia (ms)

---

## 5. Uso típico

- El cliente ejecuta un **benchmark.go** enviando 100 solicitudes TCP.
- Se calculan métricas de latencia: promedio, mínimo y máximo.
- Los resultados se registran en la tabla comparativa de escenarios.
