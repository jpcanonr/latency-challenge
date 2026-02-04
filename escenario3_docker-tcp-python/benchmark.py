# benchmark.py
import socket
import time
import statistics

HOST = "127.0.0.1"   # dirección del servidor
PORT = 8080          # puerto del servidor
ITERACIONES = 100

latencias = []

for i in range(ITERACIONES):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        start = time.time()
        s.sendall(b"estimulo")
        data = s.recv(1024)
        elapsed = (time.time() - start) * 1000  # convertir a ms
        latencias.append(elapsed)

        # opcional: imprimir cada resultado
        print(f"Iteración {i+1}: {elapsed:.3f} ms - Respuesta: {data.decode()}")

# calcular métricas
promedio = statistics.mean(latencias)
minimo = min(latencias)
maximo = max(latencias)

print("\n📊 Resultados del benchmark")
print(f"Solicitudes: {ITERACIONES}")
print(f"Latencia promedio: {promedio:.3f} ms")
print(f"Latencia mínima: {minimo:.3f} ms")
print(f"Latencia máxima: {maximo:.3f} ms")
