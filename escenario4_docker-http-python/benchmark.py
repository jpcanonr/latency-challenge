# benchmark.py
import requests
import time
import statistics

URL = "http://127.0.0.1:8080"
ITERACIONES = 100

latencias = []

for i in range(ITERACIONES):
    start = time.time()
    resp = requests.get(URL)
    elapsed = (time.time() - start) * 1000  # convertir a ms
    latencias.append(elapsed)

    # opcional: imprimir cada resultado
    print(f"Iteración {i+1}: {elapsed:.3f} ms")

# calcular métricas
promedio = statistics.mean(latencias)
minimo = min(latencias)
maximo = max(latencias)

print("\n📊 Resultados del benchmark")
print(f"Solicitudes: {ITERACIONES}")
print(f"Latencia promedio: {promedio:.3f} ms")
print(f"Latencia mínima: {minimo:.3f} ms")
print(f"Latencia máxima: {maximo:.3f} ms")
